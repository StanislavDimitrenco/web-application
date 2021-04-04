package handler

import (
	"github.com/StanislavDimitrenko/web-application/pkg/config"
	"github.com/StanislavDimitrenko/web-application/pkg/models"
	"github.com/StanislavDimitrenko/web-application/pkg/render"
	"net/http"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo the repository used by the handlers
var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringData := make(map[string]string)
	stringData["test"] = "Hello World"

	render.Template(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringData,
	})
}
