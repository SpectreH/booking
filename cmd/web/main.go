package main

import (
	"booking/internal/config"
	driver "booking/internal/driver"
	"booking/internal/handlers"
	"booking/internal/helpers"
	"booking/internal/models"
	"booking/internal/render"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8000"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errLog *log.Logger

const (
	// Initialize connection constants.
	HOST     = "mydemoserver.postgres.database.azure.com"
	DATABASE = "mypgsqldb"
	USER     = "mylogin@mydemoserver"
	PASSWORD = "<server_admin_password>"
)

// main is the main function
func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Println("Staring application on port", portNumber)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {
	var err error

	// what am I going to put in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	// change this to true when in production
	app.UseCache = false
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errLog = log.New(os.Stdout, "ERRROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	log.Println("Connecting to database...")
	// var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", HOST, USER, PASSWORD, DATABASE)
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=booking user=spectre password=admin sslmode=require")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}
	log.Println("Connected to database!")

	app.TemplateCache, err = render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}

	repo := handlers.SetNewRepo(&app, db)
	handlers.SetNewHandlers(repo)

	render.SetNewRenderer(&app)
	helpers.SetNewHelpers(&app)

	return db, nil
}
