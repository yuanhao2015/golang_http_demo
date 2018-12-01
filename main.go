package main

import (
	"./routers"
	"./utils"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	server := http.Server{
		Addr:           fmt.Sprintf(":%d", utils.Conf.ReadInt32("website", "httpport")),
		Handler:        &routers.HttpHandler{},
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 0,
	}
	log.Println(fmt.Sprintf("Listen: %d", utils.Conf.ReadInt32("website", "httpport")))
	log.Fatal(server.ListenAndServe())
}
