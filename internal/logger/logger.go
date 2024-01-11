package logger

import "fmt"

type Logger interface {
	Named(name string) Logger
	Error(args ...interface{})
	Errorw(msg string, args ...interface{})
	Warn(args ...interface{})
	Warnw(msg string, args ...interface{})
	Info(args ...interface{})
	Infow(msg string, args ...interface{})
	Debug(args ...interface{})
	Debugw(msg string, args ...interface{})
}

// A simple logger that prints to stdout.
type simpleLogger struct {
	name string
}

// You might want to use a more sophisticated logger.
// e.g. zap, sugar, ...
func New() Logger {
	return &simpleLogger{
		name: "",
	}
}

func (l *simpleLogger) Named(name string) Logger {
	return &simpleLogger{
		name: fmt.Sprintf("%s.%s", l.name, name),
	}
}

func (l *simpleLogger) Error(args ...interface{}) {
	fmt.Printf("[ERROR] %s: %s\n", l.name, fmt.Sprint(args...))
}

func (l *simpleLogger) Errorw(msg string, args ...interface{}) {
	fmt.Printf("[ERROR] %s: %s\n", l.name, fmt.Sprintf(msg, args...))
}

func (l *simpleLogger) Warn(args ...interface{}) {
	fmt.Printf("[WARN] %s: %s\n", l.name, fmt.Sprint(args...))
}

func (l *simpleLogger) Warnw(msg string, args ...interface{}) {
	fmt.Printf("[WARN] %s: %s\n", l.name, fmt.Sprintf(msg, args...))
}

func (l *simpleLogger) Info(args ...interface{}) {
	fmt.Printf("[INFO] %s: %s\n", l.name, fmt.Sprint(args...))
}

func (l *simpleLogger) Infow(msg string, args ...interface{}) {
	fmt.Printf("[INFO] %s: %s\n", l.name, fmt.Sprintf(msg, args...))
}

func (l *simpleLogger) Debug(args ...interface{}) {
	fmt.Printf("[DEBUG] %s: %s\n", l.name, fmt.Sprint(args...))
}

func (l *simpleLogger) Debugw(msg string, args ...interface{}) {
	fmt.Printf("[DEBUG] %s: %s\n", l.name, fmt.Sprintf(msg, args...))
}
