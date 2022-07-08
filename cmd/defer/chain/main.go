package main

type Obj struct {
	Hello string
}

func (o *Obj) Start() {
	println("start")
}

func (o *Obj) Stop() {
	println("stop")
}

func Start(o *Obj) interface{ Stop() } {
	o.Start()
	return o
}

func main() {
	o := &Obj{
		Hello: "tom",
	}
	defer Start(o).Stop()

	println(o.Hello)
}
