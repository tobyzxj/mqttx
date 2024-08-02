package main

import (
	"time"

	"github.com/GiterLab/mqttx"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mosquittoCallbackOnConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	cliOpts := client.OptionsReader()
	server := mqttx.FormatServerAddr(cliOpts.Servers()[0].String())
	cli := clientPool.Get(server)
	if cli == nil {
		glog.Error("mqtt(%v) - client(%v) not found", server, cliOpts.ClientID())
		return
	}

	// 清理客户端状态
	cli.ClearOtherOpts()

	// 订阅主题，每个工程需要订阅的主题不一样，需要在此处修改
	topic := ""

	topic = "mqttx/example/#"
	glog.Debug("mqtt(%v) - client(%v) subscribe topic: %v", server, cliOpts.ClientID(), topic)
	token := client.Subscribe(topic, 0, MessageHandlerForMosquitto)
	isSeccess := token.WaitTimeout(10 * time.Second)
	if !isSeccess {
		glog.Error("mqtt(%v) - client(%v) subscribe topic timeout", server, cliOpts.ClientID())
	}
	if isSeccess && token.Error() != nil {
		glog.Error("mqtt(%v) - client(%v) subscribe topic error: %v", server, cliOpts.ClientID(), token.Error())
	}
}

// MessageHandler 系统消息订阅
var MessageHandlerForMosquitto mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	cliOpts := client.OptionsReader()
	server := mqttx.FormatServerAddr(cliOpts.Servers()[0].String())
	cli := clientPool.Get(server)
	if cli == nil {
		glog.Error("mqtt(%v) - client(%v) not found", server, cliOpts.ClientID())
		return
	}

	if cli != nil {
		// test
		glog.Debug("mqtt(%v) - client(%v) client info: %v", server, cliOpts.ClientID(), cli)
	}

	glog.Info("mqtt(%v) topic - data: %s, message: % X", server, msg.Topic(), msg.Payload())
	cli.SetOtherOpts("message", string(msg.Payload()))
}
