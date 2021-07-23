package server

import (
	"bytes"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)


func setupTest(HTTPMethod string, path string, body []byte) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()

	req := httptest.NewRequest(HTTPMethod, path, bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	return e.NewContext(req, rec), rec
}
