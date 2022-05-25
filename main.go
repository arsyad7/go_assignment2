package main

import (
	"go_assignment2/controllers"
	"go_assignment2/database"
	"go_assignment2/repositories"
	"go_assignment2/services"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	PORT = ":8080"
)

func main() {
	db := database.StartDB()
	router := gin.Default()
	orderRepo := repositories.NewOrderRepo(db)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

	router.POST("/orders", orderController.CreateOrder)
	router.GET("/orders", orderController.GetOrders)
	router.PUT("/orders/:orderId", orderController.UpdateOrder)
	router.DELETE("/orders/:orderId", orderController.DeleteOrder)

	log.Println("server running at port ", PORT)
	router.Run(PORT)
}
