package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Logger 日志中间件
func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			method := c.Request().Method
			path := c.Request().URL.Path

			// 记录请求开始
			zap.L().Debug("-> 请求开始",
				zap.String("method", method),
				zap.String("path", path),
			)

			// 处理请求
			err := next(c)

			// 计算耗时
			latency := time.Since(start)

			// 记录请求完成
			zap.L().Info("<- 请求完成",
				zap.String("method", method),
				zap.String("path", path),
				zap.Int("status", c.Response().Status),
				zap.Duration("latency", latency),
				zap.String("client_ip", c.RealIP()),
			)
			return err
		}
	}
}
