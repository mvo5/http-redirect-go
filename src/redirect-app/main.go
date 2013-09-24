package main

import (
	"fmt"
	r "redirect-lib"
)

func main() {
	redirect_from := ":8080"
	redirect_to := "http://www.uni-trier.de"

        fmt.Println("Redirecting from %s to %s", redirect_from, redirect_to)
	
	r.DoRedirect(redirect_from, redirect_to)
}
