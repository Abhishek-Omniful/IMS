package init

import (
	"github.com/Abhishek-Omniful/IMS/pkg/services"
	"github.com/omniful/go_commons/log"
)

var logger *log.Logger

func init() {
	logger = log.DefaultLogger()
	services.ConnectDB()
	logger.Infof("Connected to the database successfully")
	services.ConnectRedis()
	logger.Infof("Connected to Redis successfully")
}

func Initialize() {
	logger.Info("Application initialized successfully")
}
