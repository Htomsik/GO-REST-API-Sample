package apiServer

import "github.com/sirupsen/logrus"

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
	}
}

// Start ...
func (api *APIServer) Start() error {

	if err := api.configureLogger(); err != nil {
		return err
	}

	api.logger.Info("API server starting")

	return nil
}

func (api *APIServer) configureLogger() error {
	lvl, err := logrus.ParseLevel(api.config.LogLevel)

	if err != nil {
		return err
	}

	api.logger.SetLevel(lvl)

	return nil
}
