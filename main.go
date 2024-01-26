package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	httpClient := http.Client{}

	request, requestError := http.NewRequest("GET", "https://google.com", nil)
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
