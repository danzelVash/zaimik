package main

import (
	_ "github.com/jackc/pgx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"os"
	"zaimik/internal/app/handler"
	"zaimik/internal/app/models"
	"zaimik/internal/app/repository"
	"zaimik/internal/app/repository/postgres"
	"zaimik/internal/app/service"
	"zaimik/internal/pkg/inmemory_storage"
	"zaimik/internal/pkg/logging"
)

// @title Zaimik API
// @version 1.0

// @host localhost
// @BasePath /

func main() {
	logger := logging.GetLogger()
	if err := initConfig(); err != nil {
		logger.Fatalf("error reading config file %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logger.Fatal("can`t read db password from env")
	}

	db, err := postgres.NewPostgresDb(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logger.Fatalf("error while init db: %s", err.Error())
	}

	repos := repository.NewRepository(db, inmemory_storage.NewDataStorage(), logger)
	services := service.NewService(repos, logger)
	handlers := handler.NewHandler(services, logger)

	c := handler.CorsSettings()
	c.Handler(handlers.InitRouters())
	//handlers.InitRouters()
	srv := models.Server{}
	if err := srv.Run(viper.GetString("http.port"), c.Handler(handlers.InitRouters())); err != nil {
		logger.Fatalf("error while trying to run server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("internal/app/configs")
	//viper.SetConfigName("config")
	viper.SetConfigName("deploy_config")
	return viper.ReadInConfig()
}
