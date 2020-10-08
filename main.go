package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type route struct {
	Id      string
	Balance string   `yaml:"balance,omitempty"`
	Option  []string `yaml:"option,omitempty"`
	Cluster []struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"cluster"`
}

type config struct {
	Global struct {
		Log []struct {
			Address  string `yaml:"address"`
			Facility string `yaml:"facility"`
			Level    string `yaml:"level,omitempty"`
		} `yaml:"log,omitempty"`
		Maxconn int    `yaml:"maxconn,omitempty"`
		User    string `yaml:"user,omitempty"`
		Group   string `yaml:"group,omitempty"`
		Daemon  bool   `yaml:"daemon"`
	} `yaml:"global,omitempty"`
	Defaults struct {
		Option []string `yaml:"option,omitempty"`
		Stats  struct {
			Auth struct {
				User     string `yaml:"user"`
				Password string `yaml:"password"`
			} `yaml:"auth"`
			URI string `yaml:"uri"`
		} `yaml:"stats,omitempty"`
	} `yaml:"defaults,omitempty"`
	Backends map[string]route `yaml:"backends"`
}

func main() {
	haproxyConfig := config{}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal([]byte(data), &haproxyConfig)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	i := 0
	for k, v := range haproxyConfig.Backends {
		i++
		v.Id = fmt.Sprintf("backend_%v", i)
		haproxyConfig.Backends[k] = v
	}
	tmpl, err := template.ParseFiles("haproxy.cfg")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	tmpl.Execute(os.Stdout, haproxyConfig)
}
