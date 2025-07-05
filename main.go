package log4Go

import (
	"log"
	"os"
)

const (
	ColorReset   = "\033[0m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorWhite   = "\033[37m"
)

type LogLevel int

const (
	Info LogLevel = iota
	Success
	Error
	Warning
)

type Logger struct {
	logLevel LogLevel
	logger   *log.Logger
}

func NewLogger(logLevel LogLevel) *Logger {
	return &Logger{
		logLevel: logLevel,
		logger:   log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (l *Logger) Log(level LogLevel, message string) {
	switch level {
	case Info:
		if l.logLevel <= Info {
			l.logger.Printf("%s[*] INFO: %s%s", ColorCyan, message, ColorReset)
		}
	case Success:
		if l.logLevel <= Success {
			l.logger.Printf("%s[+] SUCCESS: %s%s", ColorGreen, message, ColorReset)
		}
	case Error:
		if l.logLevel <= Error {
			l.logger.Printf("%s[!] ERROR: %s%s", ColorRed, message, ColorReset)
		}
	case Warning:
		if l.logLevel <= Warning {
			l.logger.Printf("%s[?] WARNING: %s%s", ColorYellow, message, ColorReset)
		}
	default:
		l.logger.Printf("%s[*] %s%s", ColorWhite, message, ColorReset)
	}
}
