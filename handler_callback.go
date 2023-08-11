package mqttx

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// CallbackDefaultHandlerSet set callbackDefaultHandler
func CallbackDefaultHandlerSet(handler mqtt.MessageHandler) {
	callbackDefaultHandler = handler
}

// CallbackOnConnectHandlerSet set callbackOnConnectHandler
func CallbackOnConnectHandlerSet(handler mqtt.OnConnectHandler) {
	callbackOnConnectHandler = handler
}

// CallbackConnectionLostHandlerSet set callbackConnectionLostHandler
func CallbackConnectionLostHandlerSet(handler mqtt.ConnectionLostHandler) {
	callbackConnectionLostHandler = handler
}

// CallbackReconnectingHandlerSet set callbackReconnectingHandler
func CallbackReconnectingHandlerSet(handler mqtt.ReconnectHandler) {
	callbackReconnectingHandler = handler
}
