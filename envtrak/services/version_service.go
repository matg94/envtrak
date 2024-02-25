package services

import (
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/matg94/envtrak/config"
	"github.com/matg94/envtrak/domain"
)

type VersionService interface {
	GetServiceInstances(config.Service) []domain.Instance
}

type EnvironmentListResponse struct {
	Id uint64 `json:"id"`
}

type EnvironmentInfoResponse struct {
	LastDeployment struct {
		CreatedAt string `json:"created_at"`
		Commit    struct {
			Title      string `json:"title"`
			Id         string `json:"id"`
			AuthorName string `json:"author_name"`
		} `json:"commit"`
		Pipeline struct {
			WebUrl string `json:"web_url"`
		} `json:"pipeline"`
	}
}

type GitlabVersionService struct {
	Client *resty.Client
}

func (gl GitlabVersionService) GetServiceInstances(service config.Service) []domain.Instance {
	var instances []domain.Instance
	for _, env := range service.Environments {
		instances = append(instances, gl.getInstanceFromReleases(service.GitlabProjectID, env))
	}
	return instances
}

func (gl GitlabVersionService) getInstanceFromReleases(projectId int64, env config.Environment) domain.Instance {
	environmentList := []EnvironmentListResponse{}
	resp, err := gl.
		Client.
		R().
		SetResult(&environmentList).
		EnableTrace().
		SetQueryParam("name", env.GitlabName).
		Get(fmt.Sprintf("/api/v4/projects/%d/environments", projectId))

	if err != nil {
		return domain.CreateFailedInstance(err, env)
	}

	if resp.StatusCode() != 200 || len(environmentList) != 1 {
        return domain.CreateFailedInstance(
            errors.New(
                fmt.Sprintf(
                    "Invalid HTTP Response status %d with body %s",
                    resp.StatusCode(),
                    resp.Body(),
                ),
            ),
            env,
        )
	}

	environmentDetails := &EnvironmentInfoResponse{}
	
    resp, err = gl.Client.
		R().
		SetResult(environmentDetails).
		EnableTrace().
		Get(fmt.Sprintf("api/v4/projects/%d/environments/%d", projectId, environmentList[0].Id))

	if err != nil {
		return domain.CreateFailedInstance(err, env)
	}

	if resp.StatusCode() != 200 || environmentDetails.LastDeployment.Commit.Title == "" {  
		return domain.CreateFailedInstance(
			errors.New(
				fmt.Sprintf(
					"Invalid HTTP Response status %d with body %s",
					resp.StatusCode(),
					resp.Body(),
				),
			),
			env,
		)
	}

    fmt.Println(environmentDetails)

	return domain.Instance{
		Version: domain.Version{
			Value:       environmentDetails.LastDeployment.Commit.Id,
			Title:       environmentDetails.LastDeployment.Commit.Title,
			Description: environmentDetails.LastDeployment.Commit.AuthorName,
		},
		CreatedAt:   environmentDetails.LastDeployment.CreatedAt,
		InfoUrl:     environmentDetails.LastDeployment.Pipeline.WebUrl,
		Environment: env,
	}
}
