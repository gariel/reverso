package model

type Host struct {
	Description string `json:"description"`
	Host        string `json:"host"`
	Type        string `json:"type"`
	Address     string `json:"address"`
	//Middleware  string `json:""`
}

type Handler struct {
	Port  int32  `json:"port"`
	Hosts []Host `json:"hosts"`
}

type Project struct {
	Description string    `json:"description"`
	Handlers    []Handler `json:"handlers"`
	//Plugins     []string `json:""`
}
