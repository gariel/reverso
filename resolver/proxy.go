package resolver

import (
	"errors"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type ProxyHost struct {
	Address string `json:"address"`
}

type proxyResolver struct {
	proxy http.Handler
}

func (p *proxyResolver) Resolve(writer http.ResponseWriter, request *http.Request) error {
	p.proxy.ServeHTTP(writer, request)
	return nil
}

func NewProxyResolver(host ProxyHost) (Resolver, error) {
	if len(host.Address) == 0 {
		return nil, errors.New("empty proxy address")
	}

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
	return &proxyResolver {proxy}, nil
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
