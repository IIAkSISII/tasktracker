package logger

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	log *logrus.Logger
}

func NewLogrusLogger(level, format string) (Logger, error) {
	l := logrus.New()

	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return nil, fmt.Errorf("invalid log level %q: %w", level, err)
	}
	l.SetLevel(lvl)

	switch format {
	case "json":
		l.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		l.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	default:
		return nil, fmt.Errorf("invalid log format %q", format)
	}

	return logrusLogger{log: l}, nil
}

func (l logrusLogger) Debug(ctx context.Context, msg string, fields ...Field) {
	l.log.WithFields(toLogrusFields(fields)).Debug(msg)
}

func (l logrusLogger) Info(ctx context.Context, msg string, fields ...Field) {
	l.log.WithFields(toLogrusFields(fields)).Info(msg)
}

func (l logrusLogger) Error(ctx context.Context, msg string, fields ...Field) {
	l.log.WithFields(toLogrusFields(fields)).Error(msg)
}

func (l logrusLogger) Warn(ctx context.Context, msg string, fields ...Field) {
	l.log.WithFields(toLogrusFields(fields)).Warn(msg)
}

func toLogrusFields(fields []Field) logrus.Fields {
	lf := make(logrus.Fields, len(fields))
	for _, f := range fields {
		lf[f.Key] = f.Value
	}
	return lf
}
