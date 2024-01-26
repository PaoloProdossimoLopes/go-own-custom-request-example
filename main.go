package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	httpContext := context.Background()
	httpContext, httpContextCancel := context.WithTimeout(httpContext, 3*time.Second)
	defer httpContextCancel()

	loggingHttpContext(httpContext)

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

	http.HandleFunc("/", getHandlerWithContext)
	http.ListenAndServe(":8080", nil)
}

func getHandlerWithContext(w http.ResponseWriter, r *http.Request) {
	requestContext := r.Context()
	log.Println("Request iniciada")
	defer log.Print("Request finalziada")

	select {
	case <-time.After(5 * time.Second):
		const message = "Request processada com sucesso"
		log.Println(message)
		w.Write([]byte(message))
		return

	case <-requestContext.Done():
		log.Println("Request cancelada pelo client")
		//Aqui pode para tudo que esta sendo feito pois a requisição foi cancelada
		return
	}
}
