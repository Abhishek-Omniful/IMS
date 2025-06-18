package routes

import (
	"github.com/Abhishek-Omniful/IMS/pkg/controllers"
	"github.com/omniful/go_commons/http"
)

func Initialize(s *http.Server) {
	s.GET("/", controllers.ServeHome)

	v1 := s.Engine.Group("/api/v1")
	{
		//validators
		orders := v1.Group("/validators")
		{
			orders.POST("/validate_order", controllers.ValidateHubAndSKU)
			orders.POST("/validate_inventory", controllers.ValidateAndUpdateInventory)
		}
		//hubs
		hubs := v1.Group("/hubs")
		{
			hubs.GET("", controllers.GetHubs)
			hubs.POST("", controllers.CreateHub)
			hubs.PUT("/:id", controllers.UpdateHub)
			hubs.DELETE("/:id", controllers.DeleteHub)
		}
		//skus
		skus := v1.Group("/skus")
		{
			skus.GET("", controllers.GetSKUs) // also handles the filtering logic
			skus.POST("", controllers.CreateSKU)
			skus.PUT("/:id", controllers.UpdateSKU)
			skus.DELETE("/:id", controllers.DeleteSKU)
		}
		// tenants
		tenants := v1.Group("/tenants")
		{
			tenants.GET("", controllers.GetTenants)
			tenants.POST("", controllers.CreateTenant)
			tenants.PUT("/:id", controllers.UpdateTenant)
			tenants.DELETE("/:id", controllers.DeleteTenant)
		}

	}

}
