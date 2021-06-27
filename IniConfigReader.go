package local

import (
	"bufio"
	"github.com/microdustor/common"
	"io"
	"os"
	"strings"
)

func NewIniConfiguration(file string) (*Configuration, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return iniConfigFromReader(f)
}

func iniConfigFromReader(conf io.Reader) (*Configuration, error) {
	c := &Configuration{
		data: common.New(),
	}
	scanner := bufio.NewScanner(conf)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		line := scanner.Text()
		f := strings.Split(line, "=")

		if len(f) < 1 {
			continue
		}

		switch f[0] {
		case "#":
			continue
		default:
			value := strings.Join(f[1:], "")

			c.data.Set(f[0], strings.Trim(value, "\""))
		}
	}
	return c, nil
}
