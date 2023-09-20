package controllers

import (
	"net/http"
	"strconv"
	"math/big"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type FibonacciController struct{}

func NewFibonacciController() *FibonacciController {
	return new(FibonacciController)
}

func FibonacciIterative(n int) *big.Int {
	x, y := big.NewInt(0), big.NewInt(1)
	for i := 0; i < n; i++ {
		x, y = y, new(big.Int).Add(x, y)
	}
	return x
}

func (fc *FibonacciController) GetFibonacci(c echo.Context) error {
	nStr := c.QueryParam("n")
	n, err := strconv.Atoi(nStr)

	if err != nil || n < 0 || n > 100000 {
		errorResponse := ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Bad request",
		}
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	result := FibonacciIterative(n)
	return c.JSON(http.StatusOK, map[string]interface{}{"result": result})
}
