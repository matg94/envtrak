package resources

type Deployment struct {
    Resource Versioned `yaml:"resource" json:"resource"`
    Environment Environment `yaml:"environment" json:"environment"`
}

