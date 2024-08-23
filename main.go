package main

import (
	"fmt"
	"os"
	"os/signal"
	"svpnossh/cmd"
	"svpnossh/internal"
	"syscall"
)

// init function is called before the main function
func init() {

	// set the log level to debug
	internal.DebugLevel()
}

func main() {

	// SIGINT o SIGTERM (ctrl + c)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs

		fmt.Println("Signal received: ", sig)
		if sig == syscall.SIGINT || sig == syscall.SIGTERM {
			internal.Cleanup()
			os.Exit(0)
		}
	}()
	cmd.Execute()
}
