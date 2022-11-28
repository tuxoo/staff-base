package app

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tuxoo/smart-loader/staff-base/internal/config"
	"github.com/tuxoo/smart-loader/staff-base/internal/controller/http"
	"github.com/tuxoo/smart-loader/staff-base/internal/repository"
	"github.com/tuxoo/smart-loader/staff-base/internal/server"
	"github.com/tuxoo/smart-loader/staff-base/internal/service"
	"os"
	"os/signal"
	"syscall"
)

// @title        Staff Base Application
// @version      1.0
// @description  API Server for ...

// @host      localhost:9000
// @BasePath  /api/v1

// Run initializes whole application

func Run() {
	fmt.Println(`
 _$$$$__$$$$$$__$$$$__$$$$$$_$$$$$$_$$$$$___$$$$___$$$$__$$$$$
 $$_______$$___$$__$$_$$_____$$_____$$__$$_$$__$$_$$_____$$___
 _$$$$____$$___$$$$$$_$$$$___$$$$___$$$$$__$$$$$$__$$$$__$$$$_
 ____$$___$$___$$__$$_$$_____$$_____$$__$$_$$__$$_____$$_$$___
 _$$$$____$$___$$__$$_$$_____$$_____$$$$$__$$__$$__$$$$__$$$$$
	`)

	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := config.NewPostgresPool(config.PostgresConfig{
		Host:            cfg.Postgres.Host,
		Port:            cfg.Postgres.Port,
		DB:              cfg.Postgres.DB,
		User:            cfg.Postgres.User,
		Password:        cfg.Postgres.Password,
		MaxConns:        cfg.Postgres.MaxConns,
		MinConns:        cfg.Postgres.MinConns,
		MaxConnLifetime: cfg.Postgres.MaxConnLifetime,
		MaxConnIdleTime: cfg.Postgres.MaxConnIdleTime,
	})
	if err != nil {
		logrus.Fatalf("error initializing postgres: %s", err.Error())
	}
	defer db.Close()

	repositories := repository.NewRepositories(db)
	services := service.NewServices(repositories)

	httpHandlers := http.NewHandler(services.EmployeeService)
	httpServer := server.NewHTTPServer(cfg, httpHandlers.Init(cfg.HTTP))

	go func() {
		if err := httpServer.Run(); err != nil {
			logrus.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logrus.Printf("STAFF BASE application has been started on :%s port", cfg.HTTP.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("STAFF BASE facade application shutting down")
}
