package monitor

import (
	"net/http"

	"github.com/arl/statsviz"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// RegisterRoutes 注册 statsviz 运行时监控路由
func RegisterRoutes(e *echo.Echo, path string) {
	mux := http.NewServeMux()
	err := statsviz.Register(mux)
	if err != nil {
		zap.L().Error("failed to register statsviz", zap.Error(err))
		return
	}

	e.GET(path+"/*", echo.WrapHandler(mux))
	zap.L().Info("statsviz runtime monitor registered", zap.String("path", path))
}
