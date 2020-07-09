package utils

import (
	"fmt"
	"os"
)

func TerminateWithError(err error) {
	fmt.Printf("an error occurred %v", err)
	os.Exit(1)
}
