package main

import (
	"log"
	"gomqtt/publisher"
)

func main() {
	log.Println("Start IOT")

	/* Start MQTT */
	publisher.Pub()
}
