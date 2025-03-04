package ecs_systems

import (
	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()

// LogEvent logs an event with a timestamp.
//
// Parameters:
//   - gameTime: A string representing the time of the event.
//   - event: A string describing the event.
//
// Example:
//
//	LogEvent("Game time 123", "User logged in")

func LogEvent(gameTime string, event string) {
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infow("Event", "message", "["+gameTime+"] "+event)
}

func SetLogger(customLogger *zap.Logger) {
	logger = customLogger
}
