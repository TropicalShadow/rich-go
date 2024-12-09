package ipc

import (
	"fmt"
	"gopkg.in/natefinch/npipe.v2"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"
)

// getIpcPath returns the IPC path depending on the operating system.
func getIpcPath(pipe string) string {
	ipcPrefix := "discord-ipc-"
	if pipe != "" {
		ipcPrefix = fmt.Sprintf("%s%s", ipcPrefix, pipe)
	}

	var tempdir string
	var paths []string

	switch runtime.GOOS {
	case "linux", "darwin":
		tempdir = os.Getenv("XDG_RUNTIME_DIR")
		if tempdir == "" {
			tempdir = fmt.Sprintf("/run/user/%d", syscall.Getuid())
		}
		if _, err := os.Stat(tempdir); os.IsNotExist(err) {
			tempdir = os.TempDir() // Default to temporary directory if no runtime directory exists
		}

		paths = []string{".", "snap.discord", "app/com.discordapp.Discord", "app/com.discordapp.DiscordCanary"}

	case "windows":
		tempdir = `\\.\pipe`
		paths = []string{"."}
	default:
		return ""
	}

	for _, path := range paths {
		fullPath := filepath.Join(tempdir, path)
		if runtime.GOOS == "windows" || isDir(fullPath) {
			entries, err := os.ReadDir(fullPath)
			if err != nil {
				continue
			}
			for _, entry := range entries {
				if strings.HasPrefix(entry.Name(), ipcPrefix) && testIpcPath(filepath.Join(fullPath, entry.Name())) {
					return filepath.Join(fullPath, entry.Name())
				}
			}
		}
	}
	return ""
}

// testIpcPath tests if the IPC path is valid
func testIpcPath(path string) bool {
	switch runtime.GOOS {
	case "windows":
		_, err := npipe.DialTimeout(path, time.Second*2)
		return err == nil

	default:
		conn, err := net.Dial("unix", path)
		if err != nil {
			return false
		}
		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(conn)
		return true
	}
}

// isDir checks if the path is a valid directory.
func isDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}
