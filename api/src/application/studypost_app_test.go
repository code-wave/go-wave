package application

import (
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/infrastructure/persistence"
	"github.com/code-wave/go-wave/utils/config"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"net/http"
	"testing"
)

var (
	services *persistence.Repositories
	err      error
	//redisService *persistence.RedisService
	sApp *studyPostApp
)

func init() {
	services, err = persistence.NewRepositories("pgx", config.Host, config.Port, config.DBUser, config.DBPassword, config.DBName)
	if err != nil {
		log.Fatal("init error: ", err.Error())
	}

	//redisService, err = persistence.NewRedisDB(config.RedisHost, config.RedisPort, config.RedisPassword)
	//if err != nil {
	//	log.Fatal("init error: ", err.Error())
	//}
	sApp = NewStudyPostApp(services.StudyPost, services.TechStack, services.StudyPostTechStack)
}

func TestSavePost(t *testing.T) {
	sPost := &entity.StudyPost{
		UserID:       1,
		Title:        "test title",
		Topic:        "test topic",
		Content:      "test content haha haha haha",
		NumOfMembers: 3,
		IsMentor:     false,
		StartDate:    "2021/6/19",
		EndDate:      "2021/6/20",
		IsOnline:     true,
		TechStack:    []string{"go", "react"},
	}

	restErr := sPost.Validate(http.MethodPost)
	if restErr != nil {
		t.Error(restErr.Error)
	}

	err := sApp.SavePost(sPost)
	if err != nil {
		t.Error(err)
	}

}

func TestSavePost_WithWrongUserID(t *testing.T) {
	sPost := &entity.StudyPost{
		UserID:       -1,
		Title:        "test title",
		Topic:        "test topic",
		Content:      "test content haha haha haha",
		NumOfMembers: 3,
		IsMentor:     false,
		StartDate:    "2021/6/19",
		EndDate:      "2021/6/20",
		IsOnline:     true,
		TechStack:    []string{"go", "react"},
	}

	restErr := sPost.Validate(http.MethodPost)
	if restErr == nil {
		t.Error("userID is negative but not filtered")
	}

	//err := sApp.SavePost(sPost)
	//if err != nil {
	//	t.Error(err)
	//}
}

func TestGetPost(t *testing.T) {
	var studyPostID int64

	studyPostID = -1 // wrong study post id
	_, err := sApp.GetPost(studyPostID)
	if err == nil { // an error will be occurred because of wrong id
		t.Error(err)
	}

	studyPostID = 1 // put at least one studypost before testing this
	_, err = sApp.GetPost(studyPostID)
	if err != nil {
		t.Error(err)
	}
}

func TestGetPostsInLatestOrder(t *testing.T) {
	var limit, offset int64
	limit, offset = 10, 0
	posts, err := sApp.GetPostsInLatestOrder(limit, offset)
	if err != nil {
		t.Error(err)
	}

	if len(posts) > int(limit) {
		t.Errorf("limit is %d but get more", limit)
	}

}
