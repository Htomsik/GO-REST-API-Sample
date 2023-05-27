package apiServer

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (api *APIServer) Start() error {

	if err := api.configureLogger(); err != nil {
		return err
	}

	api.configureRouter()

	api.logger.Info("API server starting")

	return http.ListenAndServe(api.config.Addr, api.router)
}

func (api *APIServer) configureLogger() error {
	lvl, err := logrus.ParseLevel(api.config.LogLevel)

	if err != nil {
		return err
	}

	api.logger.SetLevel(lvl)

	return nil
}

func (api *APIServer) configureRouter() {
	api.router.HandleFunc("/hello", api.handleHello())
}

func (api *APIServer) handleHello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "Hello")
	}
}
