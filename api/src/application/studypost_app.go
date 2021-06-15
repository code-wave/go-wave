package application

import (
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/domain/repository"
	"github.com/code-wave/go-wave/infrastructure/errors"
)

type studyPostApp struct {
	studyPostRepo repository.StudyPostRepository // interface
}

var _ StudyPostInterface = &studyPostApp{}

type StudyPostInterface interface {
	SavePost(post *entity.StudyPost) (*entity.StudyPost, *errors.RestErr)
	GetPost(id uint64) (*entity.StudyPost, *errors.RestErr)
	GetPostsInLatestOrder(limit, offset uint64) (entity.StudyPosts, *errors.RestErr)
}

func (s *studyPostApp) SavePost(post *entity.StudyPost) (*entity.StudyPost, *errors.RestErr) {
	return s.studyPostRepo.SavePost(post)
}

func (s *studyPostApp) GetPost(id uint64) (*entity.StudyPost, *errors.RestErr) {
	return s.studyPostRepo.GetPost(id)
}

func (s *studyPostApp) GetPostsInLatestOrder(limit, offset uint64) (entity.StudyPosts, *errors.RestErr) {
	return s.studyPostRepo.GetPostsInLatestOrder(limit, offset)
}
