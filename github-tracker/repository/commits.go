package repository

import (
	"context"
	"database/sql"
	"github-tracker/github-tracker/repository/entity"
	"log"
)

type Commit interface {
	Insert(ctx context.Context, commit *entity.Commit) (err error)
	GetCommitByAuthorEmail(ctx context.Context, email string) (commits []entity.Commit, err error)
}

type commit struct {
	Conn *sql.DB
}

func NewCommit(conn *sql.DB) commit {
	return commit{
		Conn: conn,
	}
}

func (m commit) Insert(ctx context.Context, commit *entity.Commit) (err error) {
	query := `
		INSERT INTO commits (repo_name, commit_id, commit_message, author_username, author_email, payload, created_at, updated_at)
		VALUIES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	stmr, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmr.Close()

	err = stmr.QueryRowContext(
		ctx,
		commit.RepoName,
		commit.CommitID,
		commit.CommitMessage,
		commit.AuthorUsername,
		commit.AuthorEmail,
		commit.Payload,
		commit.CreatedAt,
		commit.UpdatedAt,
	).Err()

	return err
}

func (m commit) GetCommitByAuthorEmail(ctx context.Context, email string) (commits []entity.Commit, err error) {
	query := `
	  SELECT * FROM commits
	  WHERE author_email = $1
	`

	rows, err := m.Conn.QueryContext(ctx, query, email)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Error closing rows:", err)
		}
	}()

	for rows.Next() {
		var commit entity.Commit
		err := rows.Scan(
			&commit.ID,
			&commit.RepoName,
			&commit.CommitID,
			&commit.CommitMessage,
			&commit.AuthorUsername,
			&commit.AuthorEmail,
			&commit.Payload,
			&commit.CreatedAt,
			&commit.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		commits = append(commits, commit)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return commits, nil
}
