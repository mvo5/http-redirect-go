package redirect

import (
	"net/http"
	"testing"
	"time"
)


func assertRedirectLocation(t *testing.T, redirect_from, redirect_to string) {

	// we can't use http.Client.Get() here as this follows redirects
	// and we want the original (no-redirected) reply
	req, err := http.NewRequest("GET", redirect_from, nil)
	if err != nil {
		t.Error("http.NewRequest failed")
	}
	tr := &http.Transport{}
	resp, err := tr.RoundTrip(req)

	if err != nil {
		t.Error("Error on http.Get (%s)", err);
	}
	if resp.Header["Location"][0] != redirect_to {
		t.Error("Unexpected Location header got: ",
			resp.Header["Location"][0],
			" want. ", redirect_to)
	}
	if resp.StatusCode != 301 {
		t.Error("Unexpected status  " + resp.Status)
	}

}

func TestRedirectRoot(t *testing.T) {
	// setup
	redirect_from := "localhost:54321"
	redirect_to := "http://example.com"
	go DoRedirect(redirect_from, redirect_to)

	// FIXME: find a more elegant way to know when GoRedirect is ready
	time.Sleep(500 * time.Millisecond)

	// test redirect for "/"
	assertRedirectLocation(t, "http://"+redirect_from, redirect_to + "/")

	// test redirect for "/some/other"
	assertRedirectLocation(t,"http://"+redirect_from+"/random/string/", redirect_to + "/random/string/")

	// teardown (how to stop it again?)
}
