package http

import (
	"context"
	"errors"
	"github.com/deissh/go-utils"
	echoPrometheus "github.com/globocom/echo-prometheus"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rl-os/api/docs"
	"github.com/rl-os/api/middlewares/customerror"
	"github.com/rl-os/api/middlewares/customlogger"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"time"
)

var ProviderSet = wire.NewSet(New, NewRouter, NewOptions)

const WAIT_FOR_CONNECTIONS_BEFORE_SHUTDOWN = time.Second * 2

// Options is server configuration struct
type Options struct {
	Port    string
	Host    string
	APIAddr string
}

// Server is echo server struct
type Server struct {
	o          *Options
	logger     *zerolog.Logger
	router     *echo.Echo
	httpServer http.Server

	// DidFinishListen channel notif about stopping server
	DidFinishListen chan struct{}
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("http", o); err != nil {
		return nil, err
	}

	return o, err
}

func NewRouter(o *Options, log *zerolog.Logger) *echo.Echo {
	docs.SwaggerInfo.Host = utils.IfThenElse(
		o.APIAddr == "NOTSET",
		"localhost:2400",
		o.APIAddr,
	).(string)

	r := echo.New()
	r.HidePort = true
	r.HideBanner = true
	r.HTTPErrorHandler = customerror.CustomHTTPErrorHandler

	echoPrometheus.DefaultConfig.Namespace = ""
	r.Use(
		echoPrometheus.MetricsMiddleware(),
		middleware.RequestID(),
		//permission.GlobalMiddleware(s.App),
		//reqest_context.GlobalMiddleware(s.App),
		customlogger.Middleware(),
	)

	root := r.Group("")

	//api.New(app, root)
	root.GET("/docs/*", echoSwagger.WrapHandler)

	return r
}

func New(o *Options, log *zerolog.Logger, router *echo.Echo) (*Server, error) {
	l := log.With().
		Str("type", "http.Server").
		Logger()

	var s = &Server{
		o:      o,
		logger: &l,
		router: router,
	}

	return s, nil
}

func (s *Server) listen(addr string) {
	s.logger.Info().Msg("starting listener")
	if err := s.router.Start(addr); err != nil {
		s.logger.Error().
			Err(err).
			Msg("stopped")
	}

	close(s.DidFinishListen)
}

func (s *Server) Start() error {
	if s.o.Host == "" {
		return errors.New("get local ipv4 error")
	}

	addr := s.o.Host + ":" + s.o.Port

	s.DidFinishListen = make(chan struct{})
	go s.listen(addr)

	return nil
}

func (s *Server) Stop() error {
	s.logger.Info().Msg("Shutdown HTTP server")

	ctx, cancel := context.WithTimeout(context.Background(), WAIT_FOR_CONNECTIONS_BEFORE_SHUTDOWN)
	defer cancel()

	didShutdown := false
	for s.DidFinishListen != nil && !didShutdown {
		if err := s.router.Shutdown(ctx); err != nil {
			s.logger.Warn().
				Err(err).
				Msg("Unable to shutdown HTTP server")
			return err
		}

		timer := time.NewTimer(time.Millisecond * 50)
		select {
		case <-s.DidFinishListen:
			didShutdown = true
		case <-timer.C:
		}

		timer.Stop()
	}

	return s.router.Close()
}
