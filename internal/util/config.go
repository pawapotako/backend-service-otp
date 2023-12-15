package util

import (
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	App     AppConfig
	DB      DatabaseConfig
	Kafka   KafkaConfig
	Grpc    GrpcConfig
	Swagger SwaggerConfig
}

type AppConfig struct {
	Port int
}

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

type KafkaConfig struct {
	Host  string
	Port  int
	Topic string
}

type GrpcConfig struct {
	Host string
	Port int
}

type SwaggerConfig struct {
	BasePath string
}

func LoadConfig() (config Config, err error) {

	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		// for debug mode
		viper.AddConfigPath("../../configs")
		if err := viper.ReadInConfig(); err != nil {
			Logger.Fatal("Error reading config file", zap.Error(err))
		}
	}

	viper.AddConfigPath("./configs")
	viper.SetConfigName("config.dev")
	viper.MergeInConfig()

	err = viper.Unmarshal(&config)
	if err != nil {
		Logger.Fatal("unable to decode into struct", zap.Error(err))
	}

	return
}
