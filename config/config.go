package config

import (
	"time"

	"github.com/spf13/viper"
)

// instance Configuration instance
var instance *Configuration

// Core core
type Core struct {
	Port string `yaml:"port"`
}

// Docker docker
type Docker struct {
	Host string `yaml:"host"`
}

// Configuration configurations
type Configuration struct {
	Version       string `mapstructure:"VERSION"`
	AppName       string `mapstructure:"APP_NAME"`
	BuildEnv      string `mapstructure:"BUILD_ENV"`
	ServerPort    string `mapstructure:"SERVER_PORT"`
	ServerBaseURL string `mapstructure:"SERVER_BASE_URL"`
	DockerHost    string `mapstructure:"DOCKER_HOST"`

	AdminUserName     string `mapstructure:"ADMIN_USER_NAME" json:"admin_user_name"`
	AdminUserPassword string `mapstructure:"ADMIN_USER_PASSWORD" json:"admin_user_password"`

	JWTSecret string        `mapstructure:"JWT_SECRET" json:"jwt_secret"`
	JWTExpiry time.Duration `mapstructure:"JWT_EXPIRY" json:"jwt_expiry"`

	DockerFilter   string `mapstructure:"DOCKER_FILTER" json:"docker_filter"`
	DockerTailSize int    `mapstructure:"DOCKER_TAILSIZE" json:"docker_tailsize"`

	DBPath string `mapstructure:"DB_PATH" json:"db_path"`
}

// New setup config
func New(cfgFile string) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".go-base" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigName("")
		viper.SetConfigType("env")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	var err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	instance = &Configuration{}

	err = viper.Unmarshal(&instance)
	if err != nil {
		panic(err)
	}

}

// GetInstance get instance
func GetInstance() *Configuration {
	if instance == nil {
		panic("instance is getting NULL")
	}
	return instance
}
