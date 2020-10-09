package main

import (
	"delivery_svc/config"
	_ "delivery_svc/docs"
	"delivery_svc/routers"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
	"googlemaps.github.io/maps"
	"log"
	"net/http"
	"os"
	"time"
)

// @title Delivery service
// @version 1.0
// @description Delivery service API
// @termsOfService https://www.andrey-derma.com

// @contact.name API Support
// @contact.url https://www.andrey-derma.com
// @contact.email andrey.derma@gmail.com

// @BasePath /v1

func main() {
	r := routers.SetupRouter()
	db := config.InitDB()
	mapsClientInstance, errGMaps := maps.NewClient(maps.WithAPIKey(os.Getenv("MAP_API_KEY")))
	if errGMaps != nil {
		log.Fatalf("fatal error: %s", errGMaps)
	}
	r.Mount("/v1/delivery_svc", SubRoutes(db, mapsClientInstance))

	s := &http.Server{
		Addr:           os.Getenv("APP_PORT"),
		Handler:        r,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

func SubRoutes(db *gorm.DB, mapsClientInstance *maps.Client) chi.Router {
	r := chi.NewRouter()
	r.Get("/health_check", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	r.Group(func(r chi.Router) {
		api := InitializeAPI(db, mapsClientInstance)
		r.Post("/orders", api.Orders)
		r.Get("/orders", api.ListOrders)
		r.Patch("/orders/{id}", api.TakeOrder)
	})

	return r
}
