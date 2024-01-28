package resources

import (
    "gopkg.in/yaml.v3"
    "encoding/json"
    "fmt"
)

type Environment struct {
    Dtap       string `yaml:"dtap" json:"dtap"` 
	BusinessId string `yaml:"business-id" json:"businessId"`
    Alias      string `yaml:"alias" json:"alias"`
	Region     string `yaml:"region" json:"region"`
}


func (e *Environment) GetId() string {
    return fmt.Sprintf("%s-%s-%s", e.Alias, e.Dtap, e.BusinessId) 
}

func (e *Environment) GetYaml() string {
    res, err := yaml.Marshal(&e)
    if err != nil {
        return ""
    }
    return string(res)
}

func (e *Environment) GetJson() string{
    res, err := json.Marshal(&e)
    if err != nil {
        return ""
    }
    return string(res)
} 
func (e *Environment) GetRow() []string {
    return []string{
        e.BusinessId,
        e.Alias,
        e.Dtap,
        e.Region,
    }
}
