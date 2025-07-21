package config

type WindowConfig struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

func NewWindowConfig(height int, width int) *WindowConfig {
	return &WindowConfig{height, width}
}
