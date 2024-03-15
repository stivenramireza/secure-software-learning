package main

import (
	"context"
	"encoding/json"
	"github-tracker/github-tracker/models"
	"github-tracker/github-tracker/repository"
	"github-tracker/github-tracker/repository/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDummy(t *testing.T) {
	c := require.New(t)

	result := 22

	c.Equal(22, result)
}

func TestInsert(t *testing.T) {
	c := require.New(t)

	webhook := models.GitHubWebhook{
		Repository: models.Repository{
			FullName: "stivenramireza/secure-software-learning",
		},
		HeadCommit: models.Commit{
			ID:      "b1dc2566b219d3941f199b206c057e1ddeaabe43",
			Message: "Add sample code for handle-github-webhook",
			Author: models.CommitUser{
				Email:    "stivenramireza@gmail.com",
				Username: "stivenramireza",
			},
		},
	}

	body, err := json.Marshal(webhook)
	c.NoError(err)

	createdTime := time.Now()

	m := mock.Mock{}
	mockCommit := repository.MockCommit{Mock: &m}

	commit := entity.Commit{
		RepoName:       webhook.Repository.FullName,
		CommitID:       webhook.HeadCommit.ID,
		CommitMessage:  webhook.HeadCommit.Message,
		AuthorUsername: webhook.HeadCommit.Author.Username,
		AuthorEmail:    webhook.HeadCommit.Author.Email,
		Payload:        string(body),
		CreatedAt:      createdTime,
		UpdatedAt:      createdTime,
	}

	ctx := context.Background()

	mockCommit.On("Insert", ctx, &commit).Return(nil)

	err = insertGitHubWebhook(ctx, mockCommit, webhook, string(body), createdTime)
	c.NoError(err)

	m.AssertExpectations(t)
}
