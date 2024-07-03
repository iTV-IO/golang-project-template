package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/quvonchbe05/golang-clean-architecture-template/app/http/handler"
	"github.com/quvonchbe05/golang-clean-architecture-template/app/rest"
	"github.com/quvonchbe05/golang-clean-architecture-template/config"
	"github.com/quvonchbe05/golang-clean-architecture-template/internal/repository"
	"github.com/quvonchbe05/golang-clean-architecture-template/internal/service"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func main() {
	cfg := config.Load(".")

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DB,
	)

	db, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("Failed to connect to postgres: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping postgres: %v", err)
	}

	log.Println("Postgres connection successfully done!")

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	router := rest.NewRoute(handler)

	srv := new(Server)
	if err := srv.Run(cfg.Port, router); err != nil {
		log.Fatalf(`Error occured while running http server: %s`, err.Error())
	}

}
