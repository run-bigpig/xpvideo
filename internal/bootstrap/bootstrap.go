package bootstrap

import (
	"context"
	"github.com/run-bigpig/xpvideo/internal/config"
	"github.com/run-bigpig/xpvideo/internal/controller"
)

type Boot struct {
	ctx context.Context
}

func NewBoot() *Boot {
	// 初始化配置
	config.Init()
	return &Boot{}
}

func (b *Boot) Startup(ctx context.Context) {
	b.ctx = ctx
}

func (b *Boot) NewController() *controller.AppController {
	return controller.NewAppController(b.ctx)
}
