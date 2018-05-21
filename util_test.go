package goutil

import (
	"testing"
	"fmt"
	"time"
)

func TestTimeCost(t *testing.T) {
	fmt.Println("start test")
	defer TimeCost(fmt.Println, 12321)()
	time.Sleep(time.Second)
}

func TestGetDateFromTimestamp(t *testing.T) {
	GetDateFromTimestamp(1526449519)

	the_time := time.Date(2018, 3, 21, 0, 0, 0, 0, time.Local)
	fmt.Println(the_time.Format("20060102"))
	d, _ := time.ParseDuration("2400h")
	new_time := the_time.Add(d)
	fmt.Println(new_time.Format("20060102"))
}


