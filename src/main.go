package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	//"migrations"

	_ "github.com/go-sql-driver/mysql"
)

type Row struct {
	gorm.Model
	Code  string
	Price uint
}

func dbConnect(some int) error {
	_, err := gorm.Open("mysql", "root:secret@/goland")

	return err
}

// var schema string = "CREATE TABLE `users` (	  	`id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,		  	`name` varchar(255) NOT NULL		)"

func main() {

	dbConfig := mysql.NewConfig()
	dbConfig.User = "root"
	dbConfig.Passwd = "secret"
	dbConfig.Addr = "mysql-container:3306"
	dbConfig.DBName = "goland"
	dbConfig.Net = "tcp"

	db, err := gorm.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//db, err := gorm.Open("mysql", "root:secret@tcp(localhost:3306)/goland")
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer db.Close()

	//err := retry.Sleep(3, 10000000000, dbConnect)
	//if err != nil {
	//	log.Fatal(err)
	//}

	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}

	db.AutoMigrate(&Row{})

	// Create
	db.Create(&Row{Code: "L1212", Price: 1000})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test: %s\n", r.URL.Path)
	})
	//http.HandleFunc("/path", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, dbConfig.FormatDSN(), nil)
	//})
	http.ListenAndServe(":"+PORT, nil)
}
