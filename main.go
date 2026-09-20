package main

import (
	"devchecks/checker"
	"devchecks/config"
	"devchecks/ui"
	"fmt"
	"os"
)



func main() {

cfg, err := config.LoadConfig("devcheck.yaml")

 if err != nil {
    fmt.Printf("Config error: %v\n", err)
    os.Exit(1)
}
	results := checker.RunChecks(cfg)

	ui.PrintResults(results)

	
	for _, res := range results {
		if !res.Status {
			os.Exit(1)
		}
	}
}