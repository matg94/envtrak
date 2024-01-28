package resources

type Resource interface {
	GetId() string
	GetYaml() string
	GetJson() string
    GetAsRow() string
}

