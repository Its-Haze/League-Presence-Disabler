package tray

import (
	"fmt"
	"log"
	"os"

	"github.com/getlantern/systray"
	"github.com/its-haze/disable-lol-rich-presence/pkg/process"
)

func OnReady() {
	systray.SetIcon(LoadIcon("assets/leagueoflegends.ico"))
	systray.SetTitle("League RPC Checker")
	systray.SetTooltip("Disables the League Native Rich Presence plugin")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuit.ClickedCh
		log.Println("Quit menu item clicked, stopping systray.")
		systray.Quit()
	}()
	log.Println("Tray icon set up successfully.")
	go process.MonitorProcesses()
}

func LoadIcon(path string) []byte {
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Failed to load icon:", err)
		return nil
	}
	return b
}

func OnExit() {
	fmt.Println("Exiting...")
	os.Exit(0)
}
