package goutil

import (
	"testing"
	"fmt"
	"time"
)

func TestTimeCost(t *testing.T) {
	fmt.Println("start test")
	defer TimeCost(fmt.Println)()
	time.Sleep(time.Second)
}


