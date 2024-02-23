package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/yokaracho/orderapps/internal/database"
	"github.com/yokaracho/orderapps/internal/logger"
	"github.com/yokaracho/orderapps/internal/logic"
	"os"
	"os/signal"
)

func main() {
	var getLogger = logger.GetLogger()
	if err := godotenv.Load("/home/dev/Desktop/ordes/.env"); err != nil {
		getLogger.Fatalf("error loading .env variables: %s", err.Error())
	}
	getLogger.Infof(".env variables successfully loaded")
	postgresCfg := database.Config{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_DATABASE"),
	}

	postgresPool, err := database.ConnectionPool(context.Background(), postgresCfg)
	if err != nil {
		getLogger.Error("database connection error: %s", err)
	}
	defer postgresPool.Close()

	getLogger.Infof("Database connected successfully")
	getLogger.Debugf("Database connected successfully")

	orderNumbers, err := database.GetOrderNumbers(postgresPool)
	if err != nil {
		getLogger.Fatalf("Error getting order numbers: %v\n", err)
	}
	getLogger.Infof("Get data order numbers")

	_, shelfProducts, err := database.GetOrderProductsAndShelves(postgresPool, orderNumbers)
	if err != nil {
		getLogger.Error("Error getting order products and shelves: %v\n", err)
	}
	getLogger.Infof("Get data order products and shelves")

	logic.PrintOrder(orderNumbers, shelfProducts)

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

	}()
}
