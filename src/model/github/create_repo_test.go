package github

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "GoLangMicroservices",
		Description: "Test Repo",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	bytes,err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	fmt.Println(string(bytes))
	var target CreateRepoRequest

	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)

	assert.EqualValues(t, target, request)



}