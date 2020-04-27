package resolver

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"reverso/model"
	"strings"
)

type proxyResolver struct {}

func (p *proxyResolver) Resolve(host model.Host, writer http.ResponseWriter, request *http.Request) error {
	director := func(req *http.Request) {
		target, _ := url.Parse(host.Address)

		targetQuery := target.RawQuery
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)

		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}

		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Host = target.Host
		if _, ok := req.Header["User-Agent"]; !ok {
			req.Header.Set("User-Agent", "")
		}
	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(writer, request)
	return nil
}

func NewProxyResolver() Resolver {
	return &proxyResolver {}
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}