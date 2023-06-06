package control

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net"
	"time"
)

func TCPServiceSet(c echo.Context, port int) error {

}

func checkPorts(ipPorts []string) {
	//now := time.Now().Format("2006-01-02 15:04:05")
	for _, ipPort := range ipPorts {
		// 检测端口
		conn, err := net.DialTimeout("tcp", ipPort, 3*time.Second)
		if err != nil {
			//fmt.Println("["+now+"]", ipPort, "端口未开启(fail)!")
			echo.Logger("")
		} else {
			if conn != nil {
				fmt.Println("["+now+"]", ipPort, "端口已开启(success)!")
				conn.Close()
			} else {
				fmt.Println("["+now+"]", ipPort, "端口未开启(fail)!")
			}
		}
	}
}
