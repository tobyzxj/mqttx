package mqttx

import (
	"testing"
)

type Args struct {
	In  string
	Out string
}

func TestFormatServerAddr(t *testing.T) {
	var tests = []Args{
		{"ssl://mqtt.test.com", "tls://mqtt.test.com"},
		{"tcp://mqtt.test.com", "tcp://mqtt.test.com"},
		{"ssl://mqtt.test.com:1883", "tls://mqtt.test.com:1883"},
		{"tcp://mqtt.test.com:1883", "tcp://mqtt.test.com:1883"},
		{"ssl://mqtt.test.com:8883", "tls://mqtt.test.com:8883"},
		{"tcp://mqtt.test.com:8883", "tcp://mqtt.test.com:8883"},
		{"ssl://mqtt.test.com:1883/abc", "tls://mqtt.test.com:1883"},
	}

	for _, test := range tests {
		if out := FormatServerAddr(test.In); out != test.Out {
			t.Errorf("FormatServerAddr(%v) = %v", test.In, out)
		}
	}
}
