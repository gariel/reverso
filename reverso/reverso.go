package reverso

import (
	"fmt"
	"net/http"
	"reverso/model"
	"sync"
)

type Reverso interface {
	Start() error
}

type reverso struct {
	project *model.Project
}

func (r *reverso) Start() error {
	var wg sync.WaitGroup
	for _, handler := range r.project.Handlers {
		handler := handler
		wg.Add(1)
		go func() {
			server := http.Server{
				Addr: fmt.Sprintf("0.0.0.0:%d", handler.Port),
				Handler: NewServerHandler(handler),
			}
			err := server.ListenAndServe()
			if err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}()
	}

	// TODO: better startup message and use a logger
	for _, handler := range r.project.Handlers {
		for _, host := range handler.Hosts {
			fmt.Printf("[%s] - %s:%d -> %s\n",
				host.Type,
				host.Host,
				handler.Port,
				host.Description,
			)
		}
	}
	fmt.Println("- Reverso Started -")
	wg.Wait()
	return nil
}

func NewReverso(project *model.Project) Reverso {
	return &reverso{
		project: project,
	}
}

