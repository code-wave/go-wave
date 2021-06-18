package application

import (
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/domain/repository"
	"github.com/code-wave/go-wave/infrastructure/errors"
)

type studyPostApp struct {
	studyPostRepo          repository.StudyPostRepository // interface
	studyPostTechStackRepo repository.StudyPostTechStackRepository
}

var _ StudyPostInterface = &studyPostApp{}

type StudyPostInterface interface {
	SavePost(studyPost *entity.StudyPost) *errors.RestErr
	//GetPost(id uint64) (*entity.StudyPost, *errors.RestErr)
	//GetPostsInLatestOrder(limit, offset uint64) (entity.StudyPosts, *errors.RestErr)
}

func NewStudyPostApp(studyPostRepo repository.StudyPostRepository, studyPostTechStackRepo repository.StudyPostTechStackRepository) *studyPostApp {
	return &studyPostApp{
		studyPostRepo:          studyPostRepo,
		studyPostTechStackRepo: studyPostTechStackRepo,
	}
}

// SavePost study_post 테이블에도 저장하고 study_post_tech_stack 테이블에 (studyPostID, techStackID) 형태로도 저장
func (s *studyPostApp) SavePost(studyPost *entity.StudyPost) *errors.RestErr {
	err := s.studyPostRepo.SavePost(studyPost)
	if err != nil {
		return errors.NewInternalServerError(err.Message)
	}

	err = s.studyPostTechStackRepo.SaveStudyPostTechStack(studyPost.ID, studyPost.TechStack)
	if err != nil {
		return errors.NewInternalServerError(err.Message)
	}

	return nil
}

//func (s *studyPostApp) GetPost(id uint64) (*entity.StudyPost, *errors.RestErr) {
//	studyPost, err := s.studyPostRepo.GetPost(id)
//	s.studyPostTechStackRepo.
//}

//func (s *studyPostApp) GetPostsInLatestOrder(limit, offset uint64) (entity.StudyPosts, *errors.RestErr) {
//	return s.studyPostRepo.GetPostsInLatestOrder(limit, offset)
//}
