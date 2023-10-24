package server

import (
	"context"
	_ "embed"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	"github.com/julienschmidt/httprouter"
)

var (
	//go:embed index.html.template
	pages string
)

type Server interface {
	Close() error
	ListenAndServe()
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type server struct {
	ctx    context.Context
	router *httprouter.Router
	status atomic.Int32
	svr    *http.Server
	tmpls *template.Template
}

func (s *server) Close() error {
	s.status.Store(http.StatusInternalServerError)
	cctx, cancel := context.WithTimeout(s.ctx, 30*time.Second)
	defer cancel()
	return s.svr.Shutdown(cctx)
}

func (s *server) ListenAndServe() {
	idleConnsClosed := make(chan interface{})

	go func() {
		defer close(idleConnsClosed)
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		s.Close()
	}()

	s.status.Store(http.StatusOK)
	if err := s.svr.ListenAndServe(); err != http.ErrServerClosed {
		slog.Error(err.Error())
	}

	<-idleConnsClosed
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func New(ctx context.Context) Server {
	return newServer(ctx)
}

func newServer(ctx context.Context) *server {
	rtr := httprouter.New()
	svr := &server{
		ctx:    ctx,
		router: rtr,
		status: atomic.Int32{},
		svr: &http.Server{
			Addr:              ":8080",
			Handler:           rtr,
			IdleTimeout:       5 * time.Minute,
			ReadHeaderTimeout: time.Minute,
		},
	}

	svr.status.Store(http.StatusServiceUnavailable)
	svr.routes()
	svr.tmpls = template.Must(template.New("").Parse(pages))

	return svr
}
