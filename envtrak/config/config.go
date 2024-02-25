package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Environment struct {
	Business   string `yaml:"business"`
	Dtap       string `yaml:"dtap"`
	Alias      string `yaml:"alias"`
	GitlabName string `yaml:"gitlab-name"`
}

type Service struct {
	AppID           string        `yaml:"app-id"`
	GitlabProjectID int64         `yaml:"gitlab-project-id"`
	Description     string        `yaml:"description"`
	Environments    []Environment `yaml:"environments"`
}

type TrackedServices struct {
	Services []Service `yaml:"services"`
}

func LoadConfig(configPath string) TrackedServices {
    data, err := ioutil.ReadFile(configPath)
    if err != nil {
        panic(err)
    }

    var trackedServices TrackedServices

    err = yaml.Unmarshal(data, &trackedServices)
    if err != nil {
        panic(err)
    }

    return trackedServices
}
