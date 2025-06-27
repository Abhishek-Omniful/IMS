package services

import (
	"time"

	"github.com/Abhishek-Omniful/IMS/mycontext"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/db/sql/postgres"
	"github.com/omniful/go_commons/log"
)

var logger = log.DefaultLogger()
var db *postgres.DbCluster

func ConnectDB() {
	ctx := mycontext.GetContext()
	myHost := config.GetString(ctx, "postgresql.master.host")
	myPort := config.GetString(ctx, "postgresql.master.port")
	myUsername := config.GetString(ctx, "postgresql.master.username")
	myPassword := config.GetString(ctx, "postgresql.master.password")
	myDbname := config.GetString(ctx, "postgresql.database")
	maxOpenConns := config.GetInt(ctx, "postgresql.maxOpenConns")
	maxIdleConns := config.GetInt(ctx, "postgresql.maxIdleConns")
	connMaxLifetime := config.GetInt(ctx, "postgresql.connMaxLifetime")
	debugMode := config.GetBool(ctx, "postgresql.debugMode")

	masterConfig := postgres.DBConfig{
		Host:               myHost,
		Port:               myPort,
		Username:           myUsername,
		Password:           myPassword,
		Dbname:             myDbname,
		MaxOpenConnections: maxOpenConns,
		ConnMaxLifetime:    time.Duration(connMaxLifetime) * time.Second,
		MaxIdleConnections: maxIdleConns,
		DebugMode:          debugMode,
	}

	slavesConfig := make([]postgres.DBConfig, 0)

	db = postgres.InitializeDBInstance(masterConfig, &slavesConfig)
	logger.Info("Connecting to the database...")
}

func GetDB() *postgres.DbCluster {
	return db
}
