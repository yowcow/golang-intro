package hello

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSimpleHTTPRequest(t *testing.T) {
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

	decoder := json.NewDecoder(w.Body)
	decoder.Decode(&resData)

	re := regexp.MustCompile("^application/json;(?:\\s.+)?")

	assert.Equal(200, result.StatusCode)
	assert.Regexp(re, result.Header["Content-Type"][0])
	assert.Equal("foobar", resData.Name)
	assert.Equal(200, resData.Status)
}

func TestRealHTTPRequest(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hi, I love %s!", req.URL.Path[1:])
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", server.URL+"/foo/bar", nil)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "Hi, I love foo/bar!", string(body))
}

func TestTimeoutRequest(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hoge", func(w http.ResponseWriter, req *http.Request) {
		time.Sleep(200 * time.Millisecond)
		w.Header().Set("content-type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success":true}`))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	client := http.Client{
		Timeout: 100 * time.Millisecond,
	}
	req, _ := http.NewRequest("GET", server.URL+"/hoge", nil)
	resp, err := client.Do(req)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
}
