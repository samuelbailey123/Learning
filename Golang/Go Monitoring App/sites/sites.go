package sites

import (
	"fmt"
	"net/http"
	"playground/webFunctions"
)

func SourceCheck(w http.ResponseWriter, r *http.Request) {
	googleCheck(w, r)
	fmt.Fprintln(w, "")
	microsoftCheck(w, r)
}

func googleCheck(w http.ResponseWriter, r *http.Request) {
	url := "http://www.google.com"
	isOnline := webFunctions.PingSite(url)
	fmt.Fprintf(w, "The site %s is %s", url, isOnline)
}

func microsoftCheck(w http.ResponseWriter, r *http.Request) {
	url := "http://www.microsoft.com"
	isOnline := webFunctions.PingSite(url)
	fmt.Fprintf(w, "The site %s is %s", url, isOnline)
}
