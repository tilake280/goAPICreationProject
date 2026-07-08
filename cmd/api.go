package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type application struct {
	config config
	// logger (make global usually)
	//db driver
}

// mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID) //Important for rate limiting
	r.Use(middleware.RealIP)    // important for rate limiting and analytics
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // recover from crashes

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good for now"))
	})

	// http.ListenAndServe(":3333", r)

	return r
}

// run

type config struct {
	addr string // port number?
	db   dbConfig
}

type dbConfig struct {
	dsn string //domain name string user= password= etc.
}
