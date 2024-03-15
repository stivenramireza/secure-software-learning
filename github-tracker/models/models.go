package models

type GitHubWebhook struct {
	Repository Repository `json:"repository"`
	HeadCommit Commit     `json:"head_commit"`
}

type Repository struct {
	FullName string `json:"full_name"`
}

type Commit struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Author  struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
}

type CommitUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
