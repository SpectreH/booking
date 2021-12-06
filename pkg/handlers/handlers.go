package handlers

import (
	"booking/pkg/config"
	"booking/pkg/models"
	"booking/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type (Repository pattern)
type Repository struct {
	App *config.AppConfig
}

// SetNewRepo creates a new repository
func SetNewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// SetNewHandlers sets the repository for the handlers
func SetNewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	// send data to the template
	render.RenderTemplate(w, "about.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
