package reverso

import (
	"fmt"
	"net/http"
	"reverso/model"
	"reverso/resolver"
	"strings"
)

type serverHandler struct {
	handler *model.Handler
}
func (s *serverHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	found := false
	for _, host := range s.handler.Hosts {
		requestHost := strings.Split(request.Host, ":")[0]
		if host.Host == requestHost {
			fmt.Printf("%s -> %s", request.Host, host.Type)
			found = true

			res, err := resolver.GetResolver(host)
			if err != nil {
				fmt.Println(err)
			}

			err = res.Resolve(writer, request)
			if err != nil {
				fmt.Println(err)
			}
			break
		}
	}
	if !found {
		fmt.Printf("???: %s -> 404\n", request.Host)
		http.NotFound(writer, request)
	}
}

func NewServerHandler(handler *model.Handler) http.Handler {
	return &serverHandler{
		handler: handler,
	}
}
