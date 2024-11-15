package envconfig

import (
	"cmp"
	"log"

	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/uszebr/loadmonitor/inner/domain/job"
	"github.com/uszebr/loadmonitor/inner/logger"
)

type Config struct {
	LogLevel      logger.LogLevel `yaml:"log_level"`
	Multiplier    int             `yaml:"complexity_multiplier"`
	MultiplyValue int             `yaml:"multiply_value"`
}

type StartValues struct {
	LogLevel      logger.LogLevel
	Multiplier    int
	MultiplyValue int
}

// Initializing Envrionment; Config; Logger; DB;
func MustLoad() StartValues {

	if err := godotenv.Load(); err != nil {
		log.Printf("godotenv issue: %v\n", err)
	}

	// reading config path from env to run proper configuration
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("CONFIG_PATH env var is not set")
	}
	if _, err := os.Stat(configPath); err != nil {
		panic("Config file does not exist:" + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("Can not read config file: :" + configPath + " :" + err.Error())
	}

	startValues := StartValues{}
	startValues.LogLevel = cmp.Or(cfg.LogLevel, logger.DEFAULTLOGLEVEL)
	logger.MustInitLogger(startValues.LogLevel)

	job.SetComplexityMultiplier(cfg.Multiplier)
	startValues.Multiplier = cfg.Multiplier

	job.SetMultiplyValue(cfg.MultiplyValue)
	startValues.MultiplyValue = cfg.MultiplyValue

	return startValues
}
