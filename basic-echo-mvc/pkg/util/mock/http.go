package mock

import (
	"io"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

type HttpMock struct {
	E *echo.Echo
}

func (em *HttpMock) NewRequest(method, path string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, path, body)
	rec := httptest.NewRecorder()
	c := em.E.NewContext(req, rec)

	return c, rec
}
