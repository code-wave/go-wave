package main

import (
	"log"
	"net/http"

	"github.com/atg0831/booking_study/infrastructure/persistence"
	"github.com/atg0831/booking_study/utils/config"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	services, err := persistence.NewRepositories("pgx", config.Host, config.Port, config.DBUser, config.DBPassword, config.DBName)
	if err != nil {
		log.Println(err)
		return
	}
	defer services.Close()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
