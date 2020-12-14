package gormlogrus

import (
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *gormLogger
	}{
		{
			name: "TestNew",
			want: &gormLogger{
				SkipErrRecordNotFound: true,
				ModuleName:            "gorm",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logger_LogMode(t *testing.T) {
	type fields struct {
		SlowThreshold         time.Duration
		SourceField           string
		SkipErrRecordNotFound bool
	}
	type args struct {
		in0 logger.LogLevel
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   logger.Interface
	}{
		{
			name:   "Test_logger_LogMode",
			fields: fields{},
			args:   args{},
			want:   &gormLogger{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gormLogger{
				SlowThreshold:         tt.fields.SlowThreshold,
				SourceField:           tt.fields.SourceField,
				SkipErrRecordNotFound: tt.fields.SkipErrRecordNotFound,
			}
			if got := l.LogMode(tt.args.in0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LogMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormLogger_Info(t *testing.T) {
	type fields struct {
		SlowThreshold         time.Duration
		SourceField           string
		SkipErrRecordNotFound bool
	}
	type args struct {
		ctx  context.Context
		s    string
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Test_gormLogger_Info",
			fields: fields{},
			args:   args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gormLogger{
				SlowThreshold:         tt.fields.SlowThreshold,
				SourceField:           tt.fields.SourceField,
				SkipErrRecordNotFound: tt.fields.SkipErrRecordNotFound,
			}
			l.Info(tt.args.ctx, tt.args.s, tt.args.args)
		})
	}
}

func Test_gormLogger_Warn(t *testing.T) {
	type fields struct {
		SlowThreshold         time.Duration
		SourceField           string
		SkipErrRecordNotFound bool
	}
	type args struct {
		ctx  context.Context
		s    string
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Test_gormLogger_Warn",
			fields: fields{},
			args:   args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gormLogger{
				SlowThreshold:         tt.fields.SlowThreshold,
				SourceField:           tt.fields.SourceField,
				SkipErrRecordNotFound: tt.fields.SkipErrRecordNotFound,
			}
			l.Warn(tt.args.ctx, tt.args.s, tt.args.args)
		})
	}
}

func Test_gormLogger_Error(t *testing.T) {
	type fields struct {
		SlowThreshold         time.Duration
		SourceField           string
		SkipErrRecordNotFound bool
	}
	type args struct {
		ctx  context.Context
		s    string
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Test_gormLogger_Error",
			fields: fields{},
			args:   args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gormLogger{
				SlowThreshold:         tt.fields.SlowThreshold,
				SourceField:           tt.fields.SourceField,
				SkipErrRecordNotFound: tt.fields.SkipErrRecordNotFound,
			}
			l.Error(tt.args.ctx, tt.args.s, tt.args.args)
		})
	}
}

func Test_gormLogger_Trace(t *testing.T) {
	type fields struct {
		SlowThreshold         time.Duration
		SourceField           string
		SkipErrRecordNotFound bool
	}
	type args struct {
		ctx   context.Context
		begin time.Time
		fc    func() (string, int64)
		err   error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test_gormLogger_Trace",
			fields: fields{
				SourceField: "test",
			},
			args: args{
				ctx:   nil,
				begin: time.Time{},
				fc: func() (string, int64) {
					return "", 0
				},
				err: nil,
			},
		},
		{
			name: "Test_gormLogger_TraceSlowThreshold",
			fields: fields{
				SourceField:   "test",
				SlowThreshold: 1,
			},
			args: args{
				ctx:   nil,
				begin: time.Time{},
				fc: func() (string, int64) {
					return "", 0
				},
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gormLogger{
				SlowThreshold:         tt.fields.SlowThreshold,
				SourceField:           tt.fields.SourceField,
				SkipErrRecordNotFound: tt.fields.SkipErrRecordNotFound,
			}
			l.Trace(tt.args.ctx, tt.args.begin, tt.args.fc, tt.args.err)
		})
	}
}

func Test_gormLogger_CreateLogrusFields(t *testing.T) {
	type fields struct {
		SlowThreshold         time.Duration
		SourceField           string
		SkipErrRecordNotFound bool
		ModuleName            string
	}
	type args struct {
		data []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   logrus.Fields
	}{
		{
			name:   "Test_gormLogger_CreateLogrusFields",
			fields: fields{},
			args: args{
				data: []interface{}{"test"},
			},
			want: logrus.Fields{
				"module": "gorm",
				"data":   []interface{}{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &gormLogger{
				SlowThreshold:         tt.fields.SlowThreshold,
				SourceField:           tt.fields.SourceField,
				SkipErrRecordNotFound: tt.fields.SkipErrRecordNotFound,
				ModuleName:            tt.fields.ModuleName,
			}
			if got := l.CreateLogrusFields(tt.args.data...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateLogrusFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
