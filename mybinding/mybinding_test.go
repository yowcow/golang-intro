package mybinding

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/mholt/binding"
)

type LoginForm struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

func (lf *LoginForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&lf.ID: binding.Field{
			Form:     "login_id",
			Required: true,
		},
		&lf.Password: binding.Field{
			Form:     "login_password",
			Required: true,
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

	var client http.Client
	resp, err := client.Get(svr.URL + "?login_id=hoge&login_password=fuga")
	if err != nil {
		t.Error("expected nil but got", err)
	}
	defer resp.Body.Close()

	actual := new(LoginForm)
	expected := &LoginForm{"hoge", "fuga"}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(actual)
	if err != nil {
		t.Error("expected nil but got", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %#v but got %#v", expected, actual)
	}
}
