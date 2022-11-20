package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/engEhabIbrahim/go-booking/pkg/config"
	"github.com/engEhabIbrahim/go-booking/pkg/handlers"
	"github.com/engEhabIbrahim/go-booking/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this to true if you're in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true // you can you different storage like badger db , sqlite , postgresql and so on ( default is cookies)
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Start listening on port : %s\n", portNumber)

	//_ = http.ListenAndServe(portNumber, nil)

	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
