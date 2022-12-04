package apiserver

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nkarpeev/telegram-logger/internal/app/sentryService"

	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {

	if err := s.ConfigureLogger(); err != nil {
		return err
	}

	s.logger.Info("starting api server...")
	s.ConfigureRouter()

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) ConfigureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) ConfigureRouter() {
	s.router.HandleFunc("/", s.handleHome()).Methods("GET")
	s.router.HandleFunc("/write", s.handleWriteMsg()).Methods("POST")
}

func (s *APIServer) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Home page telegram logger")
	}
}

func (s *APIServer) handleWriteMsg() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var p sentryService.Payload
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		go func() {
			sentryService.Write(p)
		}()

		io.WriteString(w, "Wrote!")
	}
}
