package examples

import "fmt"

type Endpoints struct {
	CurrentUserUrl   string `json:"current_user_url"`
	AuthorizationUrl string `json:"authorization_url"`
	RepositoryUrl    string `json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error) {
	response, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		// Deal with err
		return nil, err
	}

	fmt.Printf(fmt.Sprintf("Status Code: %d", response.StatusCode))
	fmt.Printf(fmt.Sprintf("Status: %s", response.Status))
	fmt.Printf(fmt.Sprintf("Body: %s/b", response.String()))

	var endpoints Endpoints
	if err := response.UnmarshalJson(&endpoints); err != nil {
		// Deal with err
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Repository URL: %s", endpoints.AuthorizationUrl))
	return &endpoints, nil
}
