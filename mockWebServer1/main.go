package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	respCode = 200
	respTime time.Duration // ms
)

func main() {
	http.HandleFunc("/test", testHandle)
	http.HandleFunc("/setRespCode", codeHandle)
	http.HandleFunc("/setRespTime", resptimeHandle)

	log.Println("mockcache is running at :80")

	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Println(err)
	}
}

func testHandle(w http.ResponseWriter, r *http.Request) {
	if respTime > 0 {
		time.Sleep(time.Millisecond * respTime)
	}
	w.WriteHeader(respCode)
}

func codeHandle(w http.ResponseWriter, r *http.Request) {
	v, err := strconv.Atoi(r.FormValue("code"))
	if err != nil {
		log.Println(err)
	}

	respCode = v
}

func resptimeHandle(w http.ResponseWriter, r *http.Request) {
	v, err := strconv.Atoi(r.FormValue("time"))
	if err != nil {
		log.Println(err)
	}

	respTime = time.Duration(v)
}

