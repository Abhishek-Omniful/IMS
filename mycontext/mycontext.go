package mycontext

import (
	"context"
	"os"
	"time"

	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/log"
)

var ctx context.Context
var logger = log.DefaultLogger()

func init() {
	if os.Getenv("CONFIG_SOURCE") == "test" {
		return // Skip config during test
	}
	err := config.Init(time.Second * 10)

	if err != nil {
		logger.Panicf("Error while initialising config, err: %v", err)
		panic(err)
	}

	ctx, err = config.TODOContext()
	if err != nil {
		logger.Panicf("Failed to create context: %v", err)
	}
}

func GetContext() context.Context {
	return ctx
}
