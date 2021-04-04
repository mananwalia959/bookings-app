package config

import "html/template"

//common config
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
