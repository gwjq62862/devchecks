package checker

import (
	"devchecks/config"
	"sync"
	"time"
)

type Result struct {
	Name    string
	Type    string
	Status  bool
	Message string
	Latency time.Duration
}

func RunChecks(cfg *config.DevConfig) []Result {
	resultChannel := make(chan Result)
	var wg sync.WaitGroup

	var result []Result

	for _, env := range cfg.Envs {
		wg.Add(1)
		go func(name string) {
			res := CheckEnv(name)
			resultChannel <- res
			wg.Done()
		}(env)
	}

	for _, host := range cfg.Hosts {
		wg.Add(1)
		go func(name string) {
			res := CheckHost(host)
			resultChannel <- res
			wg.Done()

		}(host)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()
	for res := range resultChannel {
		result = append(result, res)
	}
	return result

}
