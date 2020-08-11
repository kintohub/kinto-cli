package utils

import (
	"github.com/briandowns/spinner"
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
