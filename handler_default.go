package mqttx

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var callbackDefaultHandler mqtt.MessageHandler

// DefaultMQTTxHandler 默认的mqtt消息处理函数， 没有匹配的topic时，使用此函数处理
var MQTTxHandlerDefault mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	cliOpts := client.OptionsReader()
	server := FormatServerAddr(cliOpts.Servers()[0].String())
	TraceInfo("mqtt(%v) - client(%v) topic - default: %s, message: %X", server, cliOpts.ClientID(), msg.Topic(), msg.Payload())

	if callbackDefaultHandler != nil {
		callbackDefaultHandler(client, msg)
	}
}
