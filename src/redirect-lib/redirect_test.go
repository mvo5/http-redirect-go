package redirect

import (
	"net/http"
	"testing"
)

func assertRedirectLocation(t *testing.T, resp* http.Response, err error, redirect_to string) {
	if err != nil {
		t.Error("Error on http.Get (%s)", err);
	}
	if resp.Header["Location"][0] != redirect_to {
		t.Error("Unexpected Location header ", 
			resp.Header["Location"][0] + " != " + redirect_to)
	}
}

func TestRedirectRoot(t *testing.T) {
	// setup
	redirect_from := "localhost:54321"
	redirect_to := "http://example.com"
	go DoRedirect(redirect_from, redirect_to)

	// test redirect for "/"
	resp, err := http.Get("http://"+redirect_from) 
	assertRedirectLocation(t, resp, err, redirect_to)

	// test redirect for "/some/other"
	resp, err = http.Get("http://"+redirect_from+"/random/string") 
	assertRedirectLocation(t, resp, err, redirect_to)

	// teardown (how to stop it again?)
}
