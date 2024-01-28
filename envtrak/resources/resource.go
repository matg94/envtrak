package resources

type Resource interface {
	GetId() string
    GetRow() []string
}

