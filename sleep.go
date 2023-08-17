package mqttx

import (
	"time"
)

// SleepSecond 睡眠指定的秒数
func SleepSecond(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}
