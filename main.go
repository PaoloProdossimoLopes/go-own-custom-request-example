package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	httpContext := context.Background()
	httpContext, httpContextCancel := context.WithTimeout(httpContext, 3*time.Second)
	defer httpContextCancel()

	httpClient := http.Client{}

	request, requestError := http.NewRequestWithContext(httpContext, "GET", "https://google.com", nil)
	if requestError != nil {
		panic(requestError)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, responseError := httpClient.Do(request)
	if responseError != nil {
		panic(requestError)
	}

	defer response.Body.Close()

	responseBody, responseBodyError := ioutil.ReadAll(response.Body)
	if responseBodyError != nil {
		panic(responseBodyError)
	}

	println(responseBody)
}
