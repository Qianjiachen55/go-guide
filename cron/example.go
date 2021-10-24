package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New()

	c.AddFunc("@every 10s", func() {
		fmt.Println(time.Now())
	},
	)

	a := cron.New()
	a.AddFunc("12-59/1 * * * *", func() {
		fmt.Println("******************")
		fmt.Println(time.Now())
		fmt.Println("******************")
	})
	a.Start()


	c.Start()

}
