package logger

import (
	"os"
	"strings"

	// project import

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	PrefixServerLog = "[SERVER]"
	//PrefixDataLog           = "[DATA]"
	//PrefixConnectionLog     = "[CONNECTION]"
	//PrefixDataProcessingLog = "[DATA_PROCESSING]"
)

type ZapLogger struct {
	logLevel    string
	environment string
	serviceName *string
}

func NewZapLogger(logLevel string, env string, opts ...Options) *ZapLogger {
	c := &ZapLogger{logLevel: logLevel, environment: env}

	for _, opt := range opts {
		opt(c)
	}
	return c
}

type Options func(*ZapLogger)

func WithServiceName(serviceName string) Options {
	return func(c *ZapLogger) {
		c.serviceName = &serviceName
	}
}

// Init a zap logger service
func (z *ZapLogger) Init() *zap.Logger {
	var encoder zapcore.Encoder
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	if z.environment != "default" {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	} else {
		// used to print human-readable log in local
		encoderCfg.CallerKey = "caller"
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	}

	level := z.setLogLevel()

	core := zapcore.NewCore(
		encoder,
		zapcore.Lock(os.Stdout),
		level,
	)

	logger := zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	if z.serviceName != nil && len(*z.serviceName) > 0 {
		logger = logger.With(zap.String("service", *z.serviceName))
	}

	return logger
}

func (z *ZapLogger) setLogLevel() zap.AtomicLevel {
	switch z.logLevel {
	case "INFO":
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "DEBUG":
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "WARNING":
		return zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "ERROR":
		return zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case "FATAL":
		return zap.NewAtomicLevelAt(zapcore.FatalLevel)
	}
	return zap.NewAtomicLevelAt(zapcore.InfoLevel)
}

func MaskStringSource(s string) string {
	n := len(s)
	if n <= 6 {
		return s // troppo corta, non mascherare
	}

	prefix := s[:3]
	suffix := s[n-3:]
	masked := strings.Repeat("x", n-6)

	return prefix + masked + suffix
}

func MaskEmail(email string) string {
	at := strings.Index(email, "@")
	if at <= 1 {
		return "***"
	}

	local := email[:at]
	domain := email[at+1:]

	// local part
	maskedLocal := local[:1] + strings.Repeat("*", len(local)-2) + local[len(local)-1:]

	// domain part
	dot := strings.LastIndex(domain, ".")
	if dot <= 0 {
		return maskedLocal + "@***"
	}

	domainName := domain[:dot]
	tld := domain[dot:]

	maskedDomain := domainName[:1] + strings.Repeat("*", len(domainName)-1)

	return maskedLocal + "@" + maskedDomain + tld
}
