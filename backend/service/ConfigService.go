package service

import (
	"cat-clipboard/backend/config"
)

type ConfigService struct {
	applicationConfigContext *config.ApplicationConfigContext
}

func NewConfigService(applicationConfigContext *config.ApplicationConfigContext) *ConfigService {
	return &ConfigService{applicationConfigContext}
}

func (c *ConfigService) GetWindowConfig() *config.WindowConfig {
	return c.applicationConfigContext.WindowConfig
}
