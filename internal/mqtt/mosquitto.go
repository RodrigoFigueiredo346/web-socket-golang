package mqtt

import (
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var client MQTT.Client

// Função para iniciar o cliente MQTT
func InitMqtt() {
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://18.230.128.87:18833")
	//opts.SetClientID("go-client")
	// opts.SetUsername("master")
	// opts.SetPassword("G6*utV")
	opts.SetUsername("panel")
	opts.SetPassword("panel")

	// Função de callback para lidar com mensagens recebidas
	opts.SetDefaultPublishHandler(onMessageReceived)

	client = MQTT.NewClient(opts)

	// Conexão com servidor MQTT
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println("Error connecting to server:", token.Error())
		return
	}

	// tópico de assinatura
	topicSub := "server"
	if token := client.Subscribe(topicSub, 0, nil); token.Wait() && token.Error() != nil {
		log.Println("Error subscribing to topic:", token.Error())
		return
	}
	log.Printf("Listening to messages in topic %s...", topicSub)
}

func Publish(topic, message string) {
	if client != nil && client.IsConnected() {
		token := client.Publish(topic, 1, false, message)
		token.Wait()
		log.Printf("Message published to the topic %s: %s", topic, message)
	} else {
		log.Println("MQTT client not connected.")
	}
}

func onMessageReceived(client MQTT.Client, msg MQTT.Message) {
	log.Printf("Received message in the topic %s: %s\n", msg.Topic(), msg.Payload())
	// end := time.Now().Format("15:04:05.000")
	// fmt.Println("end:xxxxxxxxxx ", end)

}
