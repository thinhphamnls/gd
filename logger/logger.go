package gdlogger

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/thinhphamnls/gd/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	gormUtils "gorm.io/gorm/utils"
)

const (
	logTitle      = "[GORM] "
	sqlFormat     = logTitle + "%s"
	messageFormat = logTitle + "%s, %s"
	errorFormat   = logTitle + "%s, %s, %s"
	debugFormat   = logTitle + "%s, %s, %s"
	slowThreshold = 2000
)

type ILogger interface {
	Get() *zap.SugaredLogger
	LogMode(level gormLogger.LogLevel) gormLogger.Interface
	Info(ctx context.Context, msg string, data ...interface{})
	Warn(ctx context.Context, msg string, data ...interface{})
	Error(ctx context.Context, msg string, data ...interface{})
	Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error)
}

type logger struct {
	Zap *zap.SugaredLogger
}

func Init(cf gdconfig.Config) ILogger {
	zapLogger, err := build(cf)
	defer func() {
		_ = zapLogger.Sync()
	}()

	if err != nil {
		log.Fatalf("init zap logger failed: %v", err)
	}

	return &logger{Zap: zapLogger.Sugar()}
}

func build(cf gdconfig.Config) (*zap.Logger, error) {
	env := cf.Server.Env

	// configs default
	cfg := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			TimeKey:      "time",
			LevelKey:     "level",
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
			EncodeLevel: func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
				encoder.AppendString("[" + level.CapitalString() + "]")
			},
			EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
				encoder.AppendString(t.Format(time.RFC3339))
			},
		},
	}

	if env.Mode != gdconfig.ProductionEnv {
		cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}

	zapLogger, err := cfg.Build()
	return zapLogger, err
}

func (l logger) Get() *zap.SugaredLogger {
	return l.Zap
}

func (l logger) LogMode(_ gormLogger.LogLevel) gormLogger.Interface {
	return l
}

func (l logger) Info(_ context.Context, msg string, data ...interface{}) {
	l.Zap.Infof(messageFormat, append([]interface{}{msg, gormUtils.FileWithLineNum()}, data...)...)
}

func (l logger) Warn(_ context.Context, msg string, data ...interface{}) {
	l.Zap.Warnf(messageFormat, append([]interface{}{msg, gormUtils.FileWithLineNum()}, data...)...)
}

func (l logger) Error(_ context.Context, msg string, data ...interface{}) {
	l.Zap.Errorf(messageFormat, append([]interface{}{msg, gormUtils.FileWithLineNum()}, data...)...)
}

func (l logger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)

	switch {
	case err != nil:
		sql, _ := fc()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			l.Get().Debugf(debugFormat, gormUtils.FileWithLineNum(), err, sql)
		} else {
			l.Get().Errorf(errorFormat, gormUtils.FileWithLineNum(), err, sql)
		}
	case elapsed > slowThreshold*time.Millisecond:
		sql, _ := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", slowThreshold)
		l.Get().Warnf(errorFormat, gormUtils.FileWithLineNum(), slowLog, sql)
	default:
		sql, _ := fc()
		l.Get().Debugf(sqlFormat, sql)
	}
}
