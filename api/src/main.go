package main

import (
	"github.com/code-wave/go-wave/infrastructure/persistence"
	"github.com/code-wave/go-wave/utils/config"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"net/http"
)

func main() {
	services, err := persistence.NewRepositories("pgx", config.Host, config.Port, config.DBUser, config.DBPassword, config.DBName)
	if err != nil {
		log.Println(err)
		return
	}
	defer services.Close()

	//interfaces.NewStudyPost(services.StudyPost)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
