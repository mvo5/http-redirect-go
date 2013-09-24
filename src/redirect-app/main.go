package main

import (
	r "redirect-lib"
	"fmt"
)

func main() {
        fmt.Println("Redirecting")	
	r.DoRedirect("http://www.uni-trier.de")
}
