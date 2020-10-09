// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"delivery_svc/apis"
	"delivery_svc/repositories/mysql/delivery"
	"delivery_svc/services"
	"github.com/jinzhu/gorm"
	"googlemaps.github.io/maps"
)

import (
	_ "delivery_svc/docs"
	_ "github.com/joho/godotenv/autoload"
)

// Injectors from wire.go:

func InitializeAPI(db *gorm.DB, mapsClientInstance *maps.Client) apis.API {
	iDeliveryRepository := delivery.ProvideDeliveryRepository(db)
	deliveryService := services.ProvideDeliveryService(iDeliveryRepository, mapsClientInstance)
	api := apis.ProvideAPI(deliveryService)
	return api
}