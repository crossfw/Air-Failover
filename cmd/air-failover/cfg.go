package main

import (
	"encoding/json"
	"github.com/crossfw/Air-Failover/pkg/log"
	"os"
)

type baseConfig struct {
	Log               Log               `json:"log"`
	Global            Global            `json:"global"`
	Nodes             []Nodes           `json:"nodes"`
	FailoverRule      []FailoverRule    `json:"failover_rule"`
	FailoverCondition FailoverCondition `json:"failover_condition"`
}

type Log struct {
	LogLevel string `json:"log_level"`
}

type Global struct {
	CheckRate       int    `json:"check_rate"`
	CloudflareEmail string `json:"cloudflare_email"`
	CloudflareToken string `json:"cloudflare_token"`
}

type Nodes struct {
	Tag            string `json:"tag"`
	MainDomain     string `json:"main_domain"`
	FailoverDomain string `json:"failover_domain"`
	MonitorPorts   []int  `json:"monitor_ports"`
}

type FailoverRule struct {
	Tag         string   `json:"tag"`
	BackupNodes []string `json:"backup_nodes"`
}

type Check struct {
	Count          int     `json:"count"`
	IntervalTime   int     `json:"interval_time"`
	MinSuccessRate float64 `json:"min_success_rate"`
	MaxPingTime    int     `json:"max_ping_time"`
}

type FailoverOut struct {
	FailTimes int `json:"fail_times"`
}

type FailoverBack struct {
	SuccessTimes int `json:"success_times"`
}

type FailoverCondition struct {
	Check        Check        `json:"check"`
	FailoverOut  FailoverOut  `json:"failover_out"`
	FailoverBack FailoverBack `json:"failover_back"`
}

func parseBaseConfig(configPath *string) (*baseConfig, error) {
	var cfg baseConfig
	file, err := os.Open(*configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Set default config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}

	log.Debug.Println(cfg)
	return &cfg, nil
}
