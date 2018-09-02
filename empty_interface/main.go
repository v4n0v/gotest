package empty_interface

import (
	"fmt"
	"runtime"
)

func tstPanic(err error) {
	var ok bool
	defer func() {
		if r := recover(); r != nil {
			if err, ok = r.(error); ok {
				fmt.Println(err)
			}
			fmt.Println("Panic defered")
			buf := make([]byte, 10000)
			runtime.Stack(buf, false)
			fmt.Println("Stack Trace:", string(buf))
			if err != nil {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
		fmt.Println("start test")
		panic ("PANIC!!!")

}

func main() {

}
