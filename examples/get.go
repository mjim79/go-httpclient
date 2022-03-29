package examples

import "fmt"

type Endpoints struct {
	CurrentUserUrl    string `json:"current_user_url"`
	AuthorizationsUrl string `json:"authorizations_url"`
	RepositoryUrl     string `json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error) {
	response, err := httpClient.Get("https://api.github.com")
	if err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Status Code: %d", response.StatusCode))
	fmt.Println(fmt.Sprintf("Status: %s", response.Status))
	fmt.Println(fmt.Sprintf("String: %s\n", response.String()))

	var endpoints Endpoints
	if err := response.UnmarshalJson(&endpoints); err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Repository URL: %s", endpoints.RepositoryUrl))
	return &endpoints, nil
}
