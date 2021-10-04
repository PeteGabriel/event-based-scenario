package config

import (
	"github.com/spf13/viper"
	"log"
)

//Settings represents the configuration that we can provide
//from the outside in order to run the application in different ways.
type Settings struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
	ReadBufferSize int `mapstructure:"READ_BUFFER_SIZE"`
	WriteBufferSize int `mapstructure:"WRITE_BUFFER_SIZE"`
	MsgQueueHost string `mapstructure:"MSG_QUEUE_HOST"`
	MsgQueuePort string `mapstructure:"MSG_QUEUE_PORT"`
}

func New() *Settings {
	var cfg Settings
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No env file found.", err)
	}

	//try to assign read variables into golang struct
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("Error trying to unmarshal configuration.", err)
	}

	return &cfg
}