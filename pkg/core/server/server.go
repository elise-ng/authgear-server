package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/skygeario/skygear-server/pkg/core/handler"
	"github.com/skygeario/skygear-server/pkg/core/middleware"
	"github.com/skygeario/skygear-server/pkg/core/sentry"
)

// Server embeds a net/http server and has a gorillax mux internally
type Server struct {
	*http.Server

	router *mux.Router
}

// NewServer create a new Server with default option
func NewServer(
	addr string,
) Server {
	return NewServerWithOption(
		addr,
		DefaultOption(),
	)
}

// NewServerWithOption create a new Server
func NewServerWithOption(
	addr string,
	option Option,
) Server {
	rootRouter := mux.NewRouter()
	rootRouter.HandleFunc("/healthz", HealthCheckHandler)

	var appRouter *mux.Router
	if option.GearPathPrefix == "" {
		appRouter = rootRouter.NewRoute().Subrouter()
	} else {
		appRouter = rootRouter.PathPrefix(option.GearPathPrefix).Subrouter()
	}

	if option.IsAPIVersioned {
		appRouter = appRouter.PathPrefix("/{api_version}").Subrouter()
	}

	srv := Server{
		router: appRouter,
		Server: &http.Server{
			Addr:    addr,
			Handler: rootRouter,
		},
	}

	srv.Use(sentry.Middleware(sentry.DefaultClient.Hub))
	if option.RecoverPanic {
		srv.Use(middleware.RecoverMiddleware{}.Handle)
	}

	return srv
}

// Handle delegates gorilla mux Handler, and accept a HandlerFactory instead of Handler
func (s *Server) Handle(path string, hf handler.Factory) *mux.Route {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := hf.NewHandler(r)
		h.ServeHTTP(w, r)
	})

	return s.router.NewRoute().Path(path).Handler(handler)
}

// Use set middlewares to underlying router
func (s *Server) Use(mwf ...mux.MiddlewareFunc) {
	s.router.Use(mwf...)
}

// ServeHTTP makes Server a http.Handler.
// It is useful in testing.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Server.Handler.ServeHTTP(w, r)
}
