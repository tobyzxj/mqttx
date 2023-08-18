package mqttx

import "net/url"

// FormatServerAddr 转换 ssl:// 地址为 tls://
func FormatServerAddr(addr string) string {
	u, err := url.Parse(addr)
	if err != nil {
		return addr
	}
	if u.Scheme == "ssl" {
		return "tls://" + u.Host
	}
	return addr
}
