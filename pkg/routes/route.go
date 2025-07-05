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
			orders.GET("/validate_order/:sku_id/:hub_id", controllers.ValidateOrderRequest)
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

		//products
		products := v1.Group("/products")
		{
			products.GET("", controllers.GetProducts)
			products.POST("", controllers.CreateProduct)
			products.PUT("/:id", controllers.UpdateProduct)
			products.DELETE("/:id", controllers.DeleteProduct)
		}

		//sellers
		sellers := v1.Group("/sellers")
		{
			sellers.GET("", controllers.GetSellers)
			sellers.POST("", controllers.CreateSeller)
			sellers.PUT("/:id", controllers.UpdateSeller)
			sellers.DELETE("/:id", controllers.DeleteSeller)
		}
		// inventory
		inventory := v1.Group("/inventory")
		{
			inventory.POST("/upsert", controllers.UpsertInventory)                 // Atomic upsert
			inventory.GET("/by-hub/:hub_id", controllers.GetInventoryByHub)        // View inventory for a hub
			inventory.GET("/by-sku/:sku_id", controllers.GetInventoryBySKU)        // View inventory for a SKU
			inventory.GET("/:sku_id/:hub_id", controllers.GetInventoryBySKUAndHub) // View inventory for a SKU in a specific hub
			inventory.GET("", controllers.GetAllInventory)                         // View all inventory
			inventory.GET("/check", controllers.CheckInventoryStatus)              // Check inventory status
		}

	}

}
