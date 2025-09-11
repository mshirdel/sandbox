package http

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/mshirdel/sandbox/app"
	"github.com/mshirdel/sandbox/app/http/controller"
	"github.com/sirupsen/logrus"
)

const addr = "0.0.0.0:8080"

type Server struct {
	app    *app.Application
	server *http.Server
}

func NewHTTPServer(app *app.Application) *Server {
	controller := controller.NewController(app)

	return &Server{
		app: app,
		server: &http.Server{
			ReadTimeout:  time.Second * 2,
			Addr:         addr,
			WriteTimeout: time.Second * 3,
			IdleTimeout:  time.Minute * 2,
			Handler:      controller.Routes(),
		},
	}
}

func (s *Server) Start() {
	logrus.Infof("starting http server on: %s", addr)

	if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		logrus.Fatalf("failed starting http server: %v", err)
	}
}

func (s *Server) Shutdown() {
	logrus.Info("shutting down http server...")

	deadline, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.server.Shutdown(deadline); err != nil {
		logrus.Errorf("failed shutting down http server: %s", err)
	}
}
