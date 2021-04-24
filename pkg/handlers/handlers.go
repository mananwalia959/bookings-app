package handlers

import (
	"net/http"

	"github.com/mananwalia959/bookings-app/pkg/config"
	"github.com/mananwalia959/bookings-app/pkg/models"
	"github.com/mananwalia959/bookings-app/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func SetRepo(r *Repository) {
	Repo = r
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	values := make(map[string]interface{})
	values["message"] = "Hello There"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		Values: values,
	})
}
