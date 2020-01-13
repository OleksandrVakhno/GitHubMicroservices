package repositories

import (
	"github.com/OleksandrVakhno/GitHubMicroservice/src/utils/errors"
	"strings"
)

type CreateRepoRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

type CreateRepoResponse struct {
	Id int64 `json:"id"`
	Owner string `json:"owner"`
	Name string `json:"name"`
}

type CreateReposResponse struct {
	StatusCode int `json:"status_code"`
	Results []CreateRepositoriesResult `json:"results"`
}

type CreateRepositoriesResult struct {
	Response *CreateRepoResponse `json:"response"`
	Error errors.ApiError	`json:"error"`
}

func (r *CreateRepoRequest) Validate() errors.ApiError{
	r.Name = strings.TrimSpace(r.Name)
	if r.Name == "" {
		return  errors.NewBadRequestError("Invalid repository name")
	}
	return nil
}