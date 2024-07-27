package logger

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Info(msg string) {}

func (l *Logger) Infof(msg string, args ...any) {}

func (l *Logger) Error(msg string) {}

func (l *Logger) Errorf(msg string, args ...any) {}

func (l *Logger) Warn(msg string) {}

func (l *Logger) Fatal(msg string) {}
