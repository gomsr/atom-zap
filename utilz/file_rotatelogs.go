package utilz

import (
	config "github.com/gomsr/atom-zap/configz"
	"os"

	"go.uber.org/zap/zapcore"
)

var FileRotatelogs = new(fileRotatelogs)

type fileRotatelogs struct{}

// GetWriteSyncer 获取 zapcore.WriteSyncer
func (r *fileRotatelogs) GetWriteSyncer(config config.Zap, level string) zapcore.WriteSyncer {
	fileWriter := NewCutter(config.Director, level, WithCutterFormat("2006-01-02"))

	if config.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	}

	return zapcore.AddSync(fileWriter)
}
