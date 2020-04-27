package resolver

import (
	"net/http"
	"reverso/model"
)

type redirectResolver struct {

}

func (r *redirectResolver) Resolve(host model.Host, writer http.ResponseWriter, request *http.Request) error {
	// TODO status 301 or 302?
	http.Redirect(writer, request, host.Address, http.StatusTemporaryRedirect)
	return nil
}

func NewRedirectResolver() Resolver {
	return &redirectResolver {}
}
