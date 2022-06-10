package main

import "github.com/robfig/cron/v3"

func main() {
	c := cron.New()
	c.AddFunc("@every 1s", func() { println("1") })
	c.Start()

	c.AddFunc("@every 1s", func() { println("2") })
	c.AddFunc("@every 1s", func() { println("3") })
	c.AddFunc("@every 1s", func() { println("4") })
	c.AddFunc("@every 1s", func() { println("5") })
	c.AddFunc("@every 1s", func() { println("6") })

	select {}
}
