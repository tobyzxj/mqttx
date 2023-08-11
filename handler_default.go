package mqttx

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var callbackDefaultHandler mqtt.MessageHandler

// DefaultMQTTxHandler 默认的mqtt消息处理函数， 没有匹配的topic时，使用此函数处理
var MQTTxHandlerDefault mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	TraceInfo("mqtt topic - default: %s, message: %X", msg.Topic(), msg.Payload())

	if callbackDefaultHandler != nil {
		callbackDefaultHandler(client, msg)
	}
}
