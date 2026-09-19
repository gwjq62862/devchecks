package checker

import (
	"os"
	"time"
)

func CheckEnv(name string) Result {
	start := time.Now()
	val, exists := os.LookupEnv(name)
	latency := time.Since(start)

	if !exists {
		return Result{
			Status:  false,
			Type:    "ENV",
			Message: "environment variable is not set",
			Latency: latency,
		}
	}
	if exists && val == ""{
         return Result{
			Status: false,
			Type: "ENV",
			Message: "environment variable is empty",
			Latency: latency,
		 }
	}
     return Result{
		Status: true,
		Type: "ENV",
		Message: "OK",
		Latency: latency,
	 }


}
