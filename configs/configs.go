package configs

import (
	"flag"
)

type Config struct {
	HostServer  string `env:"RUN_ADDRESS"`
	DatabaseDSN string `env:"DATABASE_URI"`
	LogLevel    string
}

var newConfig Config

// initDefaultConfig Присваивает локальной не импортируемой переменной newConfig базовые значения.
// Вызывается один раз при старте проекта.
func initDefaultConfig() {
	newConfig = Config{
		HostServer:  ":8000",
		LogLevel:    "info",
		DatabaseDSN: "",
	}
}

// InitConfig() Инициализирует локальную не импортируемую переменную newConfig.
// Вызывает все доступные методы получения конфигов.
// Вызывается один раз при старте проекта
func InitConfig() *Config {
	initDefaultConfig()
	parseFlags()
	return &newConfig
}

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных глобальной переменной newConfig
func parseFlags() {

	flag.StringVar(&newConfig.HostServer, "a", newConfig.HostServer, "адрес и порт для запуска сервера")
	flag.StringVar(&newConfig.DatabaseDSN, "d", newConfig.DatabaseDSN, "строка с адресом подключения к БД")
	flag.StringVar(&newConfig.LogLevel, "l", "info", "уровень логирования")
	flag.Parse()
}
