package main

import (
	"log"
	"time"

	"github.com/Abhishek-Omniful/IMS/mycontext"

	"github.com/Abhishek-Omniful/IMS/pkg/routes"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/http"
)

func main() {

	ctx := mycontext.GetContext()

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
