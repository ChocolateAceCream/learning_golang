package recoverAndPanic

import (
	"errors"
	"fmt"
)

func Demo() (err error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
			err = errors.New("new error") // if you panic from inside recover(), then it will panic again. otherwise, recover() will always return nil
		}
	}()
	panic("panic error")
}
