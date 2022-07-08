package webFunctions

import (
	"fmt"
	"net/http"
)

func ping(url string) (bool, error) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
		return false, err
	}
	return true, nil
}

func PingSite(url string) string {
	ok, err := ping(url)
	if err != nil {
		fmt.Println(err)
	}
	if ok {
		return "Online"
	}
	return "Offline"
}
