package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/korvised/kafka-producer/handlers"
	"github.com/korvised/kafka-producer/services"
	"github.com/spf13/viper"
	"strings"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	producer, err := sarama.NewSyncProducer(viper.GetStringSlice("kafka.servers"), nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	eventProducer := services.NewEventProducer(producer)
	accountService := services.NewAccountService(eventProducer)
	accountHandler := handlers.NewAccountHandler(accountService)

	app := fiber.New()

	router := app.Group("api/v1")

	router.Post("/open-account", accountHandler.OpenAccount)
	router.Patch("/deposit-fund/:cus_id", accountHandler.DepositFund)
	router.Patch("/withdraw-fund/:cus_id", accountHandler.WithdrawFund)
	router.Delete("/close-account/:cus_id", accountHandler.CloseAccount)

	fmt.Println("Producer server started 🔥🔥🔥 ")

	if err = app.Listen(viper.GetString("app.url")); err != nil {
		panic(err)
	}
}
