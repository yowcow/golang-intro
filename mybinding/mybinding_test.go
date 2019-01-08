package mybinding

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/mholt/binding"
)

type LoginForm struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}

func (lf *LoginForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&lf.ID: binding.Field{
			Form:     "login_id",
			Required: true,
		},
		&lf.Password: binding.Field{
			Form:         "login_password",
			Required:     true,
			ErrorMessage: "Password is required!!!",
		},
	}
}

func TestBindingAForm(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		loginForm := new(LoginForm)
		errs := binding.Bind(req, loginForm)
		if errs.Handle(w) {
			return
		}
		enc := json.NewEncoder(w)
		err := enc.Encode(loginForm)
		if err != nil {
			fmt.Fprintln(w, "something has gone wrong:", err)
		}
	})

	svr := httptest.NewServer(handler)
	defer svr.Close()

	cases := []struct {
		Subtest  string
		Input    map[string]string
		Expected string
	}{
		{
			"Valid form",
			map[string]string{
				"login_id":       "12345",
				"login_password": "fugafuga",
			},
			`{"id":12345,"password":"fugafuga"}` + "\n",
		},
		{
			"Invalid form when missing login_id",
			map[string]string{
				"login_password": "fugafuga",
			},
			`[{"fieldNames":["login_id"],"classification":"RequiredError","message":"Required"}]`,
		},
		{
			"Invalid form when missing login_password",
			map[string]string{
				"login_id": "12345",
			},
			`[{"fieldNames":["login_password"],"classification":"RequiredError","message":"Password is required!!!"}]`,
		},
	}

	for _, c := range cases {
		t.Run(c.Subtest, func(t *testing.T) {
			query := &url.Values{}
			for k, v := range c.Input {
				query.Set(k, v)
			}

			body := strings.NewReader(query.Encode())
			req, err := http.NewRequest("POST", svr.URL, body)
			if err != nil {
				t.Error("expected nil but got", err)
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			var client http.Client
			resp, err := client.Do(req)
			if err != nil {
				t.Error("expected nil but got", err)
			}
			defer resp.Body.Close()

			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Error("expected nil but got", err)
			}
			if string(b) != c.Expected {
				t.Error("expected", c.Expected, "but got", string(b))
			}
		})
	}
}
