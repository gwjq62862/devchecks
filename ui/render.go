package ui

import (
	"devchecks/checker"
	"fmt"
	"os"
	"text/tabwriter"
)

const (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
)

func PrintResults(results []checker.Result) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, "TYPE\tNAME\tSTATUS\tLATENCY\tMESSAGE")
	defer w.Flush()

	

	for _, value := range results {
		statusStr := Red + "[FAIL]" + Reset
		if value.Status {
			statusStr = Green + "[OK]" + Reset
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%v\t%s\n", value.Type, value.Name, statusStr, value.Latency, value.Message)
	}

}
