//go:build linux

package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	socketPath := filepath.Join(os.TempDir(), "sync-procs.sock")

	if len(os.Args) == 1 {
		if err := client(socketPath); err != nil {
			fmt.Printf("error in client: %v\n", err)
			os.Exit(1)
		}
		return
	}

	n, _ := strconv.Atoi(os.Args[1])
	if n <= 1 {
		fmt.Printf("num procs should be >1\n")
		os.Exit(1)
	}
	if err := server(socketPath, n); err != nil {
		fmt.Printf("error in server: %v\n", err)
		os.Exit(1)
	}
	return
}

func server(socketPath string, n int) error {
	if err := os.RemoveAll(socketPath); err != nil {
		return fmt.Errorf("deleting old path: %w", err)
	}
	l, err := net.Listen("unix", socketPath)
	if err != nil {
		return fmt.Errorf("listening: %w", err)
	}
	for i := 0; i < n-1; i++ {
		if _, err := l.Accept(); err != nil {
			return fmt.Errorf("accepting: %w", err)
		}
	}
	return nil
}

func client(socketPath string) error {
	l, err := net.Dial("unix", socketPath)
	if err != nil {
		return fmt.Errorf("dialing: %w", err)
	}
	l.Read([]byte{0})
	return nil
}
