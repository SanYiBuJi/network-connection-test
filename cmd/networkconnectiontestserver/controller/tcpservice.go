package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net"
	"strconv"
	"time"
)

func TCPServiceSet(ctx echo.Context, port string) error {
	postValue, err := strconv.Atoi(port)
	if err != nil {
		ctx.Logger().Error("Port resolution failure")
		return err
	}
	if err = checkPorts(ctx, port); err != nil {
		return err
	}
	ctx.Logger().Infof("Listening on TCP post：%d", postValue)
	return nil
}

func checkPorts(ctx echo.Context, ipPort string) error {
	conn, err := net.DialTimeout("tcp", ipPort, 3*time.Second)
	if err != nil {
		ctx.Logger().Info(ipPort, "-连接探测失败，端口未开启(fail)!")
		return nil
	} else {
		if conn != nil {
			ctx.Logger().Error(ipPort, "端口已开启(success)!")
			conn.Close()
			return errors.New("the port is occupied")
		} else {
			ctx.Logger().Info(ipPort, "端口未开启(fail)!")
			return nil
		}
	}

}
