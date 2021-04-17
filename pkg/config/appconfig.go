package config

import "html/template"

//common config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
