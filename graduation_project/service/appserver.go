package service

import (
	"context"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"homework/internal"
	"net/http"
	"strconv"
)

// AppServer app server
type AppServer struct {
	ctx context.Context
	srv *http.Server
}

// Serve start server
func (appServer *AppServer) Serve(ctx context.Context) error {
	appServer.ctx = ctx
	return appServer.srv.ListenAndServe()
}
func (appServer *AppServer) Shutdown() error {
	return appServer.srv.Shutdown(appServer.ctx)
}

// NewAppServer new app server
func NewAppServer(config *internal.Config) *AppServer {
	// http服务器
	srv := &http.Server{Addr: ":" + strconv.Itoa(config.App.Port)}
	
	// 注册handler
	registerHandlers()

	return &AppServer{
		srv: srv,
	}
}

// handler
func registerHandlers() {
	//http.HandleFunc("/", handlers.SayHello)
	// Prometheus
	http.Handle("/",  promhttp.Handler())
}