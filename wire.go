//+build wireinject

package main

import (
	"delivery_svc/apis"
	"delivery_svc/repositories/mysql/delivery"
	"delivery_svc/services"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"googlemaps.github.io/maps"
)

func InitializeAPI(db *gorm.DB, mapsClientInstance *maps.Client) apis.API {
	wire.Build(
		delivery.ProvideDeliveryRepository,
		services.ProvideDeliveryService,
		apis.ProvideAPI)

	return apis.API{}
}
