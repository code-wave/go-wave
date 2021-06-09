package entity

import (
	"github.com/code-wave/go-wave/infrastructure/helpers"
)

type StudyPosts []StudyPost

type StudyPost struct {
	ID           uint64
	Title        string // 게시글 제목
	Topic        string // 프로젝트 주제
	Content      string
	TechStackID  uint64
	NumOfMembers int64
	IsMento      bool
	Price        int64  // 1인당 필요한 가격
	StartDate    string // 프로젝트 시작
	EndDate      string // 프로젝트 끝
	UserID       int64
	IsOnline     bool
	CreatedAt    string
	UpdatedAt    string
}

func (s *StudyPost) Validate() map[string]string {
	var errorMessages = make(map[string]string)
	var err error

	err = helpers.CheckStringMinChar(s.Title, 5)
	if err != nil {
		errorMessages["invalid_title"] = err.Error()
	}

	err = helpers.CheckStringMinChar(s.Topic, 5)
	if err != nil {
		errorMessages["invalid_topic"] = err.Error()
	}

	err = helpers.CheckStringMinChar(s.Content, 20)
	if err != nil {
		errorMessages["invalid_content"] = err.Error()
	}

	if s.NumOfMembers <= 0 {
		errorMessages["invalid_number_of_members"] = "number of members can't be 0 or negative"
	}

	if s.Price < 0 {
		errorMessages["invalid_price"] = "price can't be negative"
	}

	if s.StartDate == "" {
		errorMessages["invalid_start_date"] = "empty start date" // TODO: 이거랑 enddate 같이 오류 더 디테일하게
	}

	if s.EndDate == "" {
		errorMessages["invalid_end_date"] = "empty end date"
	}

	return nil
}
