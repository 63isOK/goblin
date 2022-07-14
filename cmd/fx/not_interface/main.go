package main

import (
	"context"
	"log"
	"time"

	"go.uber.org/fx"
)

type A struct {
}

func NewA(b *B) *A {
	println("a")
	return &b.AO
}

type B struct {
	AO A
}

func NewB(c *C) *B {
	println("b")
	return &c.BO
}

type C struct {
	BO B
}

func NewC() *C {
	println("c")
	return &C{}
}

func p(a *A) {
	println("开始")
}

func main() {
	app := fx.New(
		fx.Provide(
			NewA,
			NewB,
			NewC,
		),
		fx.Invoke(p),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}
