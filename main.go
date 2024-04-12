package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/min23asdw/go_api_learning/pkg/api"
	"github.com/min23asdw/go_api_learning/pkg/store"
)

func main() {
	log.Println("Starting")
	cfg := mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPassword,
		Addr:                 Envs.DBAddress,
		DBName:               Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	sqlStorage := NewMySQLStorage(cfg)

	db, err := sqlStorage.Init()

	if err != nil {
		log.Fatal(err)
	}

	s := store.NewStore(db)
	api := api.NewAPIServer(":3000", s)
	api.Serve()
}
