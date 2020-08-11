package utils

import (
	"github.com/gookit/color"
)

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
