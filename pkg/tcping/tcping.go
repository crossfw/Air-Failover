package tcping

import (
	"fmt"
	"github.com/cloverstd/tcping/ping"
	"time"
)

func Tcping(domain string, port int) error {
	var tcping ping.TCPing
	var target ping.Target

	target = ping.Target{
		Timeout:  time.Duration(5) * time.Second,
		Interval: 1,
		Host:     domain,
		Port:     port,
		Counter:  5,
		Protocol: 0,
	}

	tcping = *ping.NewTCPing()
	tcping.SetTarget(&target)
	pingStatus := tcping.Start()
	select {
	case <-pingStatus:
		break
	}
	result := tcping.Result()
	fmt.Printf("Rate: %v %% \n", result.SuccessCounter*100/result.Counter)

	return nil
}
