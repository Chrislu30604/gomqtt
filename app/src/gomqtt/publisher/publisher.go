package publisher

import (
	"fmt"
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const (
	HOST     = "tcp://localhost"
	PORT     = "9001"
	USERNAME = "TODO"
	PASSWORD = "TODO"
	CLIENTID = "TODO"
	TOPIC    = "/hello/world"
)

func f(client MQTT.Client, msg MQTT.Message) {
	log.Println("Default Publish handler: ", msg.Topic(), string(msg.Payload()))
}

func Pub() {
	log.Printf("Start Rasp Publisher on %v:%v\n", HOST, PORT)
	opts := MQTT.NewClientOptions()
	opts.AddBroker(HOST + ":" +PORT)
	opts.SetUsername(USERNAME)
	opts.SetClientID(CLIENTID)
	opts.SetDefaultPublishHandler(f)

	//create and start a client using the above ClientOptions
	clientPub := MQTT.NewClient(opts)
	if token := clientPub.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := clientPub.Publish(TOPIC, 0, false, text)
		token.Wait()
	}

	clientPub.Disconnect(250)
}
