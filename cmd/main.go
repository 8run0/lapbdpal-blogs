package main

import (
	blogs "github.com/8run0/lapbdpal-blogs"
	"github.com/labstack/echo"
)

func SetupHandler() {
	var myApi BlogImpl
	e := echo.New()
	blogs.RegisterHandlers(e, &myApi)
}
