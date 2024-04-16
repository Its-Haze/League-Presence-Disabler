package process

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/its-haze/disable-lol-rich-presence/pkg/config"
	"github.com/mitchellh/go-ps"
	"golang.org/x/sys/windows"
)

func MonitorProcesses() {
	for {
		gamePath, err := findGamePath()
		if err != nil {
			fmt.Println("Error finding game path:", err)
			time.Sleep(10 * time.Second)
			continue
		}
		if gamePath != "" {
			config.ModifyJSON(gamePath)
		}
		time.Sleep(2 * time.Second)
	}
}

func findGamePath() (string, error) {
	processes, err := ps.Processes()
	if err != nil {
		log.Printf("Error listing processes: %v\n", err)
		return "", err
	}
	for _, p := range processes {
		name := p.Executable()
		if strings.Contains(name, "RiotClientServices") {
			exePath, err := executablePath(p.Pid())
			if err != nil {
				log.Printf("Error retrieving executable path for PID %d: %v\n", p.Pid(), err)
				return "", err
			}
			if strings.Contains(exePath, "Riot Games") {
				baseRiotPath := strings.Split(exePath, "Riot Games")[0] + "Riot Games"
				fullPath := filepath.Join(baseRiotPath, "League of Legends", "Plugins", "plugin-manifest.json")
				log.Printf("Found game path: %s\n", fullPath)
				return fullPath, nil
			}
		}
	}
	log.Println("RiotClientServices process not found.")
	return "", nil
}

func executablePath(pid int) (string, error) {

	h, err := windows.OpenProcess(windows.PROCESS_QUERY_INFORMATION|windows.PROCESS_VM_READ, false, uint32(pid))
	if err != nil {
		return "", err
	}
	defer windows.CloseHandle(h)

	var buf [windows.MAX_PATH]uint16
	var size uint32 = windows.MAX_PATH
	err = windows.QueryFullProcessImageName(h, 0, &buf[0], &size)
	if err != nil {
		return "", err
	}

	// Convert from UTF16 to a normal Go string
	executablePath := windows.UTF16ToString(buf[:size])
	return executablePath, nil
}
