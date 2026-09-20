package checker

import (
	"net"
	"time"
)

func CheckHost(address string) Result {
	timeout := 2 * time.Second
	start := time.Now()
	con, err := net.DialTimeout("tcp", address, timeout)
	latency := time.Since(start)

	if err != nil {
		return Result{
			Name:    address,
			Status:  false,
			Type:    "HOST",
			Message: err.Error(),
			Latency: latency,
		}
	}

	defer con.Close()
	return Result{
		Name:    address,
		Status:  true,
		Type:    "HOST",
		Message: "OK",
		Latency: latency,
	}

}
