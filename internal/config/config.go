package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env             string        `yaml:"env"  env-default:"local"` // где мы находимся
	GRPC            GRPCConfig    `yaml:"grpc"`                     // настройки grpc
	DB              Database      `yaml:"database"`                 // настройки базы данных
	MigrationFolder string        `yaml:"maigration_folder"`        // папка с миграциями
	ClientsCfg      ClientsConfig `yaml:"clients_cfg"`              // настройки клиентов
	KafkaHost       string        `yaml:"kafka_host"`
}

type GRPCConfig struct {
	Port int `yaml:"port"`
}

type Database struct {
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Host       string `yaml:"host"`
	Name       string `yaml:"name"`
	ReqTimeOut int    `yaml:"request_timeout"` // в секундах
}

type Client struct {
	Address      string        `yaml:"address"`
	Timeout      time.Duration `yaml:"timeout"`
	RetriesCount int           `yaml:"retries_count"`
}

type ClientsConfig struct {
	GetMail Client `yaml:"get_mail_consumer"`
}

// Must говорит о том, что метод должен обязательно запуститься
// Если чтот-то пойдет не так, то паникуем, т.е. дальше нет смысла работать
func MustLoad() *Config {
	var cfg Config

	path := fetchConfigPath() // получаем путь к файлу конфига
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

// Получаем путь к файлу config из флага или среды окружения
// Приоритет: флаг > среда > по-умолчанию
// значение по-умолчанию: ""
// запуск с указанием в среде окружения: CONFIG_PATH=./path/file.yaml cards
// запуск с указанием флага: cards --config=./path/file.yaml
func fetchConfigPath() string {
	var res string

	// сначала проверяем флаг
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	// есом флага нет, то берем из среды окружения
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
