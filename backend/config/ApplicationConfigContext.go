package config

type ApplicationConfigContext struct {
	WindowConfig *WindowConfig
}

func InitApplicationConfigContext() *ApplicationConfigContext {

	windowConfig := NewWindowConfig(600, 800)

	return &ApplicationConfigContext{windowConfig}
}
