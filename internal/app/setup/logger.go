package setup

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"goApi/internal/app/config"
	"goApi/pkg/logger"
	"os"
	"path/filepath"
	"time"
)

func Logger() (func(), error) {
	c := config.C.Log
	logger.SetLevel(logger.Level(c.Level))
	logger.SetFormatter(c.Format)

	var file *rotatelogs.RotateLogs
	switch c.Output {
	case "stdout":
		logger.SetOutput(os.Stdout)
	case "stderr":
		logger.SetOutput(os.Stderr)
	case "file":
		if name := c.OutputFile; name != "" {
			_ = os.MkdirAll(filepath.Dir(name), 0777)
			f, err := rotatelogs.New(name+".%Y-%m-%d",
				rotatelogs.WithLinkName(name),
				rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
				rotatelogs.WithRotationCount(uint(20)))
			if err != nil {
				return nil, err
			}
			logger.SetOutput(f)
			file = f
		}
	}

	//TODO 钩子处理

	return func() {
		if file != nil {
			_ = file.Close()
		}
	}, nil
}
