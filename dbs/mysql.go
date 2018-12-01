package dbs

import (
	"../utils"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Conns *sql.DB

func init() {
	var err error
	username := utils.Conf.ReadStr("mysql", "username")
	password := utils.Conf.ReadStr("mysql", "password")
	dataname := utils.Conf.ReadStr("mysql", "dataname")
	port := utils.Conf.ReadStr("mysql", "port")
	host := utils.Conf.ReadStr("mysql", "host")

	dns := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dataname + "?parseTime=true"
	fmt.Println(dns)
	Conns, err = sql.Open("mysql", dns)
	if err != nil {
		fmt.Printf("dddddddddddddddddddddd====%s", err.Error())
		log.Fatal(err.Error())

	}

	err = Conns.Ping()
	if err != nil {
		fmt.Printf("xxxxxxxxxxxxxxxx====%s", err.Error())
		log.Fatal(err.Error())
	}
	Conns.SetMaxIdleConns(20)
	Conns.SetMaxOpenConns(20)
}
