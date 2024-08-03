package main

import (
	"os"

	"github.com/GiterLab/mqttx"
	"github.com/beego/beego/v2/core/logs"
)

// glog 全局日志变量
var glog *logs.BeeLogger

var clientPool *mqttx.MQTTxClientPool

func main() {
	// 设置日志
	glog = logs.NewLogger(10000)
	glog.SetLogger("console", `{"level":7}`)
	glog.EnableFuncCallDepth(true)
	glog.SetLogFuncCallDepth(3)

	mqttx.SetUserDebug(func(format string, level int, v ...interface{}) {
		switch level {
		case mqttx.LevelInformational:
			glog.Info(format, v...)
		case mqttx.LevelError:
			glog.Error(format, v...)
		case mqttx.LevelDebug:
			glog.Debug(format, v...)
		}
	})
	mqttx.Debug(true)

	// 初始化MQTTx服务
	servers := []*mqttx.MQTTxServer{
		{
			Vendor:   "emqx",
			Scheme:   "tcp",
			Domain:   "broker.emqx.io",
			IP:       "44.195.202.69",
			Port:     "1883",
			ClientID: "github_giterlab_mqttx_example",
			Username: "emqx",
			Password: "public",
		},
		{
			Vendor:   "mosquitto",
			Scheme:   "tcp",
			Domain:   "test.mosquitto.org",
			IP:       "91.121.93.94",
			Port:     "1883",
			ClientID: "github_giterlab_mqttx_example",
			Username: "",
			Password: "",
		},
	}

	clientPool = mqttx.NewMQTTxClientPool()

	// 对不同的mqtt服务器进行对应的处理
	// 建议: 每一次调用 Connect 方法时，只传入一个服务器的配置
	for _, v := range servers {
		switch v.Vendor {
		case mqttx.MQTT_BROKER_MOSQUITTO:
			var mqttServers []*mqttx.MQTTxServer
			mqttServers = append(mqttServers, v) // 只设置 1 个服务器地址
			err := mqttx.Connect(clientPool, mqttServers, nil, mosquittoCallbackOnConnectHandler, nil, nil)
			if err != nil {
				glog.Error("connect to mosquitto failed: %v", err)
				os.Exit(1)
			}
		case mqttx.MQTT_BROKER_EMQX:
			var mqttServers []*mqttx.MQTTxServer
			mqttServers = append(mqttServers, v) // 只设置 1 个服务器地址
			err := mqttx.Connect(clientPool, mqttServers, nil, emqxCallbackOnConnectHandler, nil, nil)
			if err != nil {
				glog.Error("connect to emqx failed: %v", err)
				os.Exit(1)
			}
		default:
			glog.Error("MQTT Server Vendor not support: %v", v.Vendor)
			os.Exit(1)
		}
	}

	go func() {
		for {
			client := clientPool.Get(servers[0].Server())
			if client != nil {
				glog.Debug("server(%v) IsConnected(%v) - connection count: %v, message: %s", client.Server(), client.Client.IsConnected(), client.GetServerConnectionCount(), client.GetOtherOpts("message"))
			} else {
				glog.Debug("server(%v) client is nil", mqttx.FormatServerAddr(servers[0].Server()))
			}
			mqttx.SleepSecond(5)
		}
	}()

	select {}
}
