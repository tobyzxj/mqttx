package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var emqxCallbackOnConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	cliOpts := client.OptionsReader()
	server := cliOpts.Servers()[0].String()
	cli := clientPool.Get(server)

	// 清理客户端状态
	cli.ClearOtherOpts()

	// 订阅主题，每个工程需要订阅的主题不一样，需要在此处修改
	topic := ""

	topic = "mqttx/example/#"
	glog.Debug("mqtt(%v) - client(%v) subscribe topic: %v", server, cliOpts.ClientID(), topic)
	if token := client.Subscribe(topic, 0, MessageHandlerForEmqx); token.Wait() && token.Error() != nil {
		glog.Error("mqtt(%v) - client(%v) subscribe topic error: %v", server, cliOpts.ClientID(), token.Error())
	}
}

// MessageHandlerForEmqx 系统消息订阅
var MessageHandlerForEmqx mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	cliOpts := client.OptionsReader()
	server := cliOpts.Servers()[0].String()

	cli := clientPool.Get(server)
	if cli != nil {
		// test
		glog.Debug("mqtt(%v) - client(%v) client info: %v", server, cliOpts.ClientID(), cli)
	}

	glog.Info("mqtt(%v) topic - data: %s, message: % X", server, msg.Topic(), msg.Payload())
	cli.SetOtherOpts("message", string(msg.Payload()))
}
