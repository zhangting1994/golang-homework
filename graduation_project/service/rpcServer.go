package service

import (
	"context"
	"google.golang.org/grpc"
	"homework/internal"
	"homework/protocol"
	"homework/service/rpcHandler"
	"log"
	"net"
	"strconv"
)

// RpcServer rpc server
type RpcServer struct {
	ctx    context.Context
	config internal.Config
	server *grpc.Server
}

// NewRpcServer new rpc server
func NewRpcServer(config *internal.Config) *RpcServer {
	r := &RpcServer{}
	r.config = *config
	return r
}

func (srv *RpcServer) Serve(ctx context.Context) error {
	srv.ctx = ctx

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(srv.config.Rpc.Port))
	if err != nil {
		log.Fatal(err)
		return err
	}

	// 启动grpc服务
	srv.server = grpc.NewServer()
	protocol.RegisterParallelTaskServer(srv.server, rpcHandler.NewQuery(srv.config))

	if err := srv.server.Serve(listener); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func (srv *RpcServer) Shutdown() error {
	srv.server.GracefulStop()
	return nil
}
