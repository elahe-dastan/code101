package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/eat", eat)
	e.POST("/sleep", sleep)
	e.GET("/surprise", surprise)
	e.Logger.Fatal(e.Start(":1373"))
}

type KopolComposite struct {
	Name string `json:"name"`
	//Nane Nane
	Love string `query:"love"`
	Birthday  *int `query:"birthday"`
}

type KopolEmbed struct {
	Name string `json:"name"`
	Love string `query:"love"`
	//Nane
}

type JustQuery struct {
	Something string `query:"something"`
	Nane Nane
}

type Nane struct {
	Type string `query:"type"`
}

func eat(c echo.Context)  error {
	var body KopolComposite

	err := c.Bind(&body)
	if err != nil {
		return err
	}

	//if err := (&echo.DefaultBinder{}).BindBody(c, &body); err != nil {
	//	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	//}
	//
	//if err := (&echo.DefaultBinder{}).BindQueryParams(c, &body); err != nil {
	//	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	//}

	return c.JSON(http.StatusOK, body)
}

func sleep(c echo.Context)  error {
	var body KopolEmbed

	err := c.Bind(&body)
	if err != nil {
		return err
	}

	//if err := (&echo.DefaultBinder{}).BindBody(c, &body); err != nil {
	//	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	//}
	//
	//if err := (&echo.DefaultBinder{}).BindQueryParams(c, &body); err != nil {
	//	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	//}

	return c.JSON(http.StatusOK, body)
}

func surprise(c echo.Context) error {
	var body JustQuery

	err := c.Bind(&body)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, body)
}

// In this piece of code I want to test two things
// one: the tag query
// two: a struct in another one is bound by echo properly
// curl -X POST -d '{"name": "parham"}' -H 'Content-Type: application/json' '127.0.0.1:1373?type=panda&love=raha'

// for echo to work you should use only query tag or json tag and if you want to use both you should write BindBody and
// BindQueryParams separately and the struct in another one also works