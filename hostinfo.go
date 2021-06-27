package local

import (
	"bufio"
	"errors"
	"os"
)

type HostInfo struct {
	Err      error
	HostName string
	Ip       string
}

var Host = new(HostInfo)

func init() {
	Reload()
}

func Reload() {
	f, err := os.Open("/etc/hostinfo")
	if err != nil {
		Host.Err = err
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	if scanner.Err() != nil {
		Host.Err = scanner.Err()
		return
	}
	scanner.Scan()
	Host.HostName = scanner.Text()
	if len(Host.HostName) < 1 {
		Host.Err = errors.New("host name length error")
	}
	scanner.Scan()
	Host.Ip = scanner.Text()
	if len(Host.Ip) < 7 {
		Host.Err = errors.New("ip length error")
	}
}
