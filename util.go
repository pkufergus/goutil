package goutil

import (
	"time"
	"runtime"
	"fmt"
)

func TimeCost(logfunc interface{}, logid uint32) func() {
	start := time.Now()
	pc,file,line,_ := runtime.Caller(1)
	switch logfunc.(type) {
	case func(f interface{}, v ...interface{}):
		f := logfunc.(func(f interface{}, v ...interface{}))
		f("[%s:%d %s][logid:%d]enter ", file, line, runtime.FuncForPC(pc).Name(), logid)
	case func(s string, a ...interface{}):
		f := logfunc.(func(s string, a ...interface{}))
		f("[%s:%d %s][logid:%d]enter ", file, line, runtime.FuncForPC(pc).Name(), logid)
	default:
		f := fmt.Printf
		f("[%s:%d %s][logid:%d]enter\n", file, line, runtime.FuncForPC(pc).Name(), logid)
	}
	return func() {
		pc,file,line,_ := runtime.Caller(1)
		switch logfunc.(type) {
		case func(f interface{}, v ...interface{}):
			f := logfunc.(func(f interface{}, v ...interface{}))
			f("[%s:%d %s][logid:%d]exit cost:%s", file, line, runtime.FuncForPC(pc).Name(), logid, time.Since(start))
		case func(s string, a ...interface{}):
			f := logfunc.(func(s string, a ...interface{}))
			f("[%s:%d %s][logid:%d]exit cost:%s", file, line, runtime.FuncForPC(pc).Name(), logid, time.Since(start))
		default:
			f := fmt.Printf
			f("[%s:%d %s][logid:%d]exit cost:%s\n", file, line, runtime.FuncForPC(pc).Name(), logid, time.Since(start))
		}
	}
}

func GetDateFromTimestamp(sec int64) {
	tm := time.Unix(sec, 0)
	fmt.Println(tm.Format("20060102"))
}

