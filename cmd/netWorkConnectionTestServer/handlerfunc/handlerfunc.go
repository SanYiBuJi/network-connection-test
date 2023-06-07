package handlerfunc

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"io/ioutil"
)

// 日志中间件
//func logger(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(ctx echo.Context) error {
//		startTime := time.Now()
//
//		//执行action业务逻辑
//		if err := next(ctx); err != nil {
//			log.Logger.Error("用户访问日志, action返回错误", zap.Error(err))
//		}
//
//		//请求耗时
//		duration := time.Since(startTime)
//		formParams, _ := ctx.FormParams()
//		log.Logger.Info("用户访问日志", zap.String("uri", ctx.Request().URL.Path), zap.Any("method", ctx.Request().Method), zap.Any("queryParaList", ctx.QueryParams()), zap.Any("formParaList", formParams), zap.Any("header", ctx.Request().Header), zap.Any("userAgent", ctx.Request().UserAgent()), zap.Any("cookies", ctx.Request().Cookies()), zap.String("ip", ctx.RealIP()), zap.Any("ipLocation", ip_location.GetIpLocationString(ctx.RealIP())), zap.Any("输入字节数", ctx.Request().Header.Get(echo.HeaderContentLength)), zap.Any("响应字节数", ctx.Response().Size), zap.Any("duration", duration), zap.Any("durationNanosecond", int64(duration)))
//
//		return nil
//	}
//}

func PrintRequestInfo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		req := ctx.Request()
		//ctx.Logger().Info("Request: %s", req.Body)
		if req.Body != nil {
			body, err := ioutil.ReadAll(req.Body)
			if err == nil {
				ctx.Logger().Info("Request body: %s", string(body))
			}
			req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
		return next(ctx)
	}
}
