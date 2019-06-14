package main

import (
	"testing"
	"flag"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"context"
	"os"
)

var integration = flag.Bool("integration", false, "run database integration tests")
var db *sql.DB

func TestMain(m *testing.M) {
	flag.Parse()

	if *integration {
		log.Println("Running Integration Test")
		setupDatabase()
	} else {
		log.Println("Not running Integration Test")
	}

	code := m.Run()

	if *integration {
		tearDown()
	}

	os.Exit(code)
}

func setupDatabase() {
	var err error
	user := os.Getenv("MYSQL_USER")
	host := os.Getenv("MYSQL_HOST")
	database := os.Getenv("MYSQL_DB")

	databaseUri := user+"@tcp("+host+")/"+database

	db, err = sql.Open("mysql", databaseUri)
	if err !=nil{
		log.Fatal(err)
	}
}

func tearDown()  {
	db.Close()
}

func Test_GetUserByID(t *testing.T) {
	if !*integration{
		t.SkipNow()
	}

	data, err := GetUserByID(context.Background(), 1,db)
	if err !=nil{
		t.Error(err)
	}else{
		t.Log(data)
	}
}
