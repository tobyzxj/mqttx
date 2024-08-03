package mqttx

import (
	"encoding/json"
	"sync"
)

// MQTTxClientPool MQTT客户端连接池
type MQTTxClientPool struct {
	Clients []*MQTTxClient `json:"clients"` // MQTT客户端连接池
	mux     *sync.RWMutex  // 读写锁
}

// GetClients 获取所有客户端连接
func (p *MQTTxClientPool) GetClients() []*MQTTxClient {
	if p != nil {
		p.mux.RLock()
		defer p.mux.RUnlock()

		return p.Clients
	}
	return nil
}

// Len 返回连接池的长度
func (p *MQTTxClientPool) Len() int {
	if p != nil {
		p.mux.RLock()
		defer p.mux.RUnlock()

		return len(p.Clients)
	}

	return 0
}

// Add 添加一个客户端连接
func (p *MQTTxClientPool) Add(client *MQTTxClient) {
	if p != nil && client != nil {
		p.mux.Lock()
		defer p.mux.Unlock()

		p.Clients = append(p.Clients, client)
	}
}

// Remove 移除一个客户端连接
func (p *MQTTxClientPool) Remove(server string) {
	if p != nil && server != "" {
		p.mux.Lock()
		defer p.mux.Unlock()

		for i, c := range p.Clients {
			if c.Server() == FormatServerAddr(server) {
				p.Clients = append(p.Clients[:i], p.Clients[i+1:]...)
				break
			}
		}
	}
}

// Get 获取一个客户端连接
func (p *MQTTxClientPool) Get(server string) *MQTTxClient {
	if p != nil && server != "" {
		p.mux.RLock()
		defer p.mux.RUnlock()

		for _, c := range p.Clients {
			if c.Server() == FormatServerAddr(server) {
				return c
			}
		}
	}

	return nil
}

// GetMinConnectionCountClient 获取所有客户端里面连接数最少的一个
func (p *MQTTxClientPool) GetMinConnectionCountClient() *MQTTxClient {
	if p != nil {
		p.mux.RLock()
		defer p.mux.RUnlock()

		if len(p.Clients) == 0 {
			return nil
		}

		var min *MQTTxClient
		for _, c := range p.Clients {
			if min == nil {
				if c.Client.IsConnected() {
					min = c
				}
				continue
			}
			if !c.Client.IsConnected() {
				continue
			}
			if c.GetServerConnectionCount() < min.GetServerConnectionCount() {
				min = c
			}
		}
		if min != nil {
			if min.GetServerConnectionCount() == MaxInt {
				return nil
			}
		}

		return min
	}

	return nil
}

// Iterate 遍历所有客户端连接, 主要用于定时主动检查连接状态，如有必要进行重连操作
func (p *MQTTxClientPool) Iterate(f func(c *MQTTxClient)) {
	if p != nil {
		p.mux.RLock()
		defer p.mux.RUnlock()

		for _, c := range p.Clients {
			f(c)
		}
	}
}

// String 序列化成字符串
func (p MQTTxClientPool) String() string {
	body, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(body)
}

// NewMQTTxClientPool 创建一个MQTT客户端连接池
func NewMQTTxClientPool() *MQTTxClientPool {
	clientPool := &MQTTxClientPool{
		Clients: make([]*MQTTxClient, 0),
		mux:     new(sync.RWMutex),
	}
	return clientPool
}
