package resolver

import "net/http"

type FixedHost struct {
	Status  int  `json:"status_code"`
	Content string `json:"content"`
}

type fixedResolver struct {
	host FixedHost
}

func (f *fixedResolver) Resolve(writer http.ResponseWriter, _ *http.Request) error {
	status := f.host.Status
	if status == 0 {
		status = http.StatusOK
	}

	writer.WriteHeader(status)
	_, err := writer.Write([]byte(f.host.Content))
	return err
}

func NewFixedResolver(host FixedHost) Resolver {
	return &fixedResolver{host}
}
