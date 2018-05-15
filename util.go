package goutil

import (
	"time"
	"runtime"
	"fmt"
)

func TimeCost(logfunc interface{}) func() {
	start := time.Now()
	pc,file,line,_ := runtime.Caller(1)
	switch logfunc.(type) {
	case func(f interface{}, v ...interface{}):
		f := logfunc.(func(f interface{}, v ...interface{}))
		f("[%s:%d %s]enter ", file, line, runtime.FuncForPC(pc).Name())
	case func(s string, a ...interface{}):
		f := logfunc.(func(s string, a ...interface{}))
		f("[%s:%d %s]enter ", file, line, runtime.FuncForPC(pc).Name())
	default:
		f := fmt.Printf
		f("[%s:%d %s]enter\n", file, line, runtime.FuncForPC(pc).Name())
	}
	return func() {
		pc,file,line,_ := runtime.Caller(1)
		switch logfunc.(type) {
		case func(f interface{}, v ...interface{}):
			f := logfunc.(func(f interface{}, v ...interface{}))
			f("[%s:%d %s]exit cost:%s", file, line, runtime.FuncForPC(pc).Name(), time.Since(start))
		case func(s string, a ...interface{}):
			f := logfunc.(func(s string, a ...interface{}))
			f("[%s:%d %s]exit cost:%s", file, line, runtime.FuncForPC(pc).Name(), time.Since(start))
		default:
			f := fmt.Printf
			f("[%s:%d %s]exit cost:%s\n", file, line, runtime.FuncForPC(pc).Name(), time.Since(start))
		}
	}
}

