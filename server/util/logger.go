package util

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger 系の実装は https://speakerdeck.com/hgsgtk/design-considerations-for-container-based-go-application を参考にさせていただいている。

// Writer は、ログのアウトプットの場所を指定する。
var Writer zapcore.WriteSyncer = os.Stdout

// Init は、zap の global logger をカスタム用に置き換える。
func Init(output zapcore.WriteSyncer) {
	logger := newLogger(output)
	zap.ReplaceGlobals(logger)
}

func newLogger(syncer zapcore.WriteSyncer) *zap.Logger {
	atom := zap.NewAtomicLevel()

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	bl := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(Writer),
		atom,
	))

	l := bl.With(zap.String("out", "stdout"))

	return l
}

// Logger は、zap.Logger を生成して、返す。
func Logger() *zap.Logger {
	return zap.L()
}
