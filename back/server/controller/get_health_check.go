// Package controller handles requests and returns responses. api_gen users directly edit this file.
// generated version: 2.1.0
// nolint:dupl
package controller

import (
	"github.com/labstack/echo/v4"
	types "github.com/star-integrations/project-boilerplate/back/interfaces"
	props "github.com/star-integrations/project-boilerplate/back/server/props"
)

// GetHealthCheckController ...
type GetHealthCheckController interface {
	GetHealthCheck(c echo.Context, req *types.GetHealthCheckRequest) (res *types.GetHealthCheckResponse, err error)
}

type getHealthCheckController struct {
	*props.ControllerProps
}

// NewGetHealthCheckController ...
func NewGetHealthCheckController(cp *props.ControllerProps) GetHealthCheckController {
	return &getHealthCheckController{
		ControllerProps: cp,
	}
}

// GetHealthCheck - GET health_check
func (ctrl *getHealthCheckController) GetHealthCheck(
	c echo.Context, req *types.GetHealthCheckRequest,
) (res *types.GetHealthCheckResponse, err error) {
	// return nil, apierror.NewAPIError(http.StatusBadRequest)
	//
	// return nil, apierror.NewAPIError(http.StatusBadRequest).SetError(err)
	//
	// body := map[string]interface{}{
	// 	"code": http.StatusBadRequest,
	// 	"message": "invalid request parameter.",
	// }
	// return nil, apierror.NewAPIError(http.StatusBadRequest, body).SetError(err)
	panic("require implements.") // FIXME require implements.
}

// AutoBind - use echo.Bind
func (ctrl *getHealthCheckController) AutoBind() bool {
	return true
}
