package main

import (
	"github.com/xueqiya/go_project/model"
	"log"
	"net/http"
	"time"

	"github.com/xueqiya/go_project/config"
	"github.com/xueqiya/go_project/router"
)

var sc = config.Cfg.Server

func main() {
	model.Setup()
	defer model.Close()

	server := &http.Server{
		Addr:           sc.Addr,
		Handler:        router.Setup(),
		ReadTimeout:    time.Duration(sc.ReadTimeout * int(time.Second)), // 转换成时间数据结构
		WriteTimeout:   time.Duration(sc.WriteTimeout * int(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}
