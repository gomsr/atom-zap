package zapx

import (
	"github.com/gomsr/atom-zap/configz"
	"testing"
)

func TestZap(t *testing.T) {
	zap := configz.Zap{
		Level:         "info",
		Format:        "json",
		Prefix:        "[kit-logger]",
		Director:      "log",
		ShowLine:      true,
		EncodeLevel:   "LowercaseColorLevelEncoder",
		StacktraceKey: "stacktrace",
		LogInConsole:  true,
	}

	LOGGER := Zap(zap)
	LOGGER.Info("test log output")
}
