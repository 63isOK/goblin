package main

import (
	"fmt"
	"sync"
)

func main() {
	loop()
	g()
}

func loop() {
	var l []*int
	for i := 0; i < 3; i++ {
		// i := i
		i++
		fmt.Println("i: ", i, &i)
		l = append(l, &i)
	}
	fmt.Printf("%+v\n", l)

	for _, i := range l {
		fmt.Println(*i, i)
		fmt.Println("打印i: ", &i)
	}

	fmt.Println("=====添加新变量")

	for _, i := range l {
		i := i
		fmt.Println(*i, i)
		fmt.Println("打印i: ", &i)
	}

	fmt.Println("=====再次引用")
	{
		var j []*int
		for _, i := range l {
			j = append(j, i)
		}

		for _, i := range j {
			fmt.Println(i)
		}

		fmt.Println("=====添加新变量")

		var k []*int
		for _, i := range l {
			i := i
			k = append(k, i)
		}

		for _, i := range k {
			fmt.Println(i)
		}
	}
}

func g() {
	fmt.Println("=========================")

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}

	wg.Wait()
}
