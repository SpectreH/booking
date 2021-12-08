package render

import (
	"booking/internal/config"
	"booking/internal/models"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// SetNewTemplates sets the config for the template package
func SetNewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(templateData *models.TemplateData, r *http.Request) *models.TemplateData {
	templateData.Flash = app.Session.PopString(r.Context(), "flash")
	templateData.Error = app.Session.PopString(r.Context(), "error")
	templateData.Warning = app.Session.PopString(r.Context(), "warning")
	templateData.CSRFToken = nosurf.Token(r)
	return templateData
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, templateData *models.TemplateData) {
	var templateCache map[string]*template.Template

	if app.UseCache {
		// Get the template cache from the app config
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	template, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	templateData = AddDefaultData(templateData, r)

	err := template.Execute(buf, templateData)
	if err != nil {
		fmt.Println("error executing template", err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		withOutExtension := strings.Split(filepath.Base(page), ".")

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./templates/layouts/*.html")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/layouts/*.html")
			if err != nil {
				return cache, err
			}
		}
		cache[withOutExtension[0]] = ts
	}

	return cache, nil
}
