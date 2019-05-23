package models

// LogLevel type
type LogLevel string

const (
	// LogLevelError enum
	LogLevelError = LogLevel("error")
	// LogLevelInfo enum
	LogLevelInfo = LogLevel("info")
	// LogLevelWarn enum
	LogLevelWarn = LogLevel("warn")
)

// Log type
type Log struct {
	BaseModel
	// the sequence of log
	Sequence int
	// log level
	Level string
}

// GetLevel of log
func (l *Log) GetLevel() LogLevel {
	return LogLevel(l.Level)
}
