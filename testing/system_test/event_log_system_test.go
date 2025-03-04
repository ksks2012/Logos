package system_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/logos/ecs/ecs_systems"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestLogEvent(t *testing.T) {
	var buf bytes.Buffer
	writer := zapcore.AddSync(&buf)
	encoderCfg := zap.NewProductionEncoderConfig()
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), writer, zapcore.InfoLevel)
	logger := zap.New(core)
	ecs_systems.SetLogger(logger)

	tests := []struct {
		time  string
		event string
	}{
		{"2023-10-01T12:00:00Z", "Test Event 1"},
		{"2023-10-02T13:00:00Z", "Test Event 2"},
	}

	for _, tt := range tests {
		ecs_systems.LogEvent(tt.time, tt.event)
		logged := buf.String()
		if !strings.Contains(logged, tt.time) || !strings.Contains(logged, tt.event) {
			t.Errorf("expected log to contain time %s and event %s, got %s", tt.time, tt.event, logged)
		}
		buf.Reset()
	}
}
