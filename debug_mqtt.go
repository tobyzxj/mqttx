package mqttx

type MQTTLibDebug struct {
	Level int
}

func mqttLibDebugNew(level int) *MQTTLibDebug {
	return &MQTTLibDebug{Level: level}
}

func (l *MQTTLibDebug) Println(v ...interface{}) {
	if l.Level <= LevelError {
		TraceError("%v", v...)
		return
	}
	if l.Level <= LevelInformational {
		TraceInfo("%v", v...)
		return
	}
	if l.Level <= LevelDebug {
		TraceDebug("%v", v...)
		return
	}
}

func (l *MQTTLibDebug) Printf(format string, v ...interface{}) {
	if l.Level <= LevelError {
		TraceError("%v", v...)
		return
	}
	if l.Level <= LevelInformational {
		TraceInfo("%v", v...)
		return
	}
	if l.Level <= LevelDebug {
		TraceDebug("%v", v...)
		return
	}
}
