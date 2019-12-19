package main

import "github.com/google/wire"

var BazSuperSet = wire.NewSet(
	NewFoo,
	NewBar,
	NewBaz,
)

type Foo struct {
	X int
}

func NewFoo() Foo {
	return Foo{X: 42}
}

type Bar struct {
	Foo Foo
}

func NewBar(foo Foo) Bar {
	return Bar{Foo: foo}
}

type Baz struct {
	Bar Bar
}

func NewBaz(bar Bar) Baz {
	return Baz{Bar: bar}
}
