package main

import "site/http"

func main() {
	go http.Start()

	select {}
}
