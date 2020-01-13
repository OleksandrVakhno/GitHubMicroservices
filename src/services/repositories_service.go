package services

import (
	"github.com/OleksandrVakhno/GitHubMicroservice/src/config"
	"github.com/OleksandrVakhno/GitHubMicroservice/src/model/github"
	"github.com/OleksandrVakhno/GitHubMicroservice/src/model/repositories"
	"github.com/OleksandrVakhno/GitHubMicroservice/src/providers/github_provider"
	"github.com/OleksandrVakhno/GitHubMicroservice/src/utils/errors"
	"net/http"
	"sync"
)

type repoService struct {}

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest)(*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(request []repositories.CreateRepoRequest)(*repositories.CreateReposResponse)
}


var (
	RepositoryService repoServiceInterface
	)

func init(){
	RepositoryService = &repoService{}
}


func (s *repoService) CreateRepo(input repositories.CreateRepoRequest)(*repositories.CreateRepoResponse, errors.ApiError){
	if err := input.Validate(); err!=nil{
		return nil, err
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(),request)

	if err!=nil{
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	return &repositories.CreateRepoResponse{
				Id:    response.Id,
				Owner: response.Name,
				Name:  response.Owner.Login,
			}, nil

}

func (s *repoService) CreateRepos(request []repositories.CreateRepoRequest)(*repositories.CreateReposResponse){
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)

	var wg sync.WaitGroup
	go s.handleRepoResults(&wg, input, output)
	for _,current :=range request{
		wg.Add(1)
		go s.createReposConcurrent(current, input)
	}

	wg.Wait()
	close(input)
	result := <-output

	success :=0
	for _, current:=range result.Results{
		if current.Response !=nil{
			success++
		}
	}

	if success==0{
		result.StatusCode = result.Results[0].Error.Status()
	}else if success==len(request){
		result.StatusCode = http.StatusCreated
	}else{
		result.StatusCode =http.StatusPartialContent
	}

	return &result
}


func (s *repoService) handleRepoResults(wg *sync.WaitGroup,input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse){
	var results repositories.CreateReposResponse
	for result := range input{
		results.Results = append(results.Results, result)
		wg.Done()
	}
	output <- results

}


func (s *repoService) createReposConcurrent(input repositories.CreateRepoRequest, output chan  repositories.CreateRepositoriesResult){
	if err:= input.Validate(); err!=nil{
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}
	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(),request)

	if err!=nil{
		output <- repositories.CreateRepositoriesResult{
			Error: errors.NewApiError(err.StatusCode, err.Message),
		}
		return
	}

	output <- repositories.CreateRepositoriesResult{Response: &repositories.CreateRepoResponse{
		Id:    response.Id,
		Owner: response.Name,
		Name:  response.Owner.Login,
	}}
}
