package utils

import (
	"github.com/briandowns/spinner"
	"github.com/gookit/color"
	"time"
)

var s *spinner.Spinner

func init() {
	style := spinner.CharSets[14]
	interval := 100 * time.Millisecond
	s = spinner.New(style, interval)
	s.HideCursor = true
	s.Prefix = "Retrieving data... "
}

func StartSpinner() {
	s.Start()
}

func StopSpinner() {
	s.Stop()
}

func InfoMessage(message string) {
	StopSpinner() //To stop any active spinners, if any.
	color.Gray.Println(message)
}

func SuccessMessage(message string) {
	StopSpinner()
	color.Green.Println(message)
}

func WarningMessage(message string) {
	StopSpinner()
	color.Yellow.Println(message)
}

