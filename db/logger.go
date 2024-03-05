package db

import (
	"xorm.io/core"
	log "github.com/sirupsen/logrus"
)

type Logger struct {
	*log.Entry
	level core.LogLevel
}

func (l *Logger) SetLevel(level core.LogLevel) {
	l.level = level
}

func (l *Logger) Level() core.LogLevel {
	return l.level
}

func (l *Logger) ShowSQL(show ...bool) {}
func (l *Logger) IsShowSQL() bool      { return false }
