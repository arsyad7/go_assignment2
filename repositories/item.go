package repositories

import (
	"go_assignment2/models"

	"gorm.io/gorm"
)

type ItemRepo interface {
	GetItem(id, orderId uint) (*models.Item, error)
	UpdateItem(id uint, req *models.Item) error
}

type itemRepo struct {
	db *gorm.DB
}

func NewItemRepo(db *gorm.DB) ItemRepo {
	return &itemRepo{db}
}

func (i *itemRepo) GetItem(id, orderId uint) (*models.Item, error) {
	var item models.Item
	err := i.db.Where("id = ?", id).Where("order_id = ?", orderId).First(&item).Error
	return &item, err
}

func (i *itemRepo) UpdateItem(id uint, req *models.Item) error {
	var item models.Item
	err := i.db.Model(&item).Where("id = ?", id).Updates(models.Item{Description: req.Description, ItemCode: req.ItemCode, Quantity: req.Quantity}).Error
	return err
}
