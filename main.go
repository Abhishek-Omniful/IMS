package main

import (
	"log"
	"time"

	"github.com/Abhishek-Omniful/IMS/context"
	"github.com/Abhishek-Omniful/IMS/migrations"
	"github.com/Abhishek-Omniful/IMS/pkg/appinit"
	"github.com/Abhishek-Omniful/IMS/pkg/routes"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/http"
)

func init() {
	db := appinit.GetDB()
	if db == nil {
		log.Panic("Failed to connect to the database")
	}
	log.Println("Connected to the database successfully")
	migrations.RunMigration()
}

func main() {

	ctx := context.GetContext()

	server := http.InitializeServer(
		config.GetString(ctx, "server.port"), // Port to listen
		10*time.Second,                       // Read timeout
		10*time.Second,                       // Write timeout
		70*time.Second,                       // Idle timeout
		false,
	)

	routes.Initialize(server)
	err := server.StartServer("ims-service")
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}

}
