package main

import (
	"context"
	"time"
)

func main() {
	ctx, ctxfun := context.WithTimeout(context.Background(), time.Second)
	defer ctxfun()
	<-ctx.Done()
}
