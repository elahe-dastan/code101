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
	Nane Nane
	Love string `query:"love"`
}

type Nane struct {
	Type string `query:"type"`
}

func eat(c echo.Context)  error {
	var body Kopol
	err := c.Bind(&body)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, body)
}


// In this piece of code I want to test two things
// one: the tag code works
// two: a struct in another one is bound by echo properly