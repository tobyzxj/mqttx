package mqttx

import "fmt"

type MQTTLibDebug struct {
	Level int
}

func mqttLibDebugNew(level int) *MQTTLibDebug {
	return &MQTTLibDebug{Level: level}
}

func (l *MQTTLibDebug) Println(v ...interface{}) {
	if l.Level <= LevelError {
		TraceError("%v", fmt.Sprint(v...))
		return
	}
	if l.Level <= LevelInformational {
		TraceInfo("%v", fmt.Sprint(v...))
		return
	}
	if l.Level <= LevelDebug {
		TraceDebug("%v", fmt.Sprint(v...))
		return
	}
}

func (l *MQTTLibDebug) Printf(format string, v ...interface{}) {
	if l.Level <= LevelError {
		TraceError(format, v...)
		return
	}
	if l.Level <= LevelInformational {
		TraceInfo(format, v...)
		return
	}
	if l.Level <= LevelDebug {
		TraceDebug(format, v...)
		return
	}
}
