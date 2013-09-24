package redirect

import (
	"log"
	"net/http"
	"strings"
)

// FIXME: is there a more elegant way than a "package" wide var?
var global_redirect_to string;

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, global_redirect_to + r.URL.String(), 
		      http.StatusMovedPermanently)
}

func DoRedirect(redirect_from, redirect_to string) {
	// the root URL that we redirect to can not have a trailing "/"
	// as we add it in each redirect and that would lead to "//"
	global_redirect_to = strings.Trim(redirect_to, "/")

        http.HandleFunc("/", redirectHandler)
        err := http.ListenAndServe(redirect_from, nil)
	if (err != nil) {
		log.Fatal("ListenAndServer: ", err)
	}
}
