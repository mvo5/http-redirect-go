package main

import (
	"fmt"
	"flag"
	r "redirect-lib"
)

func main() {
	redirect_from := flag.String("from", ":80", 
		"http redirect from interface:port")
	redirect_to := flag.String("to", 
		"http://www.uni-trier.de/",
		"redirect to base address")
	flag.Parse()

        fmt.Println("Redirecting from " + *redirect_from + 
		" to " + *redirect_to)
	
	r.DoRedirect(*redirect_from, *redirect_to)
}
