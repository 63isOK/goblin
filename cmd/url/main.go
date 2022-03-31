package main

import "fmt"

func main() {
	header := "http://iot.com"
	body := "hello"

	var i int
	if header[len(header)-1:] == "/" {
		i++
	}
	if body[:1] == "/" {
		i++
	}

	var url string
	switch i {
	case 0:
		url = fmt.Sprint(header, "/", body)
	case 2:
		url = fmt.Sprint(header[:len(header)-1], body)
	case 1:
		url = fmt.Sprint(header, body)
	}

	fmt.Println(url)
}
