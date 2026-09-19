package checker

import "time"

type Result struct {
	Name    string
	Type    string
	Status  bool
	Message string
	Latency time.Duration
	
}