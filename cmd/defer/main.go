package main

import "sync"

func main() {
	test()
	println("===========")
	testLock()
	println("============ done ===========")
}

// defer 支持套娃
func test() {
	defer func() {
		defer func() {
			println(2)
		}()
		println(1)
	}()

	println(3)
}

var l sync.Mutex

func testLock() {
	l.Lock()
	defer l.Unlock()

	// defer里再也获取不到锁了
	defer func() {
		l.Lock()
		defer l.Unlock()

		println("defer已加锁")
	}()

	println("已加锁")
}
