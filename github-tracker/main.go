package main

import (
	"context"
	"fmt"
	"github-tracker/github-tracker/models"
	repositoy "github-tracker/github-tracker/repository"
	"github-tracker/github-tracker/repository/entity"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received POST request!")

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading the request")

		return
	}

	fmt.Println(string(body))
}

func insertGitHubWebhook(ctx context.Context, repo repositoy.Commit, webhook models.GitHubWebhook, body string, createdTime time.Time) error {
	commit := entity.Commit{
		RepoName:       webhook.Repository.FullName,
		CommitID:       webhook.HeadCommit.ID,
		CommitMessage:  webhook.HeadCommit.Message,
		AuthorUsername: webhook.HeadCommit.Author.Username,
		AuthorEmail:    webhook.HeadCommit.Author.Email,
		Payload:        body,
		CreatedAt:      createdTime,
		UpdatedAt:      createdTime,
	}

	fmt.Println(commit)

	return nil
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", postHandler).Methods("POST")

	fmt.Println("Server listening on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err.Error())
	}
}
