package gorm_logrus

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type logger struct {
	SlowThreshold time.Duration
	SourceField   string
}

func (l *logger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *logger) Info(ctx context.Context, s string, args ...interface{}) {
	log.Infof(s, args)
}

func (l *logger) Warn(ctx context.Context, s string, args ...interface{}) {
	log.Warnf(s, args)
}

func (l *logger) Error(ctx context.Context, s string, args ...interface{}) {
	log.Errorf(s, args)
}

func (l *logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := log.Fields{}
	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if err != nil {
		fields[log.ErrorKey] = err
		log.WithFields(fields).Errorf("%s [%s]", sql, elapsed)
		return
	}

	if l.SlowThreshold != 0  && elapsed > l.SlowThreshold {
		log.WithFields(fields).Warnf("%s [%s]", sql, elapsed)
		return
	}
	
	log.WithFields(fields).Debugf("%s [%s]", sql, elapsed)
}