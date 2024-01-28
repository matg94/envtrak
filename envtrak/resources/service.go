package resources

type Service struct {
    AppId string `yaml:"app-id" json:"appId"`
    Version Version `yaml:"version" json:"version"`
}

func (s *Service) GetId() string {
    return s.AppId
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

