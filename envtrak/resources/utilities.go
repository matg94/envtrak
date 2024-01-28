package resources

import (
    "gopkg.in/yaml.v3"
    "encoding/json"
)

func ToYaml(resource Resource) string {
    res, err := yaml.Marshal(&resource)
    if err != nil {
        return ""
    }
   return string(res)

}

func ToJson(resource Resource) string {
    res, err := json.Marshal(&resource)
    if err != nil {
        return ""
    }
   return string(res)
}
