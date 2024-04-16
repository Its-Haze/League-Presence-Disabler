package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/getlantern/systray"
	"github.com/its-haze/disable-lol-rich-presence/pkg/tray"
)

func main() {
	systray.Run(tray.OnReady, tray.OnExit)
}

func init() {
	// Handle SIGINT and SIGTERM.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		systray.Quit()
	}()
}
