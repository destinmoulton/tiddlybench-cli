package util

import (
	"fmt"
	"net/http"
	"net/url"
)

// IsURL tests a string for for Scheme and Host
func IsURL(str string) bool {
	// https://stackoverflow.com/questions/31480710/validate-url-with-standard-package-in-go
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func TestURL(url string) bool {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error received in http resp")
		fmt.Println(err)
		return false
	}
	fmt.Println(resp)
	return true
}
