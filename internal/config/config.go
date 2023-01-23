package config

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/atrian/prePressFileChecker/pkg/logger"
)

type AppConfig struct {
	YamlConfigPath  string
	ImageExtension  string `yaml:"image_extension"`
	ExcelFilePath   string `yaml:"excel_file_path"`
	ImageFolderPath string `yaml:"image_folder_path"`
	logger          logger.Logger
}

func NewConfig(logger logger.Logger) (*AppConfig, error) {
	cfg := &AppConfig{
		logger: logger,
	}

	cfg.loadFlags()

	configFile, err := os.Open(cfg.YamlConfigPath)
	if err != nil {
		cfg.logger.Error("Can't read config file", err)
		return nil, err
	}

	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			cfg.logger.Error("configFile.Close error", err)
		}
	}(configFile)

	d := yaml.NewDecoder(configFile)

	if err := d.Decode(&cfg); err != nil {
		cfg.logger.Error("Can't decode yaml", err)
		return nil, err
	}

	return cfg, nil
}

// loadFlags загрузка конфигурации из флагов
func (config *AppConfig) loadFlags() {
	confPath := flag.String("p", "./config.yaml", "Path for YAML configuration file")
	flag.Parse()
	config.YamlConfigPath = *confPath
}
