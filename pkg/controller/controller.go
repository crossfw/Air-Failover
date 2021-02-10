package controller

import (
	"github.com/crossfw/Air-Failover/pkg/json"
	"github.com/crossfw/Air-Failover/pkg/log"
	"github.com/crossfw/Air-Failover/pkg/tcping"
)

func Controller(cfg *json.Config) {
	log.Debug.Println("Start")
	err := tcping.Tcping(cfg.RelayServers[0].MainDomain, cfg.RelayServers[0].MonitorPorts[0])
	for err != nil {
		log.Warn.Println(err)
	}
}
