package resources

import "strings"
import "strconv"

type Version interface {
    GetInt() int
    GetStr() string
}

type VVersion struct {
    Value string `yaml:"value" json:"value"`
}

func (v *VVersion) GetInt() int {
    version, err := strconv.Atoi(strings.Replace(v.Value, "v", "", -1))
    if err != nil {
        return -1
    }
    return version
}

func (v *VVersion) GetStr() string {
    return v.Value
}

type Versioned interface {
	GetVersion() Version
}
