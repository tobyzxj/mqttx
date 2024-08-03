package mqttx

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var callbackConnectionLostHandler mqtt.ConnectionLostHandler

// MQTTxHandlerConnectionLost 在与MQTT服务器连接断开时回调
var MQTTxHandlerConnectionLost mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	cliOpts := client.OptionsReader()
	server := FormatServerAddr(cliOpts.Servers()[0].String())
	TraceError("mqtt(%v) - client(%v) connection lost, error: %s", server, cliOpts.ClientID(), err)

	if callbackConnectionLostHandler != nil {
		callbackConnectionLostHandler(client, err)
	}
}
