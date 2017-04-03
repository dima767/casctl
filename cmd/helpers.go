package cmd

import (
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
