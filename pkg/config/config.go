package config

import (
	"log"

	"github.com/spf13/viper"
)

// Estrutura para armazenar as variáveis de config, e cada uma com mapa para o Viper
type Settings struct {
	DbHost     string `mapstructure:"DB_HOST"`
	Dbport     string `mapstructure:"DB_PORT"`
	DbName     string `mapstructure:"DB_NAME"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPass     string `mapstructure:"DB_PASS"`
	Env        string `mapstructure:"ENV"`
	JwtExpires string `mapstructure:"JWT_EXPIRES"`
	JwtSecret  string `mapstructure:"JWT_SECRET"`
}

func New() *Settings {

	var cfg Settings

	// Definimos o arquivo de configuração e o tipo para .env
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	// Ler quaisquer variáveis do arquivo de configuração
	err := viper.ReadInConfig()

	if err != nil {
		log.Println("No env file, using enviromment variables.", err)
	}

	// Desempacota todas as informações carregadas em nossa estrutura de configurações
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("Error trying to unmarshal configuration", err)
	}

	return &cfg
}
