package deb

import (
	"fmt"
	"log"
	"runtime"
)

func Say(s string) string {
	fmt.Println(s)
	return s
}

func Check(err error) (b bool) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Printf("[error] %s:%d %v", fn, line, err)
		b = true
	}
	return
}
