package roach

import (
	"context"
	"fmt"
	"go-boilerplate-api/internal/config"
	"go-boilerplate-api/pkg/pool"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PreparePoolConnection(ctx context.Context, cfg *config.Configurations) (*pgxpool.Pool, error) {
	pgxLogger := pool.NewPGXLogrusLogger(&log.Logger{})

	pgxLogLevel, err := pool.LogLevelFromEnv(cfg.LogLevel)
	if err != nil {
		log.Warn(err)
		pgxLogLevel = tracelog.LogLevelDebug
	}

	// example dsn - "postgresql://root@localhost:26257/photos?sslmode=disable",
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?application_name=$ %s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.AppName,
	)

	if cfg.AppEnv != "production" {
		dsn = dsn + "sslmode=disable"
	}

	return pool.NewPGXPool(ctx, dsn, pgxLogger, pgxLogLevel)

}

func Init(cfg *config.Configurations) (*gorm.DB, error) {
	// example dsn - "postgresql://root@localhost:26257/photos?sslmode=disable",
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?application_name=$ %s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.AppName,
	)

	if cfg.AppEnv != "production" {
		dsn = dsn + "sslmode=disable"
	}

	gormConf := gorm.Config{}

	return gorm.Open(postgres.Open(dsn), &gormConf)
}

// type logger struct {
// 	SlowThreshold         time.Duration
// 	SourceField           string
// 	SkipErrRecordNotFound bool
// 	Debug                 bool
// }

// func New() *logger {
// 	return &logger{
// 		SkipErrRecordNotFound: true,
// 		Debug:                 true,
// 	}
// }

// func (l *logger) LogMode(gorm.LogLevel) gormlogger.Interface {
// 	return l
// }

// func (l *logger) Info(ctx context.Context, s string, args ...interface{}) {
// 	log.WithContext(ctx).Infof(s, args)
// }

// func (l *logger) Warn(ctx context.Context, s string, args ...interface{}) {
// 	log.WithContext(ctx).Warnf(s, args)
// }

// func (l *logger) Error(ctx context.Context, s string, args ...interface{}) {
// 	log.WithContext(ctx).Errorf(s, args)
// }

// func (l *logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
// 	elapsed := time.Since(begin)
// 	sql, _ := fc()
// 	fields := log.Fields{}
// 	if l.SourceField != "" {
// 		fields[l.SourceField] = utils.FileWithLineNum()
// 	}
// 	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
// 		fields[log.ErrorKey] = err
// 		log.WithContext(ctx).WithFields(fields).Errorf("%s [%s]", sql, elapsed)
// 		return
// 	}

// 	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
// 		log.WithContext(ctx).WithFields(fields).Warnf("%s [%s]", sql, elapsed)
// 		return
// 	}

// 	if l.Debug {
// 		log.WithContext(ctx).WithFields(fields).Debugf("%s [%s]", sql, elapsed)
// 	}
// }
