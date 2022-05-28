package main

import "time"

func main() {
	c := make(chan int)
	for i := 0; i < 2; i++ {
		go func(j int) {
			a := j
			for {

				time.Sleep(100 * time.Millisecond)
				c <- a
			}
		}(i)
	}

	i := <-c
	close(c)
	println(i)
}
