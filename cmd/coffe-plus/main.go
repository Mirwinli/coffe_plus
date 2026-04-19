package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_config "github.com/Mirwinli/coffe_plus/internal/core/config"
	core_infrastructure_cloudinary "github.com/Mirwinli/coffe_plus/internal/core/infrastructure/cloudinary"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_goredis_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/cached/redis/pool/goredis"
	core_postgres_pgx "github.com/Mirwinli/coffe_plus/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/Mirwinli/coffe_plus/internal/core/transport/http/middleware"
	core_http_server "github.com/Mirwinli/coffe_plus/internal/core/transport/http/server"
	core_http_jwt "github.com/Mirwinli/coffe_plus/internal/core/transport/http/tokens/jwt"
	auth_service "github.com/Mirwinli/coffe_plus/internal/feature/auth"
	auth_transport_http "github.com/Mirwinli/coffe_plus/internal/feature/auth/adapters/in/transport/http"
	"github.com/Mirwinli/coffe_plus/internal/feature/auth/adapters/out/auth_repository"
	category_service "github.com/Mirwinli/coffe_plus/internal/feature/category"
	category_adapters_in "github.com/Mirwinli/coffe_plus/internal/feature/category/adapters/in"
	category_adapters_out "github.com/Mirwinli/coffe_plus/internal/feature/category/adapters/out"
	products_service "github.com/Mirwinli/coffe_plus/internal/feature/products"
	products_adapters_in_products_transport_http "github.com/Mirwinli/coffe_plus/internal/feature/products/adapters/in/transport/http"
	products_adapters_out_cache "github.com/Mirwinli/coffe_plus/internal/feature/products/adapters/out/cache"
	products_adapters_out_postgres "github.com/Mirwinli/coffe_plus/internal/feature/products/adapters/out/postgres"
	"go.uber.org/zap"

	_ "github.com/Mirwinli/coffe_plus/docs"
)

// @title Golang Cafe-Shop API
// @version 1.0
// @description Cafe-Shop Application API scheme
// @host 127.0.0.1:5050
// @Basepath /api/v1
func main() {
	config := core_config.NewMustConfig()
	time.Local = config.TimeZone

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("Failed to create logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("application time zone", zap.Any("time_zone", config.TimeZone))

	logger.Debug("initializing postgres connection pool")
	postgresPool, err := core_postgres_pgx.NewPool(
		ctx,
		core_postgres_pgx.NewMustConfig(),
	)
	if err != nil {
		logger.Fatal("failed to initialize postgres connection pool", zap.Error(err))
	}
	defer postgresPool.Close()

	logger.Debug("initializing redis connection pool")
	redisPool, err := core_goredis_pool.NewPool(
		ctx,
		core_goredis_pool.NewMustConfig(),
	)
	if err != nil {
		logger.Fatal("failed to initialize redis connection pool", zap.Error(err))
	}
	defer redisPool.Close()

	logger.Debug("initializing JWT config")
	JWTConfig := core_http_jwt.NewMustConfig()

	logger.Debug("initializing feature", zap.String("feature", "auth"))
	authRepository := auth_postgres.NewAuthRepository(postgresPool, redisPool)
	authService := auth_service.NewAuthService(authRepository, JWTConfig)
	authHTTPHandler := auth_transport_http.NewAuthHTTPHandler(authService, JWTConfig)

	logger.Debug("initializing feature", zap.String("feature", "categories"))
	categoryRepository := category_adapters_out.NewCategoryRepository(postgresPool)
	categoryService := category_service.NewCategoryService(categoryRepository)
	categoryHTTPHandler := category_adapters_in.NewCategoryHTTPHandler(
		categoryService,
		JWTConfig,
		authRepository,
	)

	logger.Debug("initializing feature", zap.String("feature", "products"))
	imageLoader, err := core_infrastructure_cloudinary.NewCloudinaryUploader(
		*core_infrastructure_cloudinary.NewMustConfig(),
	)
	if err != nil {
		logger.Fatal("failed to initialize cloudinary uploader", zap.Error(err))
	}

	productPostgresRepository := products_adapters_out_postgres.NewRepository(postgresPool)
	productsRepository := products_adapters_out_cache.NewCacheRepository(redisPool, productPostgresRepository)
	productsService := products_service.NewProductService(productsRepository, imageLoader)
	productsHTTPHandler := products_adapters_in_products_transport_http.NewProductsHTTPHandler(
		productsService,
		JWTConfig,
		authRepository,
	)

	logger.Debug("initializing HTTP server")
	httpConfig := core_http_server.NewMustConfig()
	httpServer := core_http_server.NewHTTPServer(
		httpConfig,
		logger,
		core_http_middleware.CORS(httpConfig.AllowedOrigins),
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	apiVersionRouterV1 := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouterV1.RegisterRoutes(authHTTPHandler.Routes()...)
	apiVersionRouterV1.RegisterRoutes(productsHTTPHandler.Routes()...)
	apiVersionRouterV1.RegisterRoutes(categoryHTTPHandler.Routes()...)

	httpServer.RegisterApiVersionRouter(apiVersionRouterV1)
	httpServer.RegisterSwagger()

	if err = httpServer.Run(ctx); err != nil {
		logger.Error("failed to start HTTP server", zap.Error(err))
	}
}
