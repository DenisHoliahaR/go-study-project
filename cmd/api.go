package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/DenisHoliahaR/go-beautyhub/internal/transport/http"
	"github.com/DenisHoliahaR/go-beautyhub/internal/service"
	"github.com/DenisHoliahaR/go-beautyhub/internal/infrastructure/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

  // A good base middleware stack
  r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)

  // Set a timeout value on the request context (ctx), that will signal
  // through ctx.Done() that the request has timed out and further
  // processing should be stopped.
  r.Use(middleware.Timeout(60 * time.Second))

	db := postgres.NewPostgresDB(app.config.db.dsn)

	userRepository := postgres.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService, app.logger)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)       
		r.Get("/", userHandler.GetUserList)         
		r.Get("/{id}", userHandler.GetUserById)     
		r.Put("/{id}", userHandler.UpdateUser)      
		r.Delete("/{id}", userHandler.DeleteUser)   
	})

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr: app.config.addr,
		Handler: h,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	app.logger.Info("Server has started", "address", app.config.addr)

	return srv.ListenAndServe()
}

type application struct {
	config config
	logger *slog.Logger
}

type config struct {
	addr string
	db dbConfig
}

type dbConfig struct {
	dsn string
}