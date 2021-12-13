package http

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/abdabariassalam/majoo/config"
	"github.com/abdabariassalam/majoo/internal/repository"
	"github.com/abdabariassalam/majoo/internal/service"
	"github.com/ilyakaznacheev/cleanenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Execute() {
	// Configuration
	var cfg config.Config

	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// connect db
	db, err := gorm.Open(mysql.Open(cfg.Mysql.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Tidak Konek DB Errornya : %s", err)
	}

	repository := repository.New(&repository.Args{
		DB:  db,
		Cfg: &cfg,
	})
	service := service.New(cfg, repository)

	// Run
	handler := NewHandler(service)
	r := NewRouter(*handler)
	r.routes()
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = cfg.Base.Port
	}
	r.router.Run(port)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Fatalf("app - Run - signal: %s", s.String())
	default:
	}
}
