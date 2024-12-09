//go:build windows

package ipc

import (
	npipe "gopkg.in/natefinch/npipe.v2"
	"time"
)

// OpenSocket opens the discord-ipc-0 named pipe
func OpenSocket(pipe string) error {

	pipePath := getIpcPath(pipe)
	// Connect to the Windows named pipe, this is a well known name
	// We use DialTimeout since it will block forever (or very very long) on Windows
	// if the pipe is not available (Discord not running)
	sock, err := npipe.DialTimeout(pipePath, time.Second*2)
	if err != nil {
		return err
	}

	Socket = sock
	return nil
}
