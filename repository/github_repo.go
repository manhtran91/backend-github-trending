package repository

import (
	"backend-github-trending/model"
	"context"
)

type GithubRepo interface {
	SaveRepo(context context.Context, user model.GithubRepo) (model.GithubRepo, error)
	SelectRepos(context context.Context, limit int) ([]model.GithubRepo, error)
	SelectRepoByName(context context.Context, name string) (model.GithubRepo, error)
	UpdateRepo(context context.Context, user model.GithubRepo) (model.GithubRepo, error)
}