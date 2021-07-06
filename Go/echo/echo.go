package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/", eat)
	e.Logger.Fatal(e.Start(":1373"))
}

type Kopol struct {
	Name string `json:"name"`
	Love string `query:"love"`
}

func eat(c echo.Context) error {
	var body Kopol

	if err := (&echo.DefaultBinder{}).BindBody(c, &body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, body)
}

// In this piece of code I want to test two things
// one: the tag code works
// two: a struct in another one is bound by echo properly
// curl -X POST -d '{"name": "parham"}' -H 'Content-Type: application/json' '127.0.0.1:1373?type=panda&love=raha'
