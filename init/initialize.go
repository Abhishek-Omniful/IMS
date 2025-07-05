package init

import (
	dbService "github.com/Abhishek-Omniful/IMS/pkg/integrations/db"
	redisService "github.com/Abhishek-Omniful/IMS/pkg/integrations/redis"
	"github.com/omniful/go_commons/log"
)

var logger *log.Logger

func init() {
	logger = log.DefaultLogger()
	dbService.ConnectDB()
	logger.Infof("Connected to the database successfully")
	redisService.ConnectRedis()
	logger.Infof("Connected to Redis successfully")
}

func Initialize() {
	logger.Info("Application initialized successfully")
}
