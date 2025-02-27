//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=models.cfg.yaml ../openapi.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=server.cfg.yaml ../openapi.yaml

package api

import (
	"net/http"

	"github.com/KongAirlines/routes/api/models"
	"github.com/labstack/echo/v4"
)

type RouteService struct {
	Routes        []models.Route
	PrivateRoutes []models.Route
}

func NewRouteService() *RouteService {
	rv := RouteService{}
	rv.Routes = []models.Route{
		{Id: "LHR-JFK", Origin: "LHR", Destination: "JFK", AvgDuration: 470},
		{Id: "LHR-SFO", Origin: "LHR", Destination: "SFO", AvgDuration: 660},
		{Id: "LHR-DXB", Origin: "LHR", Destination: "DXB", AvgDuration: 420},
		{Id: "LHR-HKG", Origin: "LHR", Destination: "HKG", AvgDuration: 745},
		{Id: "LHR-BOM", Origin: "LHR", Destination: "BOM", AvgDuration: 540},
		{Id: "LHR-HND", Origin: "LHR", Destination: "HND", AvgDuration: 830},
		{Id: "LHR-CPT", Origin: "LHR", Destination: "CPT", AvgDuration: 700},
		{Id: "LHR-SYD", Origin: "LHR", Destination: "SYD", AvgDuration: 1320},
		{Id: "LHR-SIN", Origin: "LHR", Destination: "SIN", AvgDuration: 800},
		{Id: "LHR-LAX", Origin: "LHR", Destination: "LAX", AvgDuration: 675},
	}
	rv.PrivateRoutes = []models.Route{
		{Id: "VIP-LHR-JFK", Origin: "LHR", Destination: "VIP-JFK", AvgDuration: 430},
		{Id: "VIP-LHR-SFO", Origin: "LHR", Destination: "VIP-SFO", AvgDuration: 620},
		{Id: "VIP-LHR-DXB", Origin: "LHR", Destination: "VIP-DXB", AvgDuration: 390},
		{Id: "VIP-LHR-HKG", Origin: "LHR", Destination: "VIP-HKG", AvgDuration: 645},
	}
	return &rv
}

func (s *RouteService) GetHealthStatus(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"status": "OK"})
}

func (s *RouteService) GetAllRoutes(ctx echo.Context) error {
	if ctx.Request().Header.Get("x-vip") == "true" {
		allRoutes := append(s.Routes, s.PrivateRoutes...)
		return ctx.JSON(200, allRoutes)
	}

	return ctx.JSON(200, s.Routes)
}

func (s *RouteService) GetRouteById(ctx echo.Context, id string) error {
	routes := s.Routes
	if ctx.Request().Header.Get("x-vip") == "true" {
		routes = append(routes, s.PrivateRoutes...)
	}
	for _, route := range routes {
		if route.Id == id {
			err := ctx.JSON(200, route)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return ctx.JSON(404, nil)
}
