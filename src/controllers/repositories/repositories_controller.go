package repositories

import (
	"github.com/OleksandrVakhno/GitHubMicroservice/src/model/repositories"
	"github.com/OleksandrVakhno/GitHubMicroservice/src/services"
	"github.com/OleksandrVakhno/GitHubMicroservice/src/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRepo(c *gin.Context){
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err!=nil{
		apiErr := errors.NewBadRequestError( "Invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result ,err := services.RepositoryService.CreateRepo(request)

	if err!=nil{
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func CreateRepos(c *gin.Context){
	var request []repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err!=nil{
		apiErr := errors.NewBadRequestError( "Invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result := services.RepositoryService.CreateRepos(request)

	c.JSON(result.StatusCode, result)
}