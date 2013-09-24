package redirect

import (
    "net/http"
)

// FIXME:
// - tests
// - hardcoded uni-trier
// - route /*


func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://www.uni-trier.de", 0)
}

func DoRedirect(redirect_to string) {
    http.HandleFunc("/", redirectHandler)
    http.ListenAndServe(":8080", nil)
}
