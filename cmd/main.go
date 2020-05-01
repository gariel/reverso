package main

import (
	"io/ioutil"
	"reverso/model"
	"reverso/reverso"
)

func main() {
	//project := model.Project{
	//	Handlers: []model.Handler{
	//		{Port: 8080, Hosts: []model.Host{
	//			{Host: "local.com", Type: "proxy", Address: "https://www.google.com"},
	//			{Host: "api.local.com", Type: "redirect", Address: "https://www.bing.com"},
	//		}},
	//		{Port: 9090, Hosts: []model.Host{
	//			{Host: "local.com", Type: "proxy", Address: "https://br.yahoo.com/"},
	//			{Host: "api.local.com", Type: "redirect", Address: "https://www.twitter.com"},
	//		}},
	//	},
	//}

	file, _ := ioutil.ReadFile("config.json")
	project, err := model.NewProjectFromContent(file)

	if err != nil {
		panic(err)
	}

	r := reverso.NewReverso(project)
	err = r.Start()
	if err != nil {
		panic(err)
	}
}
