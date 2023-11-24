package main

import (
	"context"
	"fmt"
	"gihub.com/korvised/kafka-consumer/repositories"
	"gihub.com/korvised/kafka-consumer/services"
	"github.com/IBM/sarama"
	events "github.com/korvised/kafka-events"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

func initDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)

	dial := mysql.Open(dsn)

	db, err := gorm.Open(dial, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	consumer, err := sarama.NewConsumerGroup(viper.GetStringSlice("kafka.servers"), viper.GetString("kafka.group"), nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	db := initDatabase()
	accountRepo := repositories.NewAccountRepository(db)
	accountEventHandler := services.NewAccountHandler(accountRepo)
	accountConsumerHandler := services.NewConsumerHandler(accountEventHandler)
	fmt.Println("Account consumer started...")

	for {
		if err = consumer.Consume(context.Background(), events.Topics, accountConsumerHandler); err != nil {
			panic(err)
		}
	}
}
