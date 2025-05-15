package zapx

import (
	"fmt"
	"github.com/gomsr/atom-zap/configz"
	"github.com/gomsr/atom-zap/utilz"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap 获取 zapx.Logger
func Zap(config configz.Zap) (logger *zap.Logger) {
	if ok, _ := utilz.PathExists(config.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", config.Director)
		_ = os.Mkdir(config.Director, os.ModePerm)
	}

	cores := utilz.Zap.GetZapCores(config)
	logger = zap.New(zapcore.NewTee(cores...))

	if config.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}

func NewZap() (logger *zap.Logger) {
	return Zap(configz.Zap{
		Format: "json",
	})
}

//func InitZap() *zap.Logger {
//	kg.L = Zap(kg.C.Zap)
//	zap.ReplaceGlobals(kg.L)
//	return kg.L
//}
