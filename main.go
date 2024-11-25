package main

import (
	"github.com/spf13/viper"
	"log"
	"my-first-api/internal/db"
	"my-first-api/internal/todo"
	"my-first-api/internal/transport"
)

func main() {
	d := InitializeDb()
	svc := todo.NewService(d)
	server := transport.NewServer(svc)
	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}

func InitializeDb() *db.DB {
	viper.SetEnvPrefix("DB")
	viper.AutomaticEnv()
	viper.SetDefault("PORT", "5432")

	if !viper.IsSet("USER") || !viper.IsSet("PASSWORD") || !viper.IsSet("HOST") || !viper.IsSet("NAME") {
		log.Fatal("DB Details not provided")
	}

	d, err := db.New(
		viper.GetString("USER"),
		viper.GetString("PASSWORD"),
		viper.GetString("HOST"),
		viper.GetInt("PORT"),
		viper.GetString("NAME"),
	)
	if err != nil {
		log.Fatal(err)
	}
	return d
}
