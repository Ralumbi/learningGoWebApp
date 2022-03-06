package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/ralumbi/learningGoWebApp/pkg/config"
	"github.com/ralumbi/learningGoWebApp/pkg/handlers"
	"github.com/ralumbi/learningGoWebApp/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {
	// Change true for production
	app.InProduction = false

	session = scs.New()
	// Set the lifetime of your session
	session.Lifetime = 24 * time.Hour
	// Set to false if you want to close the session once the browser is closed
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	// Set to true to force https
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create Cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	log.Fatal(err)
}
