package global

import (
	"fmt"
	"github.com/spf13/cast"
	"time"
)

const LAYOUT = "2006-01-02 15:04:05"


func ClientRun(handler func())  {
	for {
		t := time.NewTimer(time.Second * 1)
		select {
		case <-t.C:
			handler()
		}
	}
}

func ServerRun(handler func())  {
	go handler()
	select {

	}
}

func PrintTime()  {
	fmt.Println(cast.ToString(time.Now().Format(LAYOUT)))
}