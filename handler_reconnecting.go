package mqttx

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var callbackReconnectingHandler mqtt.ReconnectHandler

// MQTTxHandlerReconnecting 在与MQTT服务器重连时回调
var MQTTxHandlerReconnecting mqtt.ReconnectHandler = func(client mqtt.Client, clientOptions *mqtt.ClientOptions) {
	cliOpts := client.OptionsReader()

	TraceInfo("mqtt - client(%v) reconnecting", cliOpts.ClientID())

	if callbackReconnectingHandler != nil {
		callbackReconnectingHandler(client, clientOptions)
	}
}
