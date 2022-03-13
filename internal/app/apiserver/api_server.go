package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

// APIServer - description of api configuration
type APIServer struct {
	cfg    *Config        // config of api server
	logger *logrus.Logger // logger for server
	router *mux.Router    // router of the project
}

// NewAPIServer - return new api server
func NewAPIServer(config *Config) *APIServer {
	return &APIServer{
		cfg:    config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Run - start server
func (a *APIServer) Run() error {
	if err := a.initLogger(); err != nil {
		return err
	}

	a.initRouter()

	a.logger.Infoln("APIServer start work")

	return nil
}

func (a *APIServer) initLogger() error {
	level, err := logrus.ParseLevel(a.cfg.Loglevel)
	if err != nil {
		return err
	}

	a.initRouter()

	a.logger.SetLevel(level)

	return http.ListenAndServe(a.cfg.Host, a.router)
}
