package resolver

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reverso/model"
)

type Resolver interface {
	Resolve(writer http.ResponseWriter, request *http.Request) error
}

func GetResolver(host model.Host) (Resolver, error) {
	switch host.Type {
	case "proxy":
		var proxyHost ProxyHost
		err := json.Unmarshal(host.Data, &proxyHost)
		if err != nil {
			return nil, err
		}
		return NewProxyResolver(proxyHost)

	case "redirect":
		var redirectHost RedirectHost
		err := json.Unmarshal(host.Data, &redirectHost)
		if err != nil {
			return nil, err
		}
		return NewRedirectResolver(redirectHost)

	case "fixed":
		var fixedHost FixedHost
		err := json.Unmarshal(host.Data, &fixedHost)
		return NewFixedResolver(fixedHost), err

	case "static":
		var staticHost StaticHost
		err := json.Unmarshal(host.Data, &staticHost)
		return NewStaticResolver(staticHost), err
	}
	return nil, errors.New(fmt.Sprintf("no resolver found for %s", host.Type))
}
