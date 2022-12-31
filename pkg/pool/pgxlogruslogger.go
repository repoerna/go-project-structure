package pool

import (
	"context"

	logrusadapter "github.com/jackc/pgx-logrus"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/sirupsen/logrus"
)

// TracingLogger does two things at once, logging and tracing
type PGXLogrusLogger struct {
	logger *logrusadapter.Logger
}

func NewPGXLogrusLogger(logger logrus.FieldLogger) *PGXLogrusLogger {
	return &PGXLogrusLogger{logger: logrusadapter.NewLogger(logger)}
}

func (l *PGXLogrusLogger) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]interface{}) {
	l.logger.Log(ctx, level, msg, data)
}
