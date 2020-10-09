package delivery

import (
	"delivery_svc/config"
	"delivery_svc/interfaces"
	"delivery_svc/model"
	"delivery_svc/utils"
	"errors"
	"github.com/jinzhu/gorm"
	"log"
)

type DeliveryRepository struct {
	Db *gorm.DB
}

func ProvideDeliveryRepository(Db *gorm.DB) interfaces.IDeliveryRepository {
	return &DeliveryRepository{
		Db: Db,
	}
}

func (p *DeliveryRepository) SaveOrders(orders model.Orders) (model.Orders, error) {
	errSave := p.Db.Save(&orders).Error
	if errSave != nil {
		return model.Orders{}, errSave
	}
	return orders, nil
}

func (p *DeliveryRepository) FindAllOrders(page int, limit int) *utils.Paginator {

	var m []model.Orders

	paginator := utils.Paging(&utils.Param{
		DB:      p.Db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: false,
	}, &m)

	return paginator
}

func (p *DeliveryRepository) TakeUpOrder(id int) error {
	tx := p.Db.Begin()
	var mOrder model.Orders

	//Locking row for update to make sure only one can take the order
	errGet := tx.Where(
		"id = ?", id).Set("gorm:query_option", "FOR UPDATE").First(&mOrder).Error
	if errGet != nil {
		tx.Rollback()
		return errGet
	}
	if mOrder.Status == config.TAKEN_STATUS {
		tx.Rollback()
		return errors.New("Your order already taken")
	}

	//Update status and commit transaction
	mOrder.Status = config.TAKEN_STATUS
	if err := tx.Save(&mOrder).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	if txCommit := tx.Commit().Error; txCommit != nil {
		log.Println(txCommit.Error())
		tx.Rollback()
		return txCommit
	}
	return nil
}
