// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"homework/internal"
	"homework/service"
)

//go:generate
func InitializeApp() *service.App {
	wire.Build(service.NewApp, service.NewSignalHandler, service.NewAppServer, service.NewRpcServer, internal.AppConf)
	return nil
}