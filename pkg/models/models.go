package models

type TemplateData struct {
	Values    map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
