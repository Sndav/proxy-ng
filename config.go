package main

import (
	"net/http"
	"proxy/goproxy"
)

func UpStreamConfig() []goproxy.ForwardUpstream {
	proxy1 := goproxy.ForwardUpstream{
		Address: "http://127.0.0.1:8080",
		Before: func(r *http.Request) {
			r.Header.Add("scan_id", "1")
		},
	}
	proxy2 := goproxy.ForwardUpstream{
		Address: "http://YOU_IP_ADDRESS",
		Before: func(r *http.Request) {
			r.Header.Add("scan_id", "1")
		},
	}
	return []goproxy.ForwardUpstream{proxy1, proxy2}
}
