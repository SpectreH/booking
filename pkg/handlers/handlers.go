package handlers

import (
	"booking/pkg/config"
	"booking/pkg/models"
	"booking/pkg/render"
	"net/http"
)

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

// Repo is the repository used by the handlers
var Repo *Repository

// SetNewHandlers sets the repository for the handlers
func SetNewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	// send data to the template
	render.RenderTemplate(w, "about", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals", &models.TemplateData{})
}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors", &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability", &models.TemplateData{})
}
