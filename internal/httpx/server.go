package httpx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Router interface {
	Add(*mux.Router)
}

type Server struct {
	svr *http.Server
}

func NewServer(port string, readTimeOut time.Duration, writeTimeOut time.Duration) *Server {
	return &Server{
		svr: &http.Server{
			Addr:         fmt.Sprintf(":%s", port),
			ReadTimeout:  readTimeOut,
			WriteTimeout: writeTimeOut,
		},
	}
}

func (s *Server) index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s Server) handler(routers []Router) http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	r.Methods(http.MethodGet).Path("/").Name("info")

	apis := r.PathPrefix("/v1").Subrouter()
	for _, router := range routers {
		router.Add(apis)
	}
	return r
}

func (s *Server) Start(routers []Router) {
	s.svr.Handler = s.handler(routers)
	go func() {
		err := s.svr.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
}

func (s *Server) ShutDown(ctx context.Context) error {
	ctxTo, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	return s.svr.Shutdown(ctxTo)
}
