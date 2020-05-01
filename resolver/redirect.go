package resolver

import (
	"errors"
	"net/http"
)

type RedirectHost struct {
	Address string `json:"address"`
	Status  int    `json:"status_code"`
}

type redirectResolver struct {
	host RedirectHost
}

func (r *redirectResolver) Resolve(writer http.ResponseWriter, request *http.Request) error {
	status := r.host.Status
	if status == 0 {
		status = http.StatusFound
	}

	http.Redirect(writer, request, r.host.Address, status)
	return nil
}

func NewRedirectResolver(host RedirectHost) (Resolver, error) {
	if len(host.Address) == 0 {
		return nil, errors.New("empty redirect address")
	}
	return &redirectResolver{host}, nil
}
