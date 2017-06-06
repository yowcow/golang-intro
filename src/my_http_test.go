package hello

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/gin-gonic/gin.v1"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestSimpleHttpRequest(t *testing.T) {
	assert := assert.New(t)

	mySimpleHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world: %s", r.URL)
	}

	r := httptest.NewRequest("GET", "http://foobar/foo/bar", nil)
	w := httptest.NewRecorder()
	mySimpleHandler(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal("Hello world: http://foobar/foo/bar", string(body))
}

func createGinApp() http.Handler {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"name":   name,
			"status": http.StatusOK,
		})
	})
	return router
}

func performRequest(r http.Handler, method, url string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, url, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestGinAppRequest(t *testing.T) {
	assert := assert.New(t)

	r := createGinApp()
	w := performRequest(r, "GET", "/user/foobar")
	result := w.Result()

	resData := struct {
		Name   string `json:"name"`
		Status int    `json:"status"`
	}{}

	body, _ := ioutil.ReadAll(result.Body)
	json.Unmarshal(body, &resData)

	re := regexp.MustCompile("^application/json;(?:\\s.+)?")

	assert.Equal(200, result.StatusCode)
	assert.Regexp(re, result.Header["Content-Type"][0])
	assert.Equal("foobar", resData.Name)
	assert.Equal(200, resData.Status)
}
