package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	core_config "github.com/Mirwinli/coffe_plus/internal/core/config"
	core_infrastructure_cloudinary2 "github.com/Mirwinli/coffe_plus/internal/core/infrastructure/cloud/cloudinary"
	core_infrastructure_ordernotifier "github.com/Mirwinli/coffe_plus/internal/core/infrastructure/order_notifier"
	core_infrastructure_telegram_bot "github.com/Mirwinli/coffe_plus/internal/core/infrastructure/telegram/bot"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_goredis_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/cached/redis/pool/goredis"
	core_postgres_pgx "github.com/Mirwinli/coffe_plus/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/Mirwinli/coffe_plus/internal/core/transport/http/middleware"
	core_http_server "github.com/Mirwinli/coffe_plus/internal/core/transport/http/server"
	core_http_jwt "github.com/Mirwinli/coffe_plus/internal/core/transport/http/tokens/jwt"
	auth_service "github.com/Mirwinli/coffe_plus/internal/feature/auth"
	auth_transport_http "github.com/Mirwinli/coffe_plus/internal/feature/auth/adapters/in/transport/http"
	auth_postgres "github.com/Mirwinli/coffe_plus/internal/feature/auth/adapters/out/auth_repository"
	cart_service "github.com/Mirwinli/coffe_plus/internal/feature/cart"
	cart_adapters_in "github.com/Mirwinli/coffe_plus/internal/feature/cart/adapters/in"
	cart_adapters_out_redis "github.com/Mirwinli/coffe_plus/internal/feature/cart/adapters/out/redis"
	category_service "github.com/Mirwinli/coffe_plus/internal/feature/category"
	category_adapters_in "github.com/Mirwinli/coffe_plus/internal/feature/category/adapters/in"
	category_adapters_out "github.com/Mirwinli/coffe_plus/internal/feature/category/adapters/out"
	order_service "github.com/Mirwinli/coffe_plus/internal/feature/order"
	order_adapters_in "github.com/Mirwinli/coffe_plus/internal/feature/order/adapters/in"
	order_adapters_out_posgtres "github.com/Mirwinli/coffe_plus/internal/feature/order/adapters/out/posgtres"
	products_service "github.com/Mirwinli/coffe_plus/internal/feature/products"
	products_adapters_in_products_transport_http "github.com/Mirwinli/coffe_plus/internal/feature/products/adapters/in/transport/http"
	products_adapters_out_cache "github.com/Mirwinli/coffe_plus/internal/feature/products/adapters/out/cache"
	products_adapters_out_postgres "github.com/Mirwinli/coffe_plus/internal/feature/products/adapters/out/postgres"
	shop_service "github.com/Mirwinli/coffe_plus/internal/feature/shop"
	shop_adapters_out_redis "github.com/Mirwinli/coffe_plus/internal/feature/shop/adapters/out/redis"
	"go.uber.org/zap"

	_ "github.com/Mirwinli/coffe_plus/docs"
)

// @title Golang Cafe-Shop API
// @version 1.0
// @description Cafe-Shop Application API scheme
// @host 127.0.0.1:5050
// @Basepath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
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
	authHTTPHandler := auth_transport_http.NewAuthHTTPHandler(authService, JWTConfig, authRepository)

	logger.Debug("initializing feature", zap.String("feature", "categories"))
	categoryRepository := category_adapters_out.NewCategoryRepository(postgresPool)
	categoryService := category_service.NewCategoryService(categoryRepository)
	categoryHTTPHandler := category_adapters_in.NewCategoryHTTPHandler(
		categoryService,
		JWTConfig,
		authRepository,
	)

	logger.Debug("initializing feature", zap.String("feature", "products"))
	imageLoader, err := core_infrastructure_cloudinary2.NewCloudinaryUploader(
		*core_infrastructure_cloudinary2.NewMustConfig(),
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

	logger.Debug("initializing feature", zap.String("feature", "carts"))
	cartsRepository := cart_adapters_out_redis.NewCartRepository(redisPool, productsRepository)
	cartsService := cart_service.NewCartService(cartsRepository)
	cartsHTTPHandler := cart_adapters_in.NewCartHTTPHandler(cartsService, JWTConfig, authRepository)

	logger.Debug("initializing feature", zap.String("feature", "order"))
	orderNotifierConfig := core_infrastructure_ordernotifier.NewMustConfig()
	orderNotifier := core_infrastructure_ordernotifier.NewOrderNotifier(orderNotifierConfig)

	orderRepository := order_adapters_out_posgtres.NewOrderRepository(postgresPool)
	orderService := order_service.NewOrderService(orderRepository, cartsRepository, orderNotifier)
	orderHTTPHandler := order_adapters_in.NewOrderHTTPHandler(orderService, JWTConfig, authRepository)

	logger.Debug("initializing feature", zap.String("feature", "shop"))
	shopRepository := shop_adapters_out_redis.NewShopRepository(redisPool)
	shopService := shop_service.NewShopService(&shopRepository)

	logger.Debug("initializing telegram bot")

	chatID, err := strconv.Atoi(os.Getenv("TELEGRAM_CHAT_ID_COOKER"))
	if err != nil {
		log.Fatal("failed to get chatID from env")
	}

	bot, err := core_infrastructure_telegram_bot.NewBot(
		os.Getenv("TELEGRAM_API_TOKEN"),
		int64(chatID),
		orderService,
		&shopService,
		*logger,
	)
	if err != nil {
		logger.Fatal("failed to initialize telegram bot", zap.Error(err))
	}

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
	apiVersionRouterV1.RegisterRoutes(cartsHTTPHandler.Routes()...)
	apiVersionRouterV1.RegisterRoutes(orderHTTPHandler.Routes()...)

	httpServer.RegisterApiVersionRouter(apiVersionRouterV1)
	httpServer.RegisterSwagger()

	go bot.Start(ctx)

	if err = httpServer.Run(ctx); err != nil {
		logger.Error("failed to start HTTP server", zap.Error(err))
	}
}
