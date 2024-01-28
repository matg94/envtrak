package resources

import (
    "gopkg.in/yaml.v3"
    "encoding/json"
)

type Service struct {
    AppId string `yaml:"app-id" json:"appId"`
    Version Version `yaml:"version" json:"version"`
}

func (s *Service) GetId() string {
    return s.AppId
}

func (s *Service) GetYaml() string {
    res, err := yaml.Marshal(&s)
    if err != nil {
        return ""
    }
    return string(res)
}

func (s *Service) GetJson() string{
    res, err := json.Marshal(&s)
    if err != nil {
        return ""
    }
    return string(res)
} 
func (s *Service) GetRow() []string {
    return []string{
        s.AppId,
        s.Version.GetStr(),
    }
}

func (s *Service) GetVersion() string {
    return s.Version.GetStr()
}

