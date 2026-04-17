package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Mirwinli/coffe_plus/docs"
	core_logger "github.com/Mirwinli/coffe_plus/internal/core/logger"
	core_http_middlware "github.com/Mirwinli/coffe_plus/internal/core/transport/http/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"go.uber.org/zap"
)

type HTTPServer struct {
	mux    *http.ServeMux
	config Config
	log    *core_logger.Logger

	middleware []core_http_middlware.Middleware
}

func NewHTTPServer(
	config Config,
	log *core_logger.Logger,
	middleware ...core_http_middlware.Middleware,
) *HTTPServer {
	return &HTTPServer{
		mux:        http.NewServeMux(),
		config:     config,
		log:        log,
		middleware: middleware,
	}
}

func (s *HTTPServer) RegisterApiVersionRouter(routers ...*ApiVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)

		h := router.WithMiddleware()

		s.mux.Handle(
			prefix+"/",
			http.StripPrefix(
				prefix,
				h,
			),
		)
	}
}

func (s *HTTPServer) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		s.mux.Handle(pattern, route.WithMiddleware())
	}
}

func (s *HTTPServer) RegisterSwagger() {
	s.mux.Handle(
		"/swagger/",
		httpSwagger.Handler(
			httpSwagger.URL("/swagger/doc.json"),
			httpSwagger.DefaultModelsExpandDepth(-1),
		),
	)

	s.mux.HandleFunc(
		"/swagger/doc.json",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(docs.SwaggerInfo.ReadDoc()))
		},
	)
}

func (s *HTTPServer) Run(ctx context.Context) error {
	mux := core_http_middlware.ChainMiddleware(s.mux, s.middleware...)

	server := &http.Server{
		Addr:    s.config.Addr,
		Handler: mux,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		s.log.Warn("starting HTTP server", zap.String("addr", s.config.Addr))
		err := server.ListenAndServe()

		if err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				ch <- err
			}
		}
	}()

	select {
	case err := <-ch:
		return fmt.Errorf("listen and serve HTTP: %w", err)
	case <-ctx.Done():
		s.log.Warn("shutting down HTTP server")

		shutDownCtx, cancel := context.WithTimeout(ctx, s.config.ShutdownTimeout)
		defer cancel()

		err := server.Shutdown(shutDownCtx)
		if err != nil {
			_ = server.Close()

			return fmt.Errorf("shutdown HTTP server: %w", err)
		}

		s.log.Warn("HTTP server shutdown completed")
	}

	return nil
}
