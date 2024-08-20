package internal

import (
	"fmt"
	"log"
	"os"
)

// DebugLevel is a function that sets the log level to debug, showing the file name and line number
// if the DEBUG environment variable is set to true
func DebugLevel() {
	if os.Getenv("DEBUG") == "true" {

		//log with file name and line number
		log.SetFlags(log.LstdFlags | log.Lshortfile)

		//color yellow
		fmt.Print("\033[33m")
		fmt.Println("DEBUG mode is enabled")
		//reset color
		fmt.Print("\033[0m")
	}
}
