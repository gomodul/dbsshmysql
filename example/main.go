package main

import (
	"database/sql"
	"fmt"

	"github.com/gomodul/dbssh"
	"github.com/gomodul/dbsshmysql"
)

func main() {
	ssh, driverName, err := dbsshmysql.Open(dbssh.Config{
		Host: "127.0.0.1",
		Port: "22",
		User: "root",
		Pass: "password",
	})
	if err != nil {
		panic(err)
	}
	defer ssh.Close()

	db, err := sql.Open(driverName, "root:password@tcp(127.0.0.1:3306)/database_name?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connected to MySQL via SSH")
}
