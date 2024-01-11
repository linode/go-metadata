package integration

import (
	"bytes"
	"fmt"
)

type testLogger struct {
	Data bytes.Buffer
}

func (l *testLogger) Errorf(format string, v ...interface{}) {
	l.Data.WriteString("[ERROR] " + fmt.Sprintf(format, v...))
}

func (l *testLogger) Warnf(format string, v ...interface{}) {
	l.Data.WriteString("[WARN] " + fmt.Sprintf(format, v...))
}

func (l *testLogger) Debugf(format string, v ...interface{}) {
	l.Data.WriteString("[DEBUG] " + fmt.Sprintf(format, v...))
}
