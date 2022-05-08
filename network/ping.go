package network

import (
	"net"
	"time"

	"github.com/tatsushid/go-fastping"
)

var duration time.Duration

func Handle(hostname string) (int64, error) {
	p := fastping.NewPinger()

	ra, err := net.ResolveIPAddr("ip4:icmp", hostname)
	if err != nil {
		return 999, err
	}

	p.AddIPAddr(ra)

	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		duration = rtt
	}

	err = p.Run()
	if err != nil {
		return 999, err
	}

	miliseconds := duration.Milliseconds()
	return miliseconds, nil
}
