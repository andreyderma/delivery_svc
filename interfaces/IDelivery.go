package interfaces

import (
	"delivery_svc/model"
	"delivery_svc/utils"
)

type IDeliveryRepository interface {
	FindAllOrders(page int, limit int) *utils.Paginator
	SaveOrders(orders model.Orders) (model.Orders, error)
	TakeUpOrder(id int) error
}
