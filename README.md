# Simple PDF Generation API with wkhtmltopdf

## Quick start
Clone the repo locally and run:
```bash
go run .
```

## Supported Form-Data Params
```bash
dpi
filename
footerCenter
footerHTML
headerCenter
headerHTML
html
marginBottom
marginLeft
marginRight
marginTop
```

## Potential Improvements
- [x] Stream file into response 
- [ ] Parse down form-data directly into wickedPDF cli
- [ ] Handle requests concurrently