package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	wkpdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/labstack/echo/v4"
)

func generatePDF(ctx echo.Context, config *wkConfig) error {
	pdfg, err := wkpdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	pdfg.SetOutput(ctx.Response().Writer)
	pdfg.AddPage(wkpdf.NewPageReader(bytes.NewReader([]byte(config.html))))
	applyAttributes(pdfg, config)

	if err = pdfg.Create(); err != nil {
		return err
	}

	if err = pdfg.WriteFile(fmt.Sprintf("%v.pdf", config.filename)); err != nil {
		return err
	}

	return nil
}

func applyAttributes(pdfg *wkpdf.PDFGenerator, config *wkConfig) {
	if config.dpi != 0 {
		pdfg.Dpi.Set(uint(config.dpi))
	}

	if config.footerCenter != "" {
		pdfg.TOC.FooterCenter.Set(config.footerCenter)
	}

	if config.footerHTML != "" {
		pdfg.TOC.FooterHTML.Set(config.footerHTML)
	}

	if config.headerCenter != "" {
		pdfg.TOC.HeaderCenter.Set(config.headerCenter)
	}

	if config.headerHTML != "" {
		pdfg.TOC.HeaderHTML.Set(config.headerHTML)
	}

	if config.marginBottom != 0 {
		pdfg.MarginBottom.Set(uint(config.marginBottom))
	}

	if config.marginLeft != 0 {
		pdfg.MarginLeft.Set(uint(config.marginLeft))
	}

	if config.marginRight != 0 {
		pdfg.MarginRight.Set(uint(config.marginRight))
	}

	if config.marginTop != 0 {
		pdfg.MarginTop.Set(uint(config.marginTop))
	}
}

func cleanPDF(filename string) {
	e := os.Remove(fmt.Sprintf("%v.pdf", filename))
	if e != nil {
		log.Fatal(e)
	}
}
