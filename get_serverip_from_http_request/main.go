package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptrace"
)

// 方法一：通过httptrace.ClientTrace获取服务IP地址
func m1() {
	trace := &httptrace.ClientTrace{
		DNSStart: func(_ httptrace.DNSStartInfo) {},
		DNSDone:  func(_ httptrace.DNSDoneInfo) {},
		ConnectStart: func(net, addr string) {
			fmt.Printf("ConnectStart addr=%s\n", addr)
		},
		ConnectDone: func(net, addr string, err error) {
			fmt.Printf("ConnectDone addr=%s\n", addr)
		},
		GotConn:              func(_ httptrace.GotConnInfo) {},
		GotFirstResponseByte: func() {},
		TLSHandshakeStart:    func() {},
		TLSHandshakeDone:     func(_ tls.ConnectionState, _ error) {},
	}

	req, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	req = req.WithContext(httptrace.WithClientTrace(context.Background(), trace))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

// 方法二：通过DialContext获取服务IP地址
func m2() {
	req, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				conn, err := net.Dial(network, addr)
				req.RemoteAddr = conn.RemoteAddr().String()
				return conn, err
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("RemoteAddr:", req.RemoteAddr)
}

func main() {
	m1()
	m2()
}

