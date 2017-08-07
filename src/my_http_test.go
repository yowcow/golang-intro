package hello

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"sync"
	"testing"

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

func startRealHTTPServer(addr string, wg *sync.WaitGroup, ch <-chan struct{}) {
	srvmux := &http.ServeMux{}
	srvmux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi, I love %s!", r.URL.Path[1:])
	})
	server := &http.Server{
		Addr:    addr,
		Handler: srvmux,
	}

	defer func() {
		fmt.Println("Server is closing")
		server.Close()
		wg.Done()
	}()

	go func() {
		server.ListenAndServe()
	}()

	fmt.Println("Server is listening")

	for _ = range ch {
	}
}

func TestRealHTTPRequest(t *testing.T) {
	wg := &sync.WaitGroup{}
	ch := make(chan struct{})

	wg.Add(1)
	go startRealHTTPServer(":8899", wg, ch)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8899/foo/bar", nil)
	res, _ := client.Do(req)

	close(ch)
	wg.Wait()

	body, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, "Hi, I love foo/bar!", string(body))
}
