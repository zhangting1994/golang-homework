package service

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type SignalHandler struct {
	ctx context.Context
	cancel context.CancelFunc
	sigChan chan os.Signal
	shutdownFunc func() error
}

func (sigHandler *SignalHandler) ShutdownFuncReg(f func() error) {
	sigHandler.shutdownFunc = f
}

func (sigHandler *SignalHandler) Handle(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	sigHandler.ctx = ctx
	sigHandler.cancel = cancel

	for {
		select {
		case <- sigHandler.ctx.Done(): // 上下文取消 进行相关处理
			fmt.Println("context已关闭，开始关闭其他服务")
			return sigHandler.shutdownFunc()
		case s := <- sigHandler.sigChan: // 收到信号 取消上下文
			msg := fmt.Sprintf("捕获到SIGINT:%s，开始关闭context",s)
			fmt.Println(msg)
			sigHandler.cancel()
		}
	}
}

func NewSignalHandler() *SignalHandler {
	// 信号
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT)

	return &SignalHandler{
		sigChan: sigChan,
	}
}