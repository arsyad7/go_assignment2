package services

import (
	"go_assignment2/models"
	"go_assignment2/params"
	"go_assignment2/repositories"
	"time"

	"gorm.io/gorm"
)

type OrderService struct {
	orderRepo repositories.OrderRepo
}

var db = gorm.DB{}
var repo = repositories.NewOrderRepo(&db)

func NewOrderService(repo repositories.OrderRepo) *OrderService {
	return &OrderService{
		orderRepo: repo,
	}
}

func (o *OrderService) CreateOrder(req *params.CreateOrder) *params.Response {
	model := models.Order{
		OrderedAt:    time.Now(),
		CustomerName: req.CustomerName,
		Items:        req.Items,
	}

	err := o.orderRepo.CreateOrder(&model)
	if err != nil {
		return &params.Response{
			Status:         400,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  201,
		Message: "Create Order Success",
	}
}

func (o *OrderService) GetOrders() *params.Response {
	orders, err := o.orderRepo.GetOrders()
	if err != nil {
		return &params.Response{
			Status:         400,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  200,
		Message: "success",
		Payload: orders,
	}
}

func (o *OrderService) UpdateOrder(id uint, req *params.CreateOrder) *params.Response {
	model := models.Order{
		OrderedAt:    req.OrderedAt,
		CustomerName: req.CustomerName,
		Items:        req.Items,
	}

	err := o.orderRepo.UpdateOrders(id, &model)
	if err != nil {
		return &params.Response{
			Status:         404,
			Error:          "NOT FOUND",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  200,
		Message: "Update Order Success",
	}
}

func (o *OrderService) DeleteOrder(id uint) *params.Response {
	err := o.orderRepo.DeleteOrder(id)
	if err != nil {
		return &params.Response{
			Status:         404,
			Error:          "NOT FOUND",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  200,
		Message: "Delete Success",
	}
}
