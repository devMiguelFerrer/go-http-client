package main

import (
	"net/http"

	httpclient "github.com/devMiguelFerrer/go-http-client/httpClient"
)

func main() {
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer 123456")
	client := httpclient.New()
	resp, err := client.Get("https://api.github.com", headers)
	if err != nil {
		panic(err)
	}
	println(resp.StatusCode)

}
