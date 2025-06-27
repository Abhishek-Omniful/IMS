package main

import (
	"log"

	imsinit "github.com/Abhishek-Omniful/IMS/init"
	"github.com/Abhishek-Omniful/IMS/mycontext"

	"github.com/Abhishek-Omniful/IMS/pkg/middlewares"
	"github.com/Abhishek-Omniful/IMS/pkg/routes"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/http"
)

func main() {

	ctx := mycontext.GetContext()
	imsinit.Initialize()
	server := http.InitializeServer(
		config.GetString(ctx, "server.port"),            // Port to listen
		config.GetDuration(ctx, "server.read_timeout"),  // Read timeout
		config.GetDuration(ctx, "server.write_timeout"), // Write timeout
		config.GetDuration(ctx, "server.idle_timeout"),  // Idle timeout
		false,
	)
	server.Use(middlewares.LogRequest(ctx))
	routes.Initialize(server)
	err := server.StartServer(config.GetString(ctx, "server.name"))
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}

}
