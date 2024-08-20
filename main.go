package main

import (
	"svpnossh/cmd"
	"svpnossh/internal"
)

// init function is called before the main function
func init() {
	internal.DebugLevel()
}

func main() {
	cmd.Execute()
}
