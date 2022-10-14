package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.POST("/hello/:id", handleHello)

	if err := e.Start(":5000"); err != nil {
		e.Logger.Error(err.Error())
	}
}

func handleHello(c echo.Context) error {
	var (
		pathParams struct {
			ID int32 `param:"id"`
		}
		bodyParams struct {
			ChildList []int32 `form:"child_list" json:"child_list"`
		}
		queryParams struct {
			Name *string    `query:"name"`
			Time *time.Time `query:"time"`
		}
	)
	if err := (&echo.DefaultBinder{}).BindPathParams(c, &pathParams); err != nil {
		c.Echo().Logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"message": "매개변수 변환 실패"})
	}
	if err := (&echo.DefaultBinder{}).BindBody(c, &bodyParams); err != nil {
		c.Echo().Logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"message": "매개변수 변환 실패"})
	}
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &queryParams); err != nil {
		c.Echo().Logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"message": "매개변수 변환 실패"})
	}
	return c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"pathParams":  pathParams,
			"bodyParams":  bodyParams,
			"queryParams": queryParams,
		},
	)
}

