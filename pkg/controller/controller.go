package controller

import (
	"github.com/crossfw/air-failover/pkg/json"
	"github.com/crossfw/air-failover/pkg/log"
	"github.com/crossfw/air-failover/pkg/tcping"
)

func Controller(cfg *json.Config) {
	log.Debug.Println("Start")
	err := tcping.Tcping(cfg.RelayServers[0].MainDomain, cfg.RelayServers[0].MonitorPorts[0])
	for err != nil {
		log.Warn.Println(err)
	}
}
