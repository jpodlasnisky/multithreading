package config

import "github.com/spf13/viper"

type Config struct {
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	ViaCepHost    string `mapstructure:"VIACEP_HOST_API"`
	BrasilCepHost string `mapstructure:"BRASILCEP_HOST_API"`
	DBFilePath    string `mapstructure:"DB_FILE_PATH"`
}

func LoadConfig(path string) (Config, error) {
	var config Config
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile("config.env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return config, err
}
