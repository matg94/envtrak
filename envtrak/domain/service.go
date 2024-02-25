package domain

import (
	"github.com/matg94/envtrak/config"
)

type Version struct {
	Value       string
	Title       string
	Description string
	InfoUrl     string
}

type Instance struct {
	Version     Version
	CreatedAt   string
	InfoUrl     string
	Environment config.Environment
}

type Service struct {
	AppId       string
	Description string
	Instances   []Instance
}

func CreateService(serviceInfo config.Service, serviceInstances []Instance) Service {
	return Service{
		AppId:       serviceInfo.AppID,
		Description: serviceInfo.Description,
		Instances:   serviceInstances,
	}
}

func CreateFailedInstance(err error, env config.Environment) Instance {
    return Instance{
		Version: Version{
			Title:       "failed to get version",
			Description: err.Error(),
		},
		Environment: env,
	}
}
