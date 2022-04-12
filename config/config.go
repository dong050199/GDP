package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	Environment string
	Mqtt        MQTTConf
	Kafka       KAFKAConf
}

type MQTTConf struct {
	MqttBrokerHost string
	MqttUserName   string
	MqttPassword   string
	MqttClientID   string
	MqttPort       int
	MqttTopicPub   string
	MqttTopicSub   string
}

type KAFKAConf struct {
	KafkaBootStrapServer string
	KafkaGroupID         string
	KafkaAutoSetReset    string
	KafkaTopicSub        string
	KafkaTopicPub        string
}

var Cfg Configuration

func GetConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(Cfg)
}
