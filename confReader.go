package local

import (
	"bufio"
	"github.com/microdustor/common"
	"io"
	"os"
	"strings"
)

func NewConfiguration(file string) (*Configuration, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return configFromReader(f)
}

func configFromReader(conf io.Reader) (*Configuration, error) {
	c := &Configuration{
		data: common.New(),
	}
	scanner := bufio.NewScanner(conf)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		line := scanner.Text()
		f := strings.Fields(line)

		if len(f) < 1 {
			continue
		}

		switch f[0] {
		case "#":
			continue
		default:
			c.data.Set(f[0], f[1:])
		}
	}
	return c, nil
}

func (c *Configuration) Get(key string) (interface{}, bool) {
	return c.data.Get(key)
}
