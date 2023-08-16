package mqttx

import (
	"errors"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Connect 初始化MQTT服务
func Connect(servers []*MQTTxServer, defaultPublishHandler mqtt.MessageHandler, onConnectHandler mqtt.OnConnectHandler, connectionLostHandler mqtt.ConnectionLostHandler, reconnectingHandler mqtt.ReconnectHandler) error {
	if len(servers) == 0 {
		return errors.New("MQTTx servers is empty")
	}

	for k, v := range servers {
		client, err := NewMQTTxClient(v)
		if err != nil {
			return err
		}
		TraceDebug("MQTT Client(%v - %v) Connect...: %v", k, v.Vendor, client)
		if defaultPublishHandler == nil {
			defaultPublishHandler = MQTTxHandlerDefault
		}
		if onConnectHandler == nil {
			onConnectHandler = MQTTxHandlerOnConnect
		}
		if connectionLostHandler == nil {
			connectionLostHandler = MQTTxHandlerConnectionLost
		}
		if reconnectingHandler == nil {
			reconnectingHandler = MQTTxHandlerReconnecting
		}
		err = client.Connect(defaultPublishHandler, onConnectHandler, connectionLostHandler, reconnectingHandler)
		if err != nil {
			return err
		}
		ClientPool.Add(client)
	}
	return nil
}

// NewMQTTxClient 根据服务器的信息，创建MQTT客户端
func NewMQTTxClient(server *MQTTxServer) (*MQTTxClient, error) {
	if server == nil {
		return nil, errors.New("server is nil")
	}
	client := &MQTTxClient{
		Vendor:   server.Vendor,
		Scheme:   server.Scheme,
		Domain:   server.Domain,
		IP:       server.IP,
		Port:     server.Port,
		Cert:     server.Cert,
		ClientID: server.ClientID,
		Username: server.Username,
		Password: server.Password,
		Opts:     nil,
		Client:   nil,
	}
	return client, nil
}
