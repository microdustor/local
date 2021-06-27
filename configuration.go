package local

import (
	"fmt"
)

type SimpleConfiguration struct {
	*Configuration
	Err error
}

var Daf SimpleConfiguration

func init() {
	f, err := NewIniConfiguration("/usr/configuragion/spine/gtm.conf")
	Daf.Err = err
	Daf.Configuration = f
}

func (sc *SimpleConfiguration) Get(key string) (string, bool) {
	if sc.Err != nil {
		panic(fmt.Sprintf("something error occur when query value from simple configuration:%v", sc.Err))
	}
	v, found := sc.Configuration.Get(key)
	if v != nil {
		return v.(string), found
	} else {
		return "", found
	}
}

//must check first
func (sc *SimpleConfiguration) Obtain(key string) string {
	v, _ := sc.Get(key)
	return v
}

func (sc *SimpleConfiguration) Check(keys []string) error {
	for _, key := range keys {
		if _, found := sc.Get(key); !found {
			return fmt.Errorf("key:%s not found", key)
		}
	}
	return nil
}
