package myecho

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func getHelloAction(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}

type HelloPerson struct {
	Name string `json:"name" form:"name"`
}

func postHelloFormAction(c echo.Context) error {
	hello := &HelloPerson{}

	if err := c.Bind(hello); err != nil {
		return err
	}

	return c.String(http.StatusOK, "Hello "+hello.Name)
}

func postHelloJSONAction(c echo.Context) error {
	hello := &HelloPerson{}

	if err := c.Bind(hello); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"hello": hello.Name,
	})
}

func headerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			return err
		}
		c.Response().Header().Set("X-Server", "Echo")
		return nil
	}
}

func TestSimpleGET(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	e := echo.New()
	e.GET("/", getHelloAction)
	e.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Hello world", w.Body.String())
	assert.Equal(t, "text/plain; charset=UTF-8", w.Header().Get("Content-Type"))
}

func TestPOSTFormBody(t *testing.T) {
	query := url.Values{}
	query.Add("name", "hoge")
	body := bytes.NewBufferString(query.Encode())

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/", body)
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	e := echo.New()
	e.POST("/", postHelloFormAction)
	e.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Hello hoge", w.Body.String())
	assert.Equal(t, "text/plain; charset=UTF-8", w.Header().Get("content-type"))
}

func TestPOSTJSONBody(t *testing.T) {
	body := bytes.NewBufferString(`{"name":"fuga"}`)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/", body)
	req.Header.Set("content-type", "application/json")

	e := echo.New()
	e.POST("/", postHelloJSONAction)
	e.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"hello":"fuga"}`, w.Body.String())
	assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("content-type"))
}

func TestSimpleMiddleware(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	e := echo.New()
	e.Use(headerMiddleware)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world?")
	})
	e.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Hello world?", w.Body.String())
	assert.Equal(t, "Echo", w.Header().Get("X-Server"))
}

func TestElasticMiddleware(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("something-important", "hogefuga")
			return next(c)
		}
	})
	e.GET("/", func(c echo.Context) error {
		something := c.Get("something-important").(string)
		return c.String(http.StatusOK, something)
	})
	e.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "hogefuga", w.Body.String())
}
