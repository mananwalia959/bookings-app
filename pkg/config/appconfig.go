package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

//common config
type AppConfig struct {
	UseCache       bool
	TemplateCache  map[string]*template.Template
	UseTls         bool
	SessionManager *scs.SessionManager
}
