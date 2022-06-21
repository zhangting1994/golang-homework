package service

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func NewApp(appServer *AppServer, rpcServer *RpcServer, sigHandler *SignalHandler) *App {
	return &App{
		AppServer: appServer,
		RpcServer: rpcServer,
		SignalHandler: sigHandler,
	}
}

type App struct {
	AppServer *AppServer
	RpcServer *RpcServer
	SignalHandler *SignalHandler
}

func (app *App) Start() {
	// 创建带上下文的error group
	g, errCtx := errgroup.WithContext(context.Background())

	// 注册关闭函数
	app.SignalHandler.ShutdownFuncReg(func() error {
		_ = app.RpcServer.Shutdown()
		return app.AppServer.Shutdown()
	})

	// error group 启动
	g.Go(func() error {
		// web server 以及 Prometheus指标
		return app.AppServer.Serve(errCtx)
	})
	g.Go(func() error {
		// rpc server
		return app.RpcServer.Serve(errCtx)
	})
	g.Go(func() error {
		// signal handler
		return app.SignalHandler.Handle(errCtx)
	})

	// 等待错误抛出
	if err := g.Wait(); err != nil {
		fmt.Println("服务已退出，原因:" + err.Error())
	}
}