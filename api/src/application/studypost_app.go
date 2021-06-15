package application

import (
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/domain/repository"
)

type studyPostApp struct {
	studyPostRepo repository.StudyPostRepository // interface
}

var _ StudyPostInterface = &studyPostApp{}

type StudyPostInterface interface {
	SavePost(post *entity.StudyPost) (*entity.StudyPost, error)
	GetPost(id uint64) (*entity.StudyPost, error)
	GetPostsInLatestOrder(limit, offset uint64) (entity.StudyPosts, error)
}

func (s *studyPostApp) SavePost(post *entity.StudyPost) (*entity.StudyPost, error) {
	return s.studyPostRepo.SavePost(post)
}

func (s *studyPostApp) GetPost(id uint64) (*entity.StudyPost, error) {
	return s.studyPostRepo.GetPost(id)
}

func (s *studyPostApp) GetPostsInLatestOrder(limit, offset uint64) (entity.StudyPosts, error) {
	return s.studyPostRepo.GetPostsInLatestOrder(limit, offset)
}
