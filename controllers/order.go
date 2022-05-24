package controllers

import (
	"go_assignment2/params"
	"go_assignment2/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(srvcs *services.OrderService) *OrderController {
	return &OrderController{
		orderService: *srvcs,
	}
}

func (o *OrderController) CreateOrder(c *gin.Context) {
	var req *params.CreateOrder

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err,
		})
		return
	}

	response := o.orderService.CreateOrder(req)
	c.JSON(response.Status, response)
}

func (o *OrderController) GetOrders(c *gin.Context) {
	response := o.orderService.GetOrders()
	c.JSON(response.Status, response)
}

func (o *OrderController) UpdateOrder(c *gin.Context) {
	var req *params.CreateOrder
	orderId := c.Param("id")

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err,
		})
		return
	}

	id, err := strconv.Atoi(orderId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err,
		})
		return
	}
	response := o.orderService.UpdateOrder(uint(id), req)
	c.JSON(response.Status, response)
}
