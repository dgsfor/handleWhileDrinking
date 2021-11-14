package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sync/atomic"
	"time"
)

var running int64 = 0

func GinZap(logger *zap.Logger, timeFormat string, utc bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		atomic.AddInt64(&running, 1)
		defer func() {
			atomic.AddInt64(&running, -1)
		}()
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		host := c.Request.Host
		var userAgent string
		if gin.Mode() == "debug" {
			userAgent = ""
		} else {
			userAgent = c.Request.UserAgent()
		}
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}

		var errors string
		eVal, exists := c.Get("errors")
		if exists {
			errors = eVal.(string)
		} else {
			errors = ""
		}

		if len(c.Errors) > 0 {
			errors = errors + "\n" + c.Errors.String()
		}

		logger.Info("",
			zap.String("log_type", "access_log"),
			zap.Int("status", c.Writer.Status()),
			zap.String("host", host),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("remote_ip", c.ClientIP()),
			zap.String("user-agent", userAgent),
			zap.String("time", end.Format(timeFormat)),
			zap.Duration("duration", latency),
			zap.String("errors", errors),
			zap.String("request_id", c.Request.Header.Get("X-Request-Id")),
		)
	}
}
