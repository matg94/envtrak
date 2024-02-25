package main

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/matg94/envtrak/config"
	"github.com/matg94/envtrak/services"
)

var trackedServices config.TrackedServices
var gitlabVersionService services.GitlabVersionService

func main() {
    
    trackedServices = config.LoadConfig("envtrak.yml")
    httpClient := resty.
        New(). 
        SetBaseURL("http://gitlab.com").
        SetHeader("PRIVATE-TOKEN", os.Getenv("GITLAB_API_TOKEN"))

    gitlabVersionService = services.GitlabVersionService{
        Client: httpClient,
    }

    fmt.Println(trackedServices)
    serviceZeroInstances := gitlabVersionService.GetServiceInstances(trackedServices.Services[0])
    fmt.Println(serviceZeroInstances)
}
