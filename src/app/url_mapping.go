package app

import (
	"github.com/OleksandrVakhno/GitHubMicroservice/src/controllers/repositories"
	"github.com/OleksandrVakhno/GitHubMicroservice/src/controllers/statusReady"
)

func mapURLs(){
	router.GET("/", statusReady.Ready)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)

}