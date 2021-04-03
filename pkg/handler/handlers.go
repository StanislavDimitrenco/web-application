package handler

import "github.com/StanislavDimitrenko/web-application/pkg/config"

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
