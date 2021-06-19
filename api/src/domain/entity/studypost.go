package entity

import (
	"encoding/json"
	"github.com/code-wave/go-wave/infrastructure/errors"
	"github.com/code-wave/go-wave/infrastructure/helpers"
)

type StudyPosts []StudyPost

type StudyPost struct {
	ID           int64    `json:"id"`
	Title        string   `json:"title"` // 게시글 제목
	Topic        string   `json:"topic"` // 프로젝트 주제
	Content      string   `json:"content"`
	NumOfMembers int64    `json:"num_of_members"`
	IsMentor     bool     `json:"is_mentor"`
	Price        int64    `json:"price"`      // 1인당 필요한 가격
	StartDate    string   `json:"start_date"` // 프로젝트 시작
	EndDate      string   `json:"end_date"`   // 프로젝트 끝
	UserID       int64    `json:"user_id"`
	IsOnline     bool     `json:"is_online"`
	TechStack    []string `json:"tech_stack"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
}

func (s *StudyPost) Validate() *errors.RestErr {
	err := helpers.CheckStringMinChar(s.Title, 5)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	err = helpers.CheckStringMinChar(s.Topic, 5)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	err = helpers.CheckStringMinChar(s.Content, 20)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	if s.NumOfMembers <= 0 {
		return errors.NewBadRequestError("number of members can't be 0 or negative")
	}

	if s.Price < 0 {
		return errors.NewBadRequestError("price can't be negative")
	}

	if s.StartDate == "" {
		return errors.NewBadRequestError("empty start date") // TODO: 이거랑 enddate 같이 오류 더 디테일하게
	}

	if s.EndDate == "" {
		return errors.NewBadRequestError("empty end date")
	}

	err = helpers.ConvertStringArray(s.TechStack) // 여기서 리스트를 조작함
	if err != nil {
		errors.NewBadRequestError(err.Error())
	}

	return nil
}

func (s *StudyPost) ResponseJSON() ([]byte, *errors.RestErr) {
	sJson, err := json.Marshal(s)
	if err != nil {
		return nil, errors.NewInternalServerError("marshal error")
	}

	return sJson, nil
}
