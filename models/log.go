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
	// log level
	Level string
	// log message
	Message string `gorm:"size:10240"`
}

// GetLevel of log
func (l *Log) GetLevel() LogLevel {
	return LogLevel(l.Level)
}
