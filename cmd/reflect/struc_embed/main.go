package main

import "reflect"

type Mother struct {
	Name string `role:"mother"`
	Age  string `role:"a"`
}

type Fathter struct {
	Name string `role:"father"`
	Age  string `role:"a"`
}

type Son struct {
	Fathter
	Mother `yaml:"-"`
	Name   string `role:"son"`
}

func main() {
	// var f Fathter
	var s Son

	// println(p(f))
	// println(p(s))

	p(s)
}

func p(v any) string {
	t := reflect.TypeOf(v)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Anonymous {
			if len(f.Tag.Get("yaml")) == 0 {
				for i := 0; i < f.Type.NumField(); i++ {
					// println(f.Type.Field(i).Name)
					println(f.Type.Field(i).Tag.Get("role"))
				}
			}
		} else {
			println(f.Tag.Get("role"))
		}
	}

	return ""
	// return f.Tag.Get("role")
	// return reflect.TypeOf(v).Elem().Field(0).Tag.Get("role")
}
