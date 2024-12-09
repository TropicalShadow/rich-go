//go:build !windows

package ipc

import (
	"net"
	"time"
)

// OpenSocket opens the discord-ipc-0 unix Socket
func OpenSocket(pipe string) error {
	path := getIpcPath(pipe)

	sock, err := net.DialTimeout("unix", path, time.Second*2)
	if err != nil {
		return err
	}

	Socket = sock
	return nil
}
