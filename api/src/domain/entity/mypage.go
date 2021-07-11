package entity

import (
	"encoding/json"
	"log"

	"github.com/code-wave/go-wave/infrastructure/errors"
)

type Mypage struct {
	User PublicUser `json:"user"`
	// StudyPosts StudyPosts   `json:"study_posts"`
	StudyPosts []MypagePost `json:"study_posts"`
}

type MypagePost struct {
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
}

func (m *Mypage) ResponseJSON() (interface{}, *errors.RestErr) {
	mypage, err := json.Marshal(m)
	if err != nil {
		log.Println("marshaling error " + err.Error())
		restErr := errors.NewInternalServerError("marshalling error " + err.Error())
		return nil, restErr
	}

	return mypage, nil
}
