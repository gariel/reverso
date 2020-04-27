package resolver

import (
	"net/http"
	"reverso/model"
)

type Resolver interface {
	Resolve(host model.Host, writer http.ResponseWriter, request *http.Request) error
}

func GetResolver(typename string) Resolver {
	switch typename {
	case "proxy":
		return NewProxyResolver()
	case "redirect":
		return NewRedirectResolver()
	}
	return nil
}