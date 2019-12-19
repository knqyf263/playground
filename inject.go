//+build wireinject

package main

import "github.com/google/wire"

func InitializeEvent() Event {
	wire.Build(SuperSet)
	return Event{}
}
