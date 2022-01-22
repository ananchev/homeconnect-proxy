package mqttpublisher

import (
	"github.com/ananchev/homeconnect-proxy/internal/logger"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//define a function for the default message handler
var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	logger.Info("Event '{event}' data publised to mqtt topic '{topic}'", "event", "event", "topic", "topic")
}

//store here the root topic as set by StartMqttPublisher
var RootTopic string

//MQTT client
var client mqtt.Client

//Create a ClientOptions struct setting the broker address, clientid, turn
//off trace output and set the default message handler
func StartMqttPublisher(server string, port string, rootTopic string) {

	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://" + server + ":" + port)
	opts.SetClientID("homeconnect-proxy")
	opts.SetDefaultPublishHandler(f)

	RootTopic = rootTopic

	client = mqtt.NewClient(opts)
	logger.Info("Initialized connection to MQTT broker '{b}'", "b", server+":"+port)

}

func Publish(ev Event) {
	topic := RootTopic + "/" + ev.EquipmentID + "/" + ev.EventName
	payload := ev.EventData

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logger.Error("Error in MQTT client connection: '{err}'", "err", token.Error())
		return
	}

	token := client.Publish(topic, 0, false, payload)
	token.Wait()

	client.Disconnect(250)
}