package redirect

import (
    "net/http"
)

// FIXME:
// - route /*

// FIXME: is there a more elegant way than a "package" wide var?
var global_redirect_to string;

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, global_redirect_to + r.URL.String(), 0)

}

func DoRedirect(redirect_from, redirect_to string) {
	global_redirect_to = redirect_to
    http.HandleFunc("/", redirectHandler)
    http.ListenAndServe(redirect_from, nil)
}
