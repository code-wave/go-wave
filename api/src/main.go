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
	"github.com/rs/cors"
)

func main() {
	services, err := persistence.NewRepositories("pgx", config.Host, config.Port, config.DBUser, config.DBPassword, config.DBName)
	if err != nil {
		log.Println(err)
		return
	}
	defer services.Close()

	redisService, err := persistence.NewRedisDB(config.RedisHost, config.RedisPort, config.RedisPassword)
	if err != nil {
		log.Println(err)
		return
	}

	userApp := application.NewUserApp(services.User)
	userHandler := interfaces.NewUserHandler(userApp)

	authApp := application.NewAuthApp(redisService.Auth)
	authHandler := interfaces.NewAuthHandler(userApp, authApp)
	//interfaces.NewStudyPost(services.StudyPost)

	r := chi.NewRouter()
	// r.Use(middleware.CORSMiddleware)
	//users
	r.Get("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("hello")) })
	r.With(middleware.AuthVerifyMiddleware).Get("/users/{user_id}", userHandler.GetUser)
	r.Get("/users/limit={limit}&offset={offset}", userHandler.GetAllUsers)
	r.Post("/users/signup", userHandler.SaveUser)
	r.With(middleware.AuthVerifyMiddleware).Patch("/users/{user_id}", userHandler.UpdateUser)
	r.With(middleware.AuthVerifyMiddleware).Delete("/users/{user_id}", userHandler.DeleteUser)
  
  //auth
	r.Post("/auth/users/login", authHandler.LoginUser)
	r.With(middleware.AuthVerifyMiddleware).Post("/auth/users/logout", authHandler.LogoutUser)
	r.With(middleware.AuthVerifyMiddleware).Post("/auth/users/refresh", authHandler.Refresh)
  
  
	studyPostApp := application.NewStudyPostApp(services.StudyPost, services.TechStack, services.StudyPostTechStack)
	studyPostHandler := interfaces.NewStudyPostHandler(studyPostApp)

	r.Get("/study-post/{study_post_id}", studyPostHandler.GetPost)
	r.Get("/study-posts/limit={limit}&offset={offset}", studyPostHandler.GetPostsInLatestOrder)
	r.Get("/study-posts/user_id={user_id}&limit={limit}&offset={offset}", studyPostHandler.GetPostsByUserID)
	r.Post("/study-post", studyPostHandler.SavePost)
	r.Patch("/study-post", studyPostHandler.UpdatePost)
	r.Delete("/study-post/{study_post_id}", studyPostHandler.DeletePost)

	techStackApp := application.NewTechStackApp(services.TechStack)
	techStackHandler := interfaces.NewTechStackHandler(techStackApp)

	r.Get("/tech-stack/{tech_stack_id}", techStackHandler.GetTechStack)
	r.Get("/tech-stacks", techStackHandler.GetAllTechStack)
	r.Get("/tech-stacks/{study_post_id}", techStackHandler.GetAllTechStackByStudyPostID)
	r.Post("/tech-stack", techStackHandler.SaveTechStack)
	r.Delete("/tech-stack/tech-name={tech_name}", techStackHandler.DeleteTechStack)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	handler := cors.Default().Handler(r)
	handler = c.Handler(handler)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
