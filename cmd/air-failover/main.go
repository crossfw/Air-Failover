package main

import (
	"github.com/crossfw/air-failover/pkg/controller"
	"github.com/crossfw/air-failover/pkg/json"
	"github.com/crossfw/air-failover/pkg/log"
	"os"
)

func main() {
	log.Debug.Println("Started")
	path := "config/example.json"
	cfg, err := json.ReadConfig(&path)
	for err != nil {
		log.Warn.Println(err)
	}
	controller.Controller(cfg)
	os.Exit(1)
}
