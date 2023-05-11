package logadapter

import (
	"fmt"

	"github.com/logx-go/commons/pkg/commons"
	"github.com/logx-go/contract/pkg/logx"
)

var _ logx.Logger = (*LogAdapter)(nil)
var _ logx.Adapter = (*LogAdapter)(nil)

// New returns a pointer to a new instance of LogAdapter
func New(logger logx.Log) logx.Adapter {
	return &LogAdapter{
		logger:    logger,
		fields:    make(map[string]any),
		formatter: nil,
	}
}

// LogAdapter implementation to wrap a format Logger
type LogAdapter struct {
	logger    logx.Log
	formatter logx.Formatter
	fields    map[string]any
}

func (s *LogAdapter) clone() *LogAdapter {
	return &LogAdapter{
		logger:    s.logger,
		fields:    s.fields,
		formatter: s.formatter,
	}
}

func (s *LogAdapter) format(v ...any) string {
	if len(v) < 1 {
		if s.formatter == nil {
			return ""
		}

		return s.formatter.Format("", s.fields)
	}

	msg := fmt.Sprintf(`%v`, v[0])
	fields := s.fields

	for i := 1; i < len(v); i += 2 {
		fieldName := ""
		if i < len(v) {
			fieldName = fmt.Sprintf(`%v`, v[i])
		}

		var fieldValue any
		if i+1 < len(v) {
			fieldValue = v[i+1]
		}

		fields[fieldName] = fieldValue
	}

	if s.formatter == nil {
		return msg
	}

	return s.formatter.Format(msg, fields)
}

func (s *LogAdapter) Fatal(v ...any) {
	c := s.clone()
	c.fields = commons.SetCallerInfo(1, false, c.fields, logx.FieldNameCallerFunc, logx.FieldNameCallerFile, logx.FieldNameCallerLine)
	c.logger.Fatal(c.format(v...))
}

func (s *LogAdapter) Panic(v ...any) {
	c := s.clone()
	c.fields = commons.SetCallerInfo(1, false, c.fields, logx.FieldNameCallerFunc, logx.FieldNameCallerFile, logx.FieldNameCallerLine)
	c.logger.Panic(c.format(v...))
}

func (s *LogAdapter) Print(v ...any) {
	c := s.clone()
	c.fields = commons.SetCallerInfo(1, false, c.fields, logx.FieldNameCallerFunc, logx.FieldNameCallerFile, logx.FieldNameCallerLine)
	c.logger.Print(c.format(v...))
}

func (s *LogAdapter) Fatalf(format string, v ...any) {
	c := s.clone()
	c.fields = commons.SetCallerInfo(1, false, c.fields, logx.FieldNameCallerFunc, logx.FieldNameCallerFile, logx.FieldNameCallerLine)
	c.logger.Fatal(c.format(fmt.Sprintf(format, v...)))
}

func (s *LogAdapter) Panicf(format string, v ...any) {
	c := s.clone()
	c.fields = commons.SetCallerInfo(1, false, c.fields, logx.FieldNameCallerFunc, logx.FieldNameCallerFile, logx.FieldNameCallerLine)
	c.logger.Panic(c.format(fmt.Sprintf(format, v...)))
}

func (s *LogAdapter) Printf(format string, v ...any) {
	c := s.clone()
	c.fields = commons.SetCallerInfo(1, false, c.fields, logx.FieldNameCallerFunc, logx.FieldNameCallerFile, logx.FieldNameCallerLine)
	c.logger.Print(c.format(fmt.Sprintf(format, v...)))
}

func (s *LogAdapter) Debug(v ...any) {
	c := s.clone()
	c.fields = commons.SetCallerInfo(1, false, c.fields, logx.FieldNameCallerFunc, logx.FieldNameCallerFile, logx.FieldNameCallerLine)
	c.WithField(logx.FieldNameLogLevel, logx.LogLevelDebug).Print(v...)
}

func (s *LogAdapter) Info(v ...any) {
	c := s.clone()
	c.fields = commons.SetCallerInfo(1, false, c.fields, logx.FieldNameCallerFunc, logx.FieldNameCallerFile, logx.FieldNameCallerLine)
	c.WithField(logx.FieldNameLogLevel, logx.LogLevelInfo).Print(v...)
}

func (s *LogAdapter) Notice(v ...any) {
	c := s.clone()
	c.fields = commons.SetCallerInfo(1, false, c.fields, logx.FieldNameCallerFunc, logx.FieldNameCallerFile, logx.FieldNameCallerLine)
	c.WithField(logx.FieldNameLogLevel, logx.LogLevelNotice).Print(v...)
}

func (s *LogAdapter) Warning(v ...any) {
	c := s.clone()
	c.fields = commons.SetCallerInfo(1, false, c.fields, logx.FieldNameCallerFunc, logx.FieldNameCallerFile, logx.FieldNameCallerLine)
	c.WithField(logx.FieldNameLogLevel, logx.LogLevelWarning).Print(v...)
}

func (s *LogAdapter) Error(v ...any) {
	c := s.clone()
	c.fields = commons.SetCallerInfo(1, false, c.fields, logx.FieldNameCallerFunc, logx.FieldNameCallerFile, logx.FieldNameCallerLine)
	c.WithField(logx.FieldNameLogLevel, logx.LogLevelError).Print(v...)
}

func (s *LogAdapter) Debugf(format string, v ...any) {
	c := s.clone()
	c.fields = commons.SetCallerInfo(1, false, c.fields, logx.FieldNameCallerFunc, logx.FieldNameCallerFile, logx.FieldNameCallerLine)
	c.WithField(logx.FieldNameLogLevel, logx.LogLevelDebug).Printf(format, v...)
}

func (s *LogAdapter) Infof(format string, v ...any) {
	c := s.clone()
	c.fields = commons.SetCallerInfo(1, false, c.fields, logx.FieldNameCallerFunc, logx.FieldNameCallerFile, logx.FieldNameCallerLine)
	c.WithField(logx.FieldNameLogLevel, logx.LogLevelInfo).Printf(format, v...)
}

func (s *LogAdapter) Noticef(format string, v ...any) {
	c := s.clone()
	c.fields = commons.SetCallerInfo(1, false, c.fields, logx.FieldNameCallerFunc, logx.FieldNameCallerFile, logx.FieldNameCallerLine)
	c.WithField(logx.FieldNameLogLevel, logx.LogLevelNotice).Printf(format, v...)
}

func (s *LogAdapter) Warningf(format string, v ...any) {
	c := s.clone()
	c.fields = commons.SetCallerInfo(1, false, c.fields, logx.FieldNameCallerFunc, logx.FieldNameCallerFile, logx.FieldNameCallerLine)
	c.WithField(logx.FieldNameLogLevel, logx.LogLevelWarning).Printf(format, v...)
}

func (s *LogAdapter) Errorf(format string, v ...any) {
	c := s.clone()
	c.fields = commons.SetCallerInfo(1, false, c.fields, logx.FieldNameCallerFunc, logx.FieldNameCallerFile, logx.FieldNameCallerLine)
	c.WithField(logx.FieldNameLogLevel, logx.LogLevelError).Printf(format, v...)
}

func (s *LogAdapter) WithField(name string, value any) logx.Logger {
	c := s.clone()
	c.fields[name] = value

	return c
}

func (s *LogAdapter) WithFormatter(formatter logx.Formatter) logx.Adapter {
	c := s.clone()
	c.formatter = formatter

	return c
}
