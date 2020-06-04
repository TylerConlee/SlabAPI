package log

import (
	"os"

	"go.uber.org/zap"
)

type Logger struct{ e *zap.Logger }

var Log *Logger

func init() {
	var err error
	l, err := zap.NewProduction()
	Log = &Logger{l}

	if err != nil {
		os.Exit(1)
	}

	defer Log.e.Sync()
}
func (s *Logger) Fatal() {

	s.e.Fatal("Error")
}
func (s *Logger) Error(msg string, c map[string]interface{}) {
	s.e.Error(msg)
}
func (s *Logger) Warn(msg string, c map[string]interface{}) {
	s.e.Warn(msg)
}
func (s *Logger) Info(msg string, c map[string]interface{}) {
	s.e.Info(msg)
}
func (s *Logger) Debug(msg string, c map[string]interface{}) {
	s.e.Debug(msg)
}
