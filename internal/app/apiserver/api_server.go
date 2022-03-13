package apiserver

import (
	"github.com/sirupsen/logrus"
)

// APIServer - description of api configuration
type APIServer struct {
	cfg    *Config        // config of api server
	logger *logrus.Logger // logger for server
}

// NewAPIServer - return new api server
func NewAPIServer(config *Config) *APIServer {
	return &APIServer{
		cfg:    config,
		logger: logrus.New(),
	}
}

// Run - start server
func (a *APIServer) Run() error {
	if err := a.initLogger(); err != nil {
		return err
	}

	a.logger.Infoln("APIServer start work")

	return nil
}

func (a *APIServer) initLogger() error {
	level, err := logrus.ParseLevel(a.cfg.Loglevel)
	if err != nil {
		return err
	}

	a.logger.SetLevel(level)

	return nil
}
