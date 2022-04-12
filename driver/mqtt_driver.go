package driver

import (
	"GDP_KafkaMqtt/config"
	"GDP_KafkaMqtt/kafka"
	"GDP_KafkaMqtt/model"
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	//run controller here
	var datareceive model.Scada
	err := json.Unmarshal(msg.Payload(), &datareceive)
	if err != nil {
		//panic(err)
		fmt.Println(err)
		return
	}
	kafka.Kafka_Pub(datareceive)

}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}
var ClientMQ mqtt.Client

func Connect() {
	var broker = config.Cfg.Mqtt.MqttBrokerHost
	var port = config.Cfg.Mqtt.MqttPort
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(config.Cfg.Mqtt.MqttClientID)
	opts.SetUsername(config.Cfg.Mqtt.MqttUserName)
	opts.SetPassword(config.Cfg.Mqtt.MqttPassword)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	sub(client)
	ClientMQ = client
}

func sub(client mqtt.Client) {
	topic := config.Cfg.Mqtt.MqttTopicSub
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}
