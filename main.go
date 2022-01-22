package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type wkConfig struct {
	dpi          int
	filename     string
	footerCenter string
	footerHTML   string
	headerCenter string
	headerHTML   string
	html         string
	marginBottom int
	marginLeft   int
	marginRight  int
	marginTop    int
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/pdfs", CreatePDF)
	e.Logger.Fatal(e.Start(":1234"))
}

func CreatePDF(c echo.Context) error {
	c.Response().Header().Set("Content-type", "application/pdf")
	config := buildWKConfig(c)

	if err := generatePDF(c, config); err != nil {
		return c.JSON(http.StatusInternalServerError, "PDF could not be generated")
	}
	defer cleanPDF(config.filename)

	return c.File(fmt.Sprintf("%v.pdf", config.filename))
}

func buildWKConfig(ctx echo.Context) *wkConfig {
	config := &wkConfig{
		filename:     ctx.FormValue("filename"),
		footerCenter: ctx.FormValue("footerCenter"),
		footerHTML:   ctx.FormValue("footerHTML"),
		headerCenter: ctx.FormValue("headerCenter"),
		headerHTML:   ctx.FormValue("headerHTML"),
		html:         ctx.FormValue("html"),
	}

	if ctx.FormValue("dpi") != "" {
		config.dpi, _ = strconv.Atoi(ctx.FormValue("dpi"))
	}

	if config.filename == "" {
		config.filename = "tmp"
	}

	if ctx.FormValue("marginBottom") != "" {
		config.marginBottom, _ = strconv.Atoi(ctx.FormValue("marginBottom"))
	}

	if ctx.FormValue("marginLeft") != "" {
		config.marginLeft, _ = strconv.Atoi(ctx.FormValue("marginLeft"))
	}

	if ctx.FormValue("marginRight") != "" {
		config.marginRight, _ = strconv.Atoi(ctx.FormValue("marginRight"))
	}

	if ctx.FormValue("marginTop") != "" {
		config.marginTop, _ = strconv.Atoi(ctx.FormValue("marginTop"))
	}

	return config
}
