package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

type GreetingJob struct {
	Name string
}

func (g GreetingJob) Run() {
	fmt.Println("Hello ", g.Name)
}

func main() {
	a := cron.New()

	a.AddFunc("@every 6s", func() {
		fmt.Println(time.Now())
	},
	)

	//b := cron.New()
	//b.AddFunc("12-59/1 * * * *", func() {
	//	fmt.Println("******************")
	//	fmt.Println(time.Now())
	//	fmt.Println("******************")
	//})

	//c := cron.New()
	//c.AddJob("@every 1s", GreetingJob{Name: "jason"})
	//

	//parser := cron.NewParser(
	//	cron.Second | cron.Minute | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	//)

	//d := cron.New(cron.WithParser(parser))
	//_,err := d.AddFunc("1-59/1 * * * * *", func() {
	//	fmt.Println("d_func!")
	//})


	d := cron.New(
		cron.WithParser(
			cron.NewParser(
				cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
	_,err := d.AddFunc("1-59/1 * * * * *", func() {
		fmt.Println("d_func!")
	})
	if err != nil{
		fmt.Println(err)
	}


	//c.Start()
	//b.Start()
	a.Start()
	d.Start()

	time.Sleep(time.Hour * 1)

}
