package main

import (
	"context"
	"fmt"
	"time"
)

func loggingHttpContext(context context.Context) {
	//context `cancel` method trigger the `Done` method
	select {
	case <-context.Done():
		fmt.Println("Http Request was cancelled. Timeout reached")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Http request complete with success")
		return
	default:
		fmt.Println("Something wrong occurs ...")
		return
	}
}
