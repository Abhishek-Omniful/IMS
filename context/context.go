package context

import (
	"context"
	"log"
	"time"

	"github.com/omniful/go_commons/config"
)

var ctx context.Context

func init() {
	//mandatory to call config.init() before using the context
	err := config.Init(time.Second * 10)
	if err != nil {
		log.Panicf("Error while initialising config, err: %v", err)
		panic(err)
	}

	ctx, err = config.TODOContext()
	if err != nil {
		log.Panicf("Failed to create context: %v", err)
	}
}

func GetContext() context.Context {
	return ctx
}
