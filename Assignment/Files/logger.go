package main

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Log(data any) {}

func (l *Logger) Info(data any) {}

func (l *Logger) Error(data any) {}
