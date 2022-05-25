package repositories

import (
	"fmt"
	"go_assignment2/models"

	"gorm.io/gorm"
)

type OrderRepo interface {
	CreateOrder(order *models.Order) error
	GetOrders() (*[]models.Order, error)
	UpdateOrders(id uint, order *models.Order) error
	DeleteOrder(id uint) error
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return &orderRepo{db}
}

func (o *orderRepo) CreateOrder(order *models.Order) error {
	return o.db.Create(order).Error
}

func (o *orderRepo) GetOrders() (*[]models.Order, error) {
	var orders []models.Order
	err := o.db.Preload("Items").Find(&orders).Error
	return &orders, err
}

func (o *orderRepo) UpdateOrders(id uint, req *models.Order) error {
	var order models.Order

	err := o.db.First(&order, "id = ?", id).Error
	if err != nil {
		msg := fmt.Sprintf("Order with id %v not found", id)
		return fmt.Errorf(msg)
	}

	itemRepo := NewItemRepo(o.db)
	for _, v := range req.Items {
		item, err := itemRepo.GetItem(v.BaseItem.ID, id)

		if err != nil {
			msg := fmt.Sprintf("Item with id %v not found", v.BaseItem.ID)
			return fmt.Errorf(msg)
		}
		fmt.Println(item)
		err = itemRepo.UpdateItem(v.BaseItem.ID, &models.Item{
			Description: v.Description,
			Quantity:    v.Quantity,
			ItemCode:    v.ItemCode,
		})
		if err != nil {
			msg := fmt.Sprintf("Item with id %v failed to update", v.BaseItem.ID)
			return fmt.Errorf(msg)
		}
	}

	err = o.db.Model(&order).Where("id = ?", id).Updates(models.Order{CustomerName: req.CustomerName, OrderedAt: req.OrderedAt}).Error
	return err
}

func (o *orderRepo) DeleteOrder(id uint) error {
	itemRepo := NewItemRepo(o.db)
	var order models.Order

	err := o.db.Where("id = ?", id).First(&order).Error
	if err != nil {
		msg := fmt.Sprintf("Ordes with id %v not found", id)
		return fmt.Errorf(msg)
	}

	err = itemRepo.DeleteItem(id)
	if err != nil {
		msg := fmt.Sprintf("Item with order id %v not found", id)
		return fmt.Errorf(msg)
	}

	err = o.db.Where("id = ?", id).Delete(&order).Error
	if err != nil {
		msg := fmt.Sprintf("Ordes with id %v failed to delete", id)
		return fmt.Errorf(msg)
	}
	return nil
}
