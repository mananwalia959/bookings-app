package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/mananwalia959/bookings-app/pkg/config"
)

var app *config.AppConfig
var functions = template.FuncMap{}

func AddAppConfig(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var templateCache map[string]*template.Template
	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		t, err := BuildTemplateCache()
		templateCache = t
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	parsedTemplate, isTemplateInCache := templateCache[tmpl]
	if !isTemplateInCache {
		fmt.Println("Template not in Cache")
		return
	}

	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func BuildTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		//load layout templates seperately
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
