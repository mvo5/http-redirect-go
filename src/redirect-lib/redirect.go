package redirect

import (
	"log"
	"net/http"
	"strings"
)

// type that implements http.Handler
type redirectHandler struct {
	base_redirect_to string
}
func (rh *redirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, 
		rh.base_redirect_to + r.URL.String(), 
		http.StatusMovedPermanently)
}

// FIXME: provide a way to stop a redirected port
func DoRedirect(redirect_from, redirect_to string) {
	// the root URL that we redirect to can not have a trailing "/"
	// as we add it in each redirect and that would lead to "//"
	redirect_to = strings.Trim(redirect_to, "/")

	// build a custom handler to allow supporting different redirects
	// on different interface:port combinations
	rh := redirectHandler{base_redirect_to: redirect_to}
        err := http.ListenAndServe(redirect_from, &rh)
	if (err != nil) {
		log.Fatal("ListenAndServer: ", err)
	}
}
