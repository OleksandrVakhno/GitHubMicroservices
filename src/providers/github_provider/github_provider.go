package github_provider

import (
	"encoding/json"
	"fmt"
	"github.com/OleksandrVakhno/GitHubMicroservice/src/clients/restclient"
	"github.com/OleksandrVakhno/GitHubMicroservice/src/model/github"
	"io/ioutil"
	"log"
	"net/http"
)

const(
	headerAuthorization = "Authorization"
	headerAuthorizationFormat = "token %s"

	urlCreateRepo = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string{
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GitHubErrorResponse){
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	response, err := restclient.Post(urlCreateRepo, request, headers )
	
	if err!= nil{
		log.Printf("Error when trying to create new repo in github: %s", err.Error())
		return nil, &github.GitHubErrorResponse{
			StatusCode:       http.StatusInternalServerError,
			Message:          err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if err!=nil{
		return nil, &github.GitHubErrorResponse{
			StatusCode:       http.StatusInternalServerError,
			Message:          "Invalid response body from Github",
		}
	}

	if response.StatusCode > 299{
		var errResponse github.GitHubErrorResponse
		if err:=json.Unmarshal(bytes, &errResponse); err!=nil{
			return nil, &github.GitHubErrorResponse{
				StatusCode:       http.StatusInternalServerError,
				Message:          "Invalid JSON response body",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	if err:=json.Unmarshal(bytes, &result); err!=nil{
		return nil, &github.GitHubErrorResponse{
			StatusCode:       http.StatusInternalServerError,
			Message:          "Invalid JSON response body",

		}
	}

	return &result, nil
}
