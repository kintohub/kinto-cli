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
	color.Gray.Println("INFO    ",message)
}

func NoteMessage(message string) {
	StopSpinner() //To stop any active spinners, if any.
	color.Magenta.Println("NOTE    ",message)
}
func SuccessMessage(message string) {
	StopSpinner()
	//color.BgGreen.Printf("SUCCESS:")
	color.Green.Println("SUCCESS ",message)
}

func WarningMessage(message string) {
	StopSpinner()
	//color.BgYellow.Printf("WARNING:")
	color.Yellow.Println("WARNING ",message)
}

