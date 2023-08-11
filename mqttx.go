package mqttx

import (
	"encoding/json"
	"errors"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var MQTTxClientPools []*MQTTxClient // MQTT客户端连接

// Cert 证书信息
type Cert struct {
	Cafile         string `json:"cafile"`           // CA证书文件路径
	ClientCertFile string `json:"client_cert_file"` // 客户端证书文件路径
	ClientKeyFile  string `json:"client_key_file"`  // 客户端私钥文件路径
}

// IsValid 是否有效
func (c *Cert) IsValid() bool {
	if c != nil {
		if c.Cafile == "" || c.ClientCertFile == "" || c.ClientKeyFile == "" {
			return false
		}
		return true
	}
	return false
}

// String 序列化成字符串
func (c Cert) String() string {
	body, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(body)
}

// MQTTxServer MQTT服务器信息
type MQTTxServer struct {
	Vendor   string `json:"vendor"`    // MQTT服务软件厂商
	Scheme   string `json:"scheme"`    // MQTT协议 tcp/tls(ssl)/ws
	Domain   string `json:"domain"`    // MQTT服务器地址
	IP       string `json:"ip"`        // MQTT服务器IP
	Port     string `json:"port"`      // MQTT服务器端口
	Cert     Cert   `json:"cert"`      // MQTT证书信息
	ClientID string `json:"client_id"` // MQTT客户端ID, GID_GW101@@@deviceid
	Username string `json:"username"`  // MQTT服务器用户名
	Password string `json:"password"`  // MQTT服务器密码
}

// SetVendor 设置MQTT服务软件厂商
func (m *MQTTxServer) SetVendor(vendor string) {
	if m != nil {
		m.Vendor = vendor
	}
}

// GetVendor 获取MQTT服务软件厂商
func (m *MQTTxServer) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

// SetScheme 设置MQTT协议
func (m *MQTTxServer) SetScheme(scheme string) {
	if m != nil {
		m.Scheme = scheme
	}
}

// GetScheme 获取MQTT协议
func (m *MQTTxServer) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

// SetDomain 设置MQTT服务器地址
func (m *MQTTxServer) SetDomain(domain string) {
	if m != nil {
		m.Domain = domain
	}
}

// GetDomain 获取MQTT服务器地址
func (m *MQTTxServer) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

// SetIP 设置MQTT服务器IP
func (m *MQTTxServer) SetIP(ip string) {
	if m != nil {
		m.IP = ip
	}
}

// GetIP 获取MQTT服务器IP
func (m *MQTTxServer) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

// SetPort 设置MQTT服务器端口
func (m *MQTTxServer) SetPort(port string) {
	if m != nil {
		m.Port = port
	}
}

// GetPort 获取MQTT服务器端口
func (m *MQTTxServer) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

// SetCert 设置MQTT证书信息
func (m *MQTTxServer) SetCert(cert Cert) {
	if m != nil {
		m.Cert = cert
	}
}

// GetCert 获取MQTT证书信息
func (m *MQTTxServer) GetCert() Cert {
	if m != nil {
		return m.Cert
	}
	return Cert{}
}

// SetClientID 设置MQTT客户端ID
func (m *MQTTxServer) SetClientID(clientID string) {
	if m != nil {
		m.ClientID = clientID
	}
}

// GetClientID 获取MQTT客户端ID
func (m *MQTTxServer) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

// SetUsername 设置MQTT服务器用户名
func (m *MQTTxServer) SetUsername(username string) {
	if m != nil {
		m.Username = username
	}
}

// GetUsername 获取MQTT服务器用户名
func (m *MQTTxServer) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

// SetPassword 设置MQTT服务器密码
func (m *MQTTxServer) SetPassword(password string) {
	if m != nil {
		m.Password = password
	}
}

// GetPassword 获取MQTT服务器密码
func (m *MQTTxServer) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// String 序列化成字符串
func (m *MQTTxServer) String() string {
	body, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(body)
}

// MQTTxClient MQTT客户端信息
type MQTTxClient struct {
	Vendor   string `json:"vendor"`    // MQTT服务软件厂商
	Scheme   string `json:"scheme"`    // MQTT协议 tcp/tls(ssl)/ws
	Domain   string `json:"domain"`    // MQTT服务器地址
	IP       string `json:"ip"`        // MQTT服务器IP
	Port     string `json:"port"`      // MQTT服务器端口
	Cert     Cert   `json:"cert"`      // MQTT证书信息
	ClientID string `json:"client_id"` // MQTT客户端ID, GID_GW101@@@deviceid
	Username string `json:"username"`  // MQTT服务器用户名
	Password string `json:"password"`  // MQTT服务器密码

	// MQTT连接相关
	Opts   *mqtt.ClientOptions `json:"-"` // MQTT连接参数
	Client mqtt.Client         `json:"-"` // MQTT客户端连接
}

// SetVendor 设置MQTT服务软件厂商
func (m *MQTTxClient) SetVendor(vendor string) {
	if m != nil {
		m.Vendor = vendor
	}
}

// GetVendor 获取MQTT服务软件厂商
func (m *MQTTxClient) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

// SetScheme 设置MQTT协议
func (m *MQTTxClient) SetScheme(scheme string) {
	if m != nil {
		m.Scheme = scheme
	}
}

// GetScheme 获取MQTT协议
func (m *MQTTxClient) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

// SetDomain 设置MQTT服务器地址
func (m *MQTTxClient) SetDomain(domain string) {
	if m != nil {
		m.Domain = domain
	}
}

// GetDomain 获取MQTT服务器地址
func (m *MQTTxClient) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

// SetIP 设置MQTT服务器IP
func (m *MQTTxClient) SetIP(ip string) {
	if m != nil {
		m.IP = ip
	}
}

// GetIP 获取MQTT服务器IP
func (m *MQTTxClient) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

// SetPort 设置MQTT服务器端口
func (m *MQTTxClient) SetPort(port string) {
	if m != nil {
		m.Port = port
	}
}

// GetPort 获取MQTT服务器端口
func (m *MQTTxClient) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

// SetCert 设置MQTT证书信息
func (m *MQTTxClient) SetCert(cert Cert) {
	if m != nil {
		m.Cert = cert
	}
}

// GetCert 获取MQTT证书信息
func (m *MQTTxClient) GetCert() Cert {
	if m != nil {
		return m.Cert
	}
	return Cert{}
}

// SetClientID 设置MQTT客户端ID
func (m *MQTTxClient) SetClientID(clientID string) {
	if m != nil {
		m.ClientID = clientID
	}
}

// GetClientID 获取MQTT客户端ID
func (m *MQTTxClient) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

// SetUsername 设置MQTT服务器用户名
func (m *MQTTxClient) SetUsername(username string) {
	if m != nil {
		m.Username = username
	}
}

// GetUsername 获取MQTT服务器用户名
func (m *MQTTxClient) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

// SetPassword 设置MQTT服务器密码
func (m *MQTTxClient) SetPassword(password string) {
	if m != nil {
		m.Password = password
	}
}

// GetPassword 获取MQTT服务器密码
func (m *MQTTxClient) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// Connect 连接MQTT服务器
func (m *MQTTxClient) Connect(defaultPublishHandler mqtt.MessageHandler, onConnectHandler mqtt.OnConnectHandler, connectionLostHandler mqtt.ConnectionLostHandler, reconnectingHandler mqtt.ReconnectHandler) error {
	if m == nil {
		return errors.New("MQTTxClient is nil")
	}

	if m.Opts == nil {
		m.Opts = mqtt.NewClientOptions()
		scheme := m.Scheme
		if scheme == "tls" {
			scheme = "ssl" // mqtt sdk 不识别 tls, 需要转换成 ssl
			if !m.Cert.IsValid() {
				return errors.New("MQTTxClient.Cert is invalid")
			}
		}
		m.Opts.AddBroker(scheme + "://" + m.Domain + ":" + m.Port)
		m.Opts.SetClientID(m.ClientID)
		m.Opts.SetUsername(m.Username)
		m.Opts.SetPassword(m.Password)
		if m.Scheme == "tls" {
			tlsConfig, err := NewTLSConfig(m.Cert.Cafile, m.Cert.ClientCertFile, m.Cert.ClientKeyFile)
			if err != nil {
				return err
			}
			m.Opts.SetTLSConfig(tlsConfig)
		}

		// 设置相关回调函数
		if defaultPublishHandler != nil {
			m.Opts.SetDefaultPublishHandler(defaultPublishHandler)
		}
		if onConnectHandler != nil {
			m.Opts.SetOnConnectHandler(onConnectHandler)
		}
		if connectionLostHandler != nil {
			m.Opts.SetConnectionLostHandler(connectionLostHandler)
		}
		if reconnectingHandler != nil {
			m.Opts.SetReconnectingHandler(reconnectingHandler)
		}
	}

	m.Client = mqtt.NewClient(m.Opts)
	if token := m.Client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

// Disconnect 断开MQTT服务器
func (m *MQTTxClient) Disconnect() error {
	if m == nil {
		return errors.New("MQTTxClient is nil")
	}
	if m.Client != nil {
		m.Client.Disconnect(250)
		m.Client = nil
	}
	return nil
}

// Publish 发布消息
func (m *MQTTxClient) Publish(topic string, qos byte, retained bool, payload interface{}) error {
	if m == nil {
		return errors.New("MQTTClient is nil")
	}
	if m.Client == nil {
		return errors.New("MQTTClient.Client is nil")
	}
	if token := m.Client.Publish(topic, qos, retained, payload); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

// Subscribe 订阅消息
func (m *MQTTxClient) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) error {
	if m == nil {
		return errors.New("MQTTClient is nil")
	}
	if m.Client == nil {
		return errors.New("MQTTClient.Client is nil")
	}
	if token := m.Client.Subscribe(topic, qos, callback); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

// String 序列化成字符串
func (m MQTTxClient) String() string {
	body, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(body)
}
