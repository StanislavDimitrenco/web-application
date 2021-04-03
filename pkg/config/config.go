package config

import (
	"html/template"
	"log"
)

const PORT = ":8080"

//AppConfig holds application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
