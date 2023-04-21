package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func init() {
	// Log, _ = zap.NewDevelopment()
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "Timestame"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	var err error
	Log, err = config.Build()

	defer Log.Sync()

	if err != nil {
		panic(err)
	}
}
