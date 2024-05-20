package configs

type Config struct {
	HostServer  string `env:"RUN_ADDRESS"`
	DatabaseDSN string `env:"DATABASE_URI"`
	LogLevel    string
}
