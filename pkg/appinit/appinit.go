package appinit

import (
	"log"
	"time"

	"github.com/Abhishek-Omniful/IMS/mycontext"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/db/sql/postgres"
	"github.com/omniful/go_commons/redis"
)

func ConnectDB() *postgres.DbCluster {
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
	// Initialize slavesConfig as an empty slice
	// they can be added later if needed
	slavesConfig := make([]postgres.DBConfig, 0) // read replicas

	db := postgres.InitializeDBInstance(masterConfig, &slavesConfig)
	log.Println("Connecting to the database...")
	return db
}

func ConnectRedis() *redis.Client {
	config := &redis.Config{
		Hosts:       []string{"127.0.0.1:6379"},
		PoolSize:    50,
		MinIdleConn: 10,
		DB:          0,
	}
	client := redis.NewClient(config)
	log.Println("Connecting to Redis...")
	return client
}

func GetDB() *postgres.DbCluster {
	db := ConnectDB()
	return db
}

func GetRedis() *redis.Client {
	rds := ConnectRedis()
	return rds
}
