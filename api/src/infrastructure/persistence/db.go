package persistence

import (
	"database/sql"
	"fmt"
	"github.com/code-wave/go-wave/domain/repository"
	"log"
	"time"
)

const maxOpenDBConn = 10
const maxIdleDBConn = 5
const maxDBLifetime = 5 * time.Minute

type Repositories struct {
	db                 *sql.DB
	StudyPost          repository.StudyPostRepository
	User               repository.UserRepository
	TechStack          repository.TechStackRepository
	StudyPostTechStack repository.StudyPostTechStackRepository
}

func NewRepositories(driver, host, port, dbUser, password, dbName string) (*Repositories, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, dbUser, password, dbName)

	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDBConn)
	db.SetMaxIdleConns(maxIdleDBConn)
	db.SetConnMaxLifetime(maxDBLifetime)

	if err = db.Ping(); err != nil {
		return nil, err
	}
	log.Println("db connected successfully")

	return &Repositories{
		db:                 db,
		StudyPost:          NewStudyPostRepo(db),
		User:               NewUserRepository(db),
		TechStack:          NewTechStackRepo(db),
		StudyPostTechStack: NewStudyPostTechStackRepo(db),
	}, nil
}

func (s *Repositories) Close() error {
	return s.db.Close()
}
