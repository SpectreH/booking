package render

import (
	"booking/pkg/config"
	"booking/pkg/models"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(templateData *models.TemplateData) *models.TemplateData {
	return templateData
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {
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

	templateData = AddDefaultData(templateData)

	_ = template.Execute(buf, templateData)

	_, err := buf.WriteTo(w)
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
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*.html")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.html")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts
	}

	return cache, nil
}
