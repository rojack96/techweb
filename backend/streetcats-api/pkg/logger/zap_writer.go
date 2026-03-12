package logger

import (
	"strings"

	"go.uber.org/zap"
)

// used to replace gin logger with Zap

type ZapGinWriter struct {
	Logger *zap.Logger
}

func (w *ZapGinWriter) Write(p []byte) (n int, err error) {
	line := strings.TrimSpace(string(p))
	if line == "" {
		return len(p), nil
	}

	// Parsing rudimentale del log di Gin: [GIN] 2025/11/18 - 10:12 | 200 | 1.2ms | 127.0.0.1 | GET "/"
	parts := strings.Split(line, "|")
	if len(parts) >= 5 {
		status := strings.TrimSpace(parts[1])
		latency := strings.TrimSpace(parts[2])
		ip := strings.TrimSpace(parts[3])
		methodPath := strings.TrimSpace(parts[4])

		// separa metodo e path
		method := ""
		path := ""
		if strings.Contains(methodPath, " ") {
			mp := strings.SplitN(methodPath, " ", 2)
			method = mp[0]
			path = strings.Trim(mp[1], `"`)
		}

		w.Logger.Info("HTTP Request",
			zap.String("status", status),
			zap.String("latency", latency),
			zap.String("ip", ip),
			zap.String("method", method),
			zap.String("path", path),
		)
	} else {
		// fallback: log generico
		w.Logger.Info(line)
	}

	return len(p), nil
}
