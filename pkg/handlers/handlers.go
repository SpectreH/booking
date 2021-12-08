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
	render.RenderTemplate(w, r, "home", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	// send data to the template
	render.RenderTemplate(w, r, "about", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals", &models.TemplateData{})
}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors", &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability", &models.TemplateData{})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := (r.Form.Get("start"))
	end := (r.Form.Get("end"))

	w.Write([]byte(start + end))
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact", &models.TemplateData{})
}

// Reservation renders the reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation", &models.TemplateData{})
}
