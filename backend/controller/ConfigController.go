package controller

import (
	"cat-clipboard/backend/entry"
	"cat-clipboard/backend/service"
	"context"
)

type ConfigController struct {
	ctx           context.Context
	configService *service.ConfigService
}

func NewConfigController(ctx context.Context, configService *service.ConfigService) *ConfigController {
	return &ConfigController{ctx: ctx, configService: configService}
}

func (c *ConfigController) QueryAllConfig() *entry.Config {

	//value := c.ctx.Value("height")
	//println(value)

	println("查询配置信息")
	windowConfig := c.configService.GetWindowConfig()
	return entry.NewConfig(windowConfig.Height, windowConfig.Width)
}

func InitConfig() {
	println("初始化配置类")
}
