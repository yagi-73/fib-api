package router

import (
	"fibonacci-api/controllers"
	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	fc := controllers.NewFibonacciController()
	e.GET("/fib", fc.GetFibonacci)
	return e
}