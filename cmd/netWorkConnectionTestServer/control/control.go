package control

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

const (
	TCP  int = 1
	UDP  int = 2
	ICMP int = 3
)

type SetServiceRequest struct {
	protocol int
	port     int
}

// 解析Context转为SetService请求体
func ParseSetServiceRequest(c echo.Context) (*SetServiceRequest, error) {
	var setServiceREQ *SetServiceRequest = new(SetServiceRequest)
	var err error
	setServiceREQ.protocol, err = strconv.Atoi(c.Param("protocol"))
	if err == nil {
		setServiceREQ.port, err = strconv.Atoi(c.Param("port"))
		if err == nil {
			return setServiceREQ, err
		}
	}
	return nil, err
}

func SetService(c echo.Context) error {
	setServiceReuqest, err := ParseSetServiceRequest(c)
	if err != nil {
		return err
	}
	switch setServiceReuqest.protocol {
	case TCP:

	}
	return nil
}
