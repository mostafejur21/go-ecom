package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/mostafejur21/go-ecom/cmd/api"
	"github.com/mostafejur21/go-ecom/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 "root",
		Passwd:               "asd",
		Addr:                 "127.0.1:3306",
		DBName:               "ecom",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	server := api.NewAPIServe(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
