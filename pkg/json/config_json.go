package json

import (
	"encoding/json"
	"github.com/crossfw/air-failover/pkg/log"
	"os"
)

type config struct {
	Global struct {
		CheckRate       int    `json:"check_rate"`
		CloudflareEmail string `json:"cloudflare_email"`
		CloudflareToken string `json:"cloudflare_token"`
	} `json:"global"`
	RelayServers []struct {
		Tag            string `json:"tag"`
		MainDomain     string `json:"main_domain"`
		FailoverDomain string `json:"failover_domain"`
		MonitorPorts   []int  `json:"monitor_ports"`
	} `json:"relay_servers"`
	FailoverRule []struct {
		Tag           string   `json:"tag"`
		FailoverNodes []string `json:"failover_nodes"`
	} `json:"failover_rule"`
}


func ReadConfig( configPath *string) (*config, error) {
	var cfg config
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