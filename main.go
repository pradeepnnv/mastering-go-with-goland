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

	for _, eName := range []string{
		"USER",
		"PASSWORD",
		"HOST",
		"NAME",
	} {
		if !viper.IsSet(eName) {
			log.Fatalf("Environment value DB_%s is not set", eName)
		}
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
