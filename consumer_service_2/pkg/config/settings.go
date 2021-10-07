package config

import (
	"github.com/spf13/viper"
	"log"
)

//Settings represents the configuration that we can provide
//from the outside in order to run the application in different ways.
type Settings struct {
	Host                string `mapstructure:"HOST"`
	Port                string `mapstructure:"PORT"`

	ReadBufferSize      int    `mapstructure:"READ_BUFFER_SIZE"`
	WriteBufferSize     int    `mapstructure:"WRITE_BUFFER_SIZE"`

	MsgQueueCheckOrigin bool   `mapstructure:"MSG_QUEUE_CHECK_ORIGIN"`
	MsgQueueName        string `mapstructure:"MSG_QUEUE_NAME"`
	MsgQueueRoutingKey  string `mapstructure:"MSG_QUEUE_ROUTE_KEY"`
	MsgQueueExchangeName  string `mapstructure:"MSG_QUEUE_EXCHANGE_NAME"`
	MsgQueueExchangeKind  string `mapstructure:"MSG_QUEUE_EXCHANGE_KIND"`
	MsgQueueConnString string `mapstructure:"MSG_QUEUE_CONN_STRING"`
	MsgQueueConsumerName string `mapstructure:"MSG_QUEUE_CONSUMER_NAME"`
}

func New(envPath string) *Settings {
	var cfg Settings

	viper.SetConfigFile(envPath)
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