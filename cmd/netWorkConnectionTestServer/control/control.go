package control

import (
	"errors"
	"github.com/labstack/echo/v4"
)

const (
	TCP  string = "TCP"
	UDP  string = "UDP"
	ICMP string = "ICMP"
)

type SetServiceRequest struct {
	Protocol string `json:"protocol"`
	Port     string `json:"port"`
}

// ParseSetServiceRequest 解析Context转为SetService请求体
//func ParseSetServiceRequest(c echo.Context) (*SetServiceRequest, error) {
//	var setServiceREQ *SetServiceRequest = new(SetServiceRequest)
//	var err error
//	setServiceREQ.protocol, err = strconv.Atoi(c.Param("protocol"))
//	if err == nil {
//		setServiceREQ.port, err = strconv.Atoi(c.Param("port"))
//		if err == nil {
//			return setServiceREQ, err
//		}
//	}
//	return nil, err
//}

func SetService(ctx echo.Context) error {
	setServiceRequest := new(SetServiceRequest)
	if err := ctx.Bind(setServiceRequest); err != nil {
		return err
	}

	switch setServiceRequest.Protocol {
	case TCP:
		return TCPServiceSet(ctx, setServiceRequest.Port)
	case UDP:
		return nil
	case ICMP:
		return nil
	default:
		return errors.New("[error] Protocol error in request : " + setServiceRequest.Protocol)
	}
}
