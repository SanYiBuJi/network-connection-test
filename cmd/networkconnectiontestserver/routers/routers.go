package routers

import (
	"NetWorkConnectionTest/cmd/networkconnectiontestserver/controller"
	"github.com/labstack/echo/v4"
)

func SetRoutersTable(e *echo.Echo) {
	e.POST("/set/port", controller.SetService)

}
