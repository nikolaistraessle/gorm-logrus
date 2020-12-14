package gormlogrus

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

type gormLogger struct {
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
	ModuleName            string
}

func New() *gormLogger {
	return &gormLogger{
		SkipErrRecordNotFound: true,
		ModuleName:            "gorm",
	}
}

func (l *gormLogger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *gormLogger) Info(ctx context.Context, s string, args ...interface{}) {
	log.WithContext(ctx).WithFields(l.CreateLogrusFields(args)).Infof(s)
}

func (l *gormLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	log.WithContext(ctx).WithFields(l.CreateLogrusFields(args)).Warnf(s)
}

func (l *gormLogger) Error(ctx context.Context, s string, args ...interface{}) {
	log.WithContext(ctx).WithFields(l.CreateLogrusFields(args)).Errorf(s)
}

func (l *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := []interface{}{}
	if l.SourceField != "" {
		fields = append(fields, utils.FileWithLineNum())
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		log.WithContext(ctx).WithFields(l.CreateLogrusFields(fields...)).WithField(log.ErrorKey, err).Errorf("%s [%s]", sql, elapsed)
		return
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		log.WithContext(ctx).WithFields(l.CreateLogrusFields(fields...)).Warnf("%s [%s]", sql, elapsed)
		return
	}

	log.WithContext(ctx).WithFields(l.CreateLogrusFields(fields...)).Debugf("%s [%s]", sql, elapsed)
}

func (l *gormLogger) CreateLogrusFields(data ...interface{}) log.Fields {
	return log.Fields{
		"module": "gorm",
		"data":   data,
	}
}
