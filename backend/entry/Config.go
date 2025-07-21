package entry

type Config struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

func NewConfig(height int, width int) *Config {
	return &Config{height, width}
}
