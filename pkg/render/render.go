package render

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/mananwalia959/bookings-app/pkg/config"
	"github.com/mananwalia959/bookings-app/pkg/models"
)

var app *config.AppConfig
var functions = template.FuncMap{}

func AddAppConfig(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		t, err := BuildTemplateCache()
		tc = t
		if err != nil {
			fmt.Println("Can't build Template Cache", err)
		}
	}

	parsedTemplate, isTemplateInCache := tc[tmpl]
	if !isTemplateInCache {
		fmt.Println("Template not in Cache")
		return
	}

	buf := new(bytes.Buffer)

	err := parsedTemplate.ExecuteTemplate(buf, "base", td)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = buf.WriteTo(w)
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
