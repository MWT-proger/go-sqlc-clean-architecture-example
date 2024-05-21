package config

import (
	"flag"

	"gopkg.in/yaml.v2"

	"github.com/MWT-proger/go-sqlc-clean-architecture-example/configs"
)

type Config struct {
	HostServer string `yaml:"host"`
	Database   struct {
		DSN string `yaml:"dsn"`
	} `yaml:"database"`
	LogLevel string `yaml:"log_level"`
}

var newConfig Config

// initDefaultConfig Присваивает локальной не импортируемой переменной newConfig базовые значения.
// Вызывается один раз при старте проекта.
func initDefaultConfig() {
	newConfig = Config{
		HostServer: ":8000",
		LogLevel:   "info",
	}
}

// InitConfig() Инициализирует локальную не импортируемую переменную newConfig.
// Вызывает все доступные методы получения конфигов.
// Вызывается один раз при старте проекта
func InitConfig() (*Config, error) {
	initDefaultConfig()
	if err := parseYAML(); err != nil {
		return nil, err
	}
	parseFlags()
	return &newConfig, nil
}

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных глобальной переменной newConfig
func parseFlags() {

	flag.StringVar(&newConfig.HostServer, "a", newConfig.HostServer, "адрес и порт для запуска сервера")
	flag.StringVar(&newConfig.Database.DSN, "d", newConfig.Database.DSN, "строка с адресом подключения к БД")
	flag.StringVar(&newConfig.LogLevel, "l", "info", "уровень логирования")
	flag.Parse()
}

// parseYAML собирает файл конфигурации из configs.yaml
func parseYAML() error {
	file, err := configs.FSConfigsYAML.Open("configs.yaml")
	if err != nil {
		return err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&newConfig); err != nil {
		return err
	}
	return nil
}
