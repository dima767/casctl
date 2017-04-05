package cmd

import (
	"net/http"
	"os"
)

func erAndExit(msg interface{}) {
	redPrintln("casctl:", msg)
	os.Exit(-1)
}

func infoAndExit(msg interface{}) {
	cyanPrintln("casctl:", msg)
	os.Exit(0)
}

func checkResponseAndExitIfNecessary(resp *http.Response) {
	if resp == nil || resp.StatusCode == 404 {
		erAndExit("Invalid or unreachable CAS server.")
	}
}
