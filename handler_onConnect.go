package mqttx

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var callbackOnConnectHandler mqtt.OnConnectHandler

// MQTTxHandlerOnConnect 在与MQTT服务器建立连接时，订阅主题
var MQTTxHandlerOnConnect mqtt.OnConnectHandler = func(client mqtt.Client) {
	cliOpts := client.OptionsReader()

	TraceInfo("mqtt - client(%v) connected", cliOpts.ClientID())

	if callbackOnConnectHandler != nil {
		callbackOnConnectHandler(client)
	}
}
