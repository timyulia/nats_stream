package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"nats"
	cache "nats/pkg/cashe"
	"nats/pkg/handler"
	"nats/pkg/repository"
	"nats/pkg/service"
	"nats/pkg/streaming"
	"net/http"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Password: os.Getenv("DB_PASSWORD"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	var c cache.InMemory
	c.InitCache()
	services := service.NewService(repos, c)
	err = services.Restore()
	if err != nil {
		logrus.Errorf("cannot restore cache: %s", err.Error())
	}
	handlers := handler.NewHandler(services)
	str := streaming.NewStream(services)
	err = str.NatsStreamingSetup()
	if err != nil {
		logrus.Fatalf("cannot connect to nats-streaming-server: %s", err.Error())
	}

	srv := new(nats.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection closing: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
