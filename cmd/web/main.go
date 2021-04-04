package main

import (
	"fmt"
	"github.com/StanislavDimitrenko/web-application/pkg/config"
	"github.com/StanislavDimitrenko/web-application/pkg/handler"
	"github.com/StanislavDimitrenko/web-application/pkg/render"
	"log"
	"net/http"
)

func main() {
	var app config.AppConfig

	//get the template cache from the app config
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can't get templates")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handler.NewRepo(&app)
	handler.NewHandler(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handler.Repo.Home)
	http.HandleFunc("/about", handler.Repo.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", config.PORT))
	_ = http.ListenAndServe(config.PORT, nil)

}
