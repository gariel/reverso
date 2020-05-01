package model

import "encoding/json"

type Host struct {
	Description string `json:"description"`
	Host        string `json:"host"`
	Type        string `json:"type"`
	Data        []byte
}

func (h *Host) Specialize(to interface{}) {

}

type Handler struct {
	Port      int32 `json:"port"`
	Hosts     []Host
	HostsData []map[string]interface{} `json:"hosts"`
}

type Project struct {
	Description string    `json:"description"`
	Handlers    []*Handler `json:"handlers"`
}

func (p *Project) load() {
	for _, handler := range p.Handlers {
		handler.Hosts = []Host{}
		for _, m := range handler.HostsData {
			data, _ := json.Marshal(m)
			var host Host
			_ = json.Unmarshal(data, &host)
			host.Data = data
			handler.Hosts = append(handler.Hosts, host)
		}
	}
}

func NewProjectFromContent(data []byte) (*Project, error) {
	project := Project{}
	err := json.Unmarshal(data, &project)
	if err != nil {
		return nil, err
	}
	project.load()
	return &project, nil
}
