package main

import (
	"log"
	"net/http"

	"github.com/code-wave/go-wave/application"
	"github.com/code-wave/go-wave/infrastructure/persistence"
	"github.com/code-wave/go-wave/interfaces"
	"github.com/code-wave/go-wave/interfaces/middleware"
	"github.com/code-wave/go-wave/utils/config"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	services, err := persistence.NewRepositories("pgx", config.Host, config.Port, config.DBUser, config.DBPassword, config.DBName)
	if err != nil {
		log.Println(err)
		return
	}
	defer services.Close()

	//interfaces.NewStudyPost(services.StudyPost)

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("hello")) })
	userApp := application.NewUserApp(services.User)
	userHandler := interfaces.NewUserHandler(userApp)
	r.Post("/users/signup", userHandler.SaveUser)
	r.Get("/users/{user_id}", userHandler.GetUser)
	r.Get("/users/limit={limit}&offset={offset}", userHandler.GetAllUsers)
	r.Patch("/users/{user_id}", userHandler.UpdateUser)
	r.Delete("/users/{user_id}", userHandler.DeleteUser)

	authService, err := persistence.NewRedisDB(config.RedisHost, config.RedisPort, config.RedisPassword)
	if err != nil {
		log.Println(err)
		return
	}

	authApp := application.NewAuthApp(authService.Auth)
	authHandler := interfaces.NewAuthHandler(userApp, authApp)
	r.Post("/auth/users/login", authHandler.LoginUser)

	ar := chi.NewRouter()
	ar.Use(middleware.AuthVerifyMiddleware)
	ar.Post("/auth/users/logout", authHandler.LogoutUser)
	ar.Post("/auth/users/refresh", authHandler.Refresh)
	r.Mount("/", ar)

	log.Fatal(http.ListenAndServe(":8080", r))
}
