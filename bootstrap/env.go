package bootstrap

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type Env struct {
	AppEnv          string        `mapstructure:"APP_ENV"`
	ServerAddress   string        `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout  int           `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost          string        `mapstructure:"DB_HOST"`
	DBPort          string        `mapstructure:"DB_PORT"`
	DBUser          string        `mapstructure:"DB_USER"`
	DBPass          string        `mapstructure:"DB_PASS"`
	DBName          string        `mapstructure:"DB_NAME"`
	TokenExpiryHour time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRATION"`
	TokenSecret     string        `mapstructure:"ACCESS_TOKEN_SECRET"`
}

func NewEnv() *Env {

	// LoadConfig reads configuration from file or environment variables.

	log.Print("Reading configuration form environment")

	env := Env{}

	// default values
	env.TokenExpiryHour = 4 * time.Hour
	env.TokenSecret = "some-secret"

	viper.SetConfigFile(".env")

	// call viper.AutomaticEnv() to tell viper to automatically override values
	// that it has read from config file with the values of the corresponding
	// environment variables if they exist.
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
