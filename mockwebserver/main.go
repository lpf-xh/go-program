package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	code     = 200
	resptime time.Duration // ms
)

func main() {
	http.HandleFunc("/test", testHandle)
	http.HandleFunc("/setCode", codeHandle)
	http.HandleFunc("/setResptime", resptimeHandle)

	log.Println("mockcache is running at :80")

	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Println(err)
	}
}

func testHandle(w http.ResponseWriter, r *http.Request) {
	if resptime > 0 {
		time.Sleep(time.Millisecond * resptime)
	}
	w.WriteHeader(code)
}

func codeHandle(w http.ResponseWriter, r *http.Request) {
	v, err := strconv.Atoi(r.FormValue("code"))
	if err != nil {
		log.Println(err)
	}

	code = v
}

func resptimeHandle(w http.ResponseWriter, r *http.Request) {
	v, err := strconv.Atoi(r.FormValue("resptime"))
	if err != nil {
		log.Println(err)
	}

	resptime = time.Duration(v)
}

