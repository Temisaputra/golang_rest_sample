package logger

import (
	"github.com/Temisaputra/warOnk/internal/infrastructure/config"
	"go.uber.org/zap"
	zapcore "go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

// NewLogger bikin logger baru
func NewLogger(config *config.Config) *zap.Logger {
	// Bisa juga pakai zap.NewProduction() kalau di server beneran
	cfg := zap.NewProductionConfig()
	cfg.Encoding = "json"
	cfg.OutputPaths = []string{"stdout"} // biar dikonsumsi docker log driver
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, _ := cfg.Build()
	return logger
}

func (l *Logger) Sync() {
	_ = l.Logger.Sync() // flush buffer
}

// Custom log dengan tambahan field request_id & status_code
func (l *Logger) WithRequest(requestID string, statusCode int) *Logger {
	return &Logger{
		Logger: l.With(
			zap.String("request_id", requestID),
			zap.Int("status_code", statusCode),
		).Sugar().Desugar(),
	}
}
