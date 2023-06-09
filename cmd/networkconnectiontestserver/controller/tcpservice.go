package controller

import (
	"errors"
	"fmt"
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
	return listenerTCPPort(ctx, port)
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

func listenerTCPPort(ctx echo.Context, port string) error {
	listener, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		ctx.Logger().Error("开启端口监听失败：", err.Error())
	}
	defer listener.Close()
	ctx.Logger().Info("Listening on 0.0.0.0:", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			ctx.Logger().Error("ERROR accenting : ", err.Error())
			return err
		}
		go handleConnection(ctx, conn)
	}
}

func handleConnection(ctx echo.Context, conn net.Conn) {
	buffer := make([]byte, 1024)
	length, err := conn.Read(buffer)
	if err != nil {
		ctx.Logger().Error("解析client请求失败：", err.Error())
		return
	}
	fmt.Println("Info : ", string(buffer[:length]))

	conn.Write([]byte("200"))
	conn.Close()
}
