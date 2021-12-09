package core

import (
	"fmt"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/faker")

	if n == nil {
		t.Fatal("nil should not be returned")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "faker" {
		t.Fatal("name should be equal to faker")
	}

	fmt.Printf("matched path: %s, params['name']: %s]n", n.pattern, ps["name"])
}