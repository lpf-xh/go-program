package main

import (
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/miekg/dns"
)

var (
	rcode    int
	resptime time.Duration
)

func main() {
	go func() {
		http.HandleFunc("/setRcode", rcodeHandle)
		http.HandleFunc("/setResptime", resptimeHandle)
		log.Println("mockdns http is running at :8053")
		log.Panic(http.ListenAndServe(":10053", nil))
	}()

	go func() {
		dns.HandleFunc(".", dnsReply)
		log.Println("mockdns dns is running at :53")
		log.Panic(dns.ListenAndServe(":53", "udp", nil))
	}()

	select {}
}

func rcodeHandle(w http.ResponseWriter, r *http.Request) {
	v, err := strconv.Atoi(r.FormValue("code"))
	if err != nil {
		log.Println(err)
	}

	rcode = v
}

func resptimeHandle(w http.ResponseWriter, r *http.Request) {
	v, err := strconv.Atoi(r.FormValue("resptime"))
	if err != nil {
		log.Println(err)
	}

	resptime = time.Duration(v)
}

func dnsReply(w dns.ResponseWriter, req *dns.Msg) {
	if resptime > 0 {
		time.Sleep(resptime * time.Millisecond)
	}

	a := &dns.A{}
	a.Hdr.Name = req.Question[0].Name
	a.Hdr.Rrtype = dns.TypeA
	a.Hdr.Class = dns.ClassINET
	a.Hdr.Ttl = 300
	a.Hdr.Rdlength = 4
	a.A = net.ParseIP("2.0.1.9")

	res := *req
	res.Rcode = rcode
	res.Answer = append(res.Answer, a)

	w.WriteMsg(&res)
}

