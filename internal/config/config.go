package config

type Config struct {
	DefaultPath string
	MaxFileSize int64
	// Add more configuration options here
}

func Load() (*Config, error) {
	// Load configuration from file or environment variables
	return &Config{
		DefaultPath: ".",
		MaxFileSize: 1024 * 1024 * 100, // 100MB
	}, nil
}
