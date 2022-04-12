package main

import (
	D "GDP_KafkaMqtt/driver"
	"GDP_KafkaMqtt/kafka"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	D.Connect()
	go func() {
		kafka.Kafka_Sub(D.ClientMQ)
	}()
	r := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))

}
