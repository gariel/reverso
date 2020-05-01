package resolver

import (
	"net/http"
	"path/filepath"
)

type StaticHost struct {
	Directory string `json:"directory"`
	AllowListing bool `json:"allow_listing"`
}

type staticResolver struct {
	server http.Handler
}

func (s *staticResolver) Resolve(writer http.ResponseWriter, request *http.Request) error {
	s.server.ServeHTTP(writer, request)
	return nil
}

func NewStaticResolver(host StaticHost) Resolver {
	var dir http.FileSystem = http.Dir(host.Directory)

	if !host.AllowListing {
		dir = neuteredFileSystem{dir}
	}

	server := http.FileServer(dir)
	return &staticResolver{server}
}


// https://www.alexedwards.net/blog/disable-http-fileserver-directory-listings
type neuteredFileSystem struct {
    fs http.FileSystem
}

func (n neuteredFileSystem) Open(path string) (http.File, error) {
    f, err := n.fs.Open(path)
    if err != nil {
        return nil, err
    }

    s, _ := f.Stat()
    if s.IsDir() {
        index := filepath.Join(path, "index.html")
        if _, err := n.fs.Open(index); err != nil {
            closeErr := f.Close()
            if closeErr != nil {
                return nil, closeErr
            }

            return nil, err
        }
    }

    return f, nil
}
