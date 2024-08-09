package main

import (
	"log"

	"github.com/quvonchbe05/golang-clean-architecture-template/app/http/controller"
	"github.com/quvonchbe05/golang-clean-architecture-template/app/routes"
	server "github.com/quvonchbe05/golang-clean-architecture-template/cmd"
	"github.com/quvonchbe05/golang-clean-architecture-template/config"
	"github.com/quvonchbe05/golang-clean-architecture-template/internal/repository"
	"github.com/quvonchbe05/golang-clean-architecture-template/internal/service"
)

// @title iTV-IPTV
// @version 2.0.0
// @contact.email davronbekov.otabek@gmail.com
// @host 127.0.0.1:8000
// @BasePath /api/v1
func main() {
	cfg := config.Load(".")
	db := config.NewPostgresDB(cfg)

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	controller := controller.NewController(service)

	router := routes.NewRoute(controller)

	srv := new(server.Server)
	if err := srv.Run(cfg.Port, router); err != nil {
		log.Fatalf(`Error occured while running http server: %s`, err.Error())
	}
}