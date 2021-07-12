package entity

import (
	"encoding/json"
	"log"

	"github.com/code-wave/go-wave/infrastructure/errors"
)

type Mypage struct {
	User                   PublicUser   `json:"user"`
	WritedStudyPosts       []MypagePost `json:"writed_study_posts"`
	ParticipatedStudyPosts []MypagePost `json:"participated_study_posts"`
}

type MypagePost struct {
	Writer    string `json:"wirter"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
}

func (m *Mypage) ResponseJSON() ([]byte, *errors.RestErr) {
	mypage, err := json.Marshal(m)
	if err != nil {
		log.Println("marshaling error " + err.Error())
		restErr := errors.NewInternalServerError("marshalling error " + err.Error())
		return nil, restErr
	}

	return mypage, nil
}
