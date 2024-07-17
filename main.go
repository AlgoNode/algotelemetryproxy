package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type OKResponse struct {
	Result string `json:"result"`
}

func main() {

	e := echo.New()
	e.IPExtractor = echo.ExtractIPFromXFFHeader(echo.TrustPrivateNet(true))
	e.Use(middleware.Decompress())
	e.POST("/*", func(c echo.Context) error {
		b, err := io.ReadAll(c.Request().Body)
		if err == nil {
			fmt.Println(string(b))
		}
		return c.JSON(http.StatusOK, &OKResponse{Result: "OK"})
	})
	e.HEAD("/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})
	e.Logger.Fatal(e.Start(":8881"))
}
