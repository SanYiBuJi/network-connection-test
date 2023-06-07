package controller

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
