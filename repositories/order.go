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
	// var item models.Item
	// fmt.Printf("req.Items: %v\n", req.Items[0].BaseItem.ID)

	err := o.db.First(&order, "id = ?", id).Error
	if err != nil {
		return fmt.Errorf("Order with id %v not found", id)
	}

	itemRepo := NewItemRepo(o.db)
	for _, v := range req.Items {
		item, err := itemRepo.GetItem(v.BaseItem.ID, id)
		// err := o.db.Where("id = ?", v.BaseItem.ID).First(&item).Error
		if err != nil {
			return fmt.Errorf("Item with id %v not found", v.BaseItem.ID)
		}
		fmt.Println(item)
		err = itemRepo.UpdateItem(v.BaseItem.ID, &models.Item{
			Description: v.Description,
			Quantity:    v.Quantity,
			ItemCode:    v.ItemCode,
		})
		if err != nil {
			return fmt.Errorf("Item with id %v failed to update", v.BaseItem.ID)
		}
	}

	err = o.db.Model(&order).Where("id = ?", id).Updates(models.Order{CustomerName: req.CustomerName, OrderedAt: req.OrderedAt}).Error
	return err
}
