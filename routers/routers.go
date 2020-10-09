package routers

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/swaggo/http-swagger"
	"os"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler)
	//Swagger only available in dev, we are not provide swagger docs in production
	port := os.Getenv("APP_PORT")
	env := os.Getenv("ENV")
	if env == "dev" {
		var swagUrl = fmt.Sprintf("http://localhost%s/v1/delivery_svc/swagger/doc.json", port)
		r.Get("/v1/delivery_svc/swagger/*", httpSwagger.Handler(
			httpSwagger.URL(swagUrl), //The url pointing to API definition"
		))
	}
	return r
}
