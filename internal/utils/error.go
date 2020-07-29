package utils

import (
	"os"
	"github.com/gookit/color"
)

func TerminateWithError(err error) {
	color.Red.Printf("\nan error occurred %v\n", err)
	os.Exit(1)
}
