package mqttx

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var callbackOnConnectHandler mqtt.OnConnectHandler

// MQTTxHandlerOnConnect 在与MQTT服务器建立连接时，订阅主题
var MQTTxHandlerOnConnect mqtt.OnConnectHandler = func(client mqtt.Client) {
	cliOpts := client.OptionsReader()
	server := FormatServerAddr(cliOpts.Servers()[0].String())
	TraceInfo("mqtt(%v) - client(%v) connected", server, cliOpts.ClientID())

	if callbackOnConnectHandler != nil {
		callbackOnConnectHandler(client)
	}
}
