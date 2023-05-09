package netWorkConnectionTestServer

import (
	"NetWorkConnectionTest/internal"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// 创建一个数据缓存对象
	dataCache := Internal.NewDataCache()

	// 创建一个Echo实例
	e := echo.New()

	// 添加中间件
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 设置开启指定协议端口监听
	e.POST("/set/post", func(c echo.Context) error {
		key := c.FormValue("key")
		value := c.FormValue("value")
		dataCache.Set(key, value)
		return c.String(200, "OK")
	})

	// 获取一个数据
	e.GET("/get", func(c echo.Context) error {
		key := c.QueryParam("key")
		value := dataCache.Get(key)
		if value == nil {
			return c.String(404, "Not Found")
		}
		return c.String(200, value.(string))
	})

	// 启动服务
	e.Logger.Fatal(e.Start(":8080"))
}
