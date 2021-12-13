package render

import (
	"booking/internal/config"
	"booking/internal/models"
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{}

var app *config.AppConfig
var pathToTemplates = "./templates"

// SetNewRenderer sets the config for the template package
func SetNewRenderer(a *config.AppConfig) {
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
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, templateData *models.TemplateData) error {
	var templateCache map[string]*template.Template

	if app.UseCache {
		// Get the template cache from the app config
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	template, ok := templateCache[tmpl]
	if !ok {
		return errors.New("can't get template from cache")
	}

	buf := new(bytes.Buffer)

	templateData = AddDefaultData(templateData, r)

	err := template.Execute(buf, templateData)
	if err != nil {
		fmt.Println("error executing template", err)
		return err
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
		return err
	}

	return nil
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(pathToTemplates + "/*.html")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		withOutExtension := strings.Split(filepath.Base(page), ".")

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			fmt.Println(err)
			return cache, err
		}

		matches, err := filepath.Glob(pathToTemplates + "/layouts/*.html")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(pathToTemplates + "/layouts/*.html")
			if err != nil {
				return cache, err
			}
		}
		cache[withOutExtension[0]] = ts
	}

	return cache, nil
}
