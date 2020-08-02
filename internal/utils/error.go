package utils

import (
	"github.com/gookit/color"
	"os"
)

func TerminateWithError(err error) {
	color.Red.Printf("\nAn error occurred %v\n", err)
	os.Exit(1)
}

func TerminateWithCustomError(message string) {
	color.Red.Printf("\n" + message + "\n")
	os.Exit(1)
}
