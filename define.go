package local

import (
	"github.com/microdustor/common"
	"sync"
)

type Configuration struct {
	path string
	data common.FixedMap
}

type ReadAndWriteConfiguration struct {
	path string
	lock sync.RWMutex
	split byte
	eol   byte
	Data *common.SimpleFile
	Err error
}