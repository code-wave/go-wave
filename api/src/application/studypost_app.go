package application

import (
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/domain/repository"
	"github.com/code-wave/go-wave/infrastructure/errors"
)

type studyPostApp struct {
	studyPostRepo          repository.StudyPostRepository // interface
	techStackRepo          repository.TechStackRepository
	studyPostTechStackRepo repository.StudyPostTechStackRepository
}

var _ StudyPostInterface = &studyPostApp{}

type StudyPostInterface interface {
	SavePost(studyPost *entity.StudyPost) *errors.RestErr
	GetUserIDByPostID(studyPostID int64) (int64, *errors.RestErr)
	GetPost(id int64) (*entity.StudyPost, *errors.RestErr)
	GetPostsInLatestOrder(limit, offset int64) (entity.StudyPosts, *errors.RestErr)
	GetPostsByUserID(userID, limit, offset int64) (entity.StudyPosts, *errors.RestErr)
	UpdatePost(studyPost *entity.StudyPost) (*entity.StudyPost, *errors.RestErr)
	DeletePost(studyPostID int64) *errors.RestErr
}

func NewStudyPostApp(studyPostRepo repository.StudyPostRepository, techStackRepo repository.TechStackRepository, studyPostTechStackRepo repository.StudyPostTechStackRepository) *studyPostApp {
	return &studyPostApp{
		studyPostRepo:          studyPostRepo,
		techStackRepo:          techStackRepo,
		studyPostTechStackRepo: studyPostTechStackRepo,
	}
}

// SavePost study_post 테이블에도 저장하고 study_post_tech_stack 테이블에 (studyPostID, techStackID) 형태로도 저장
func (s *studyPostApp) SavePost(studyPost *entity.StudyPost) *errors.RestErr {
	err := s.techStackRepo.CheckTechStack(studyPost.TechStack)
	if err != nil {
		return err
	}

	studyPost, err = s.studyPostRepo.SavePost(studyPost)
	if err != nil {
		return err
	}

	err = s.studyPostTechStackRepo.SaveStudyPostTechStack(studyPost.ID, studyPost.TechStack)
	if err != nil {
		return err
	}

	return nil
}

func (s *studyPostApp) GetUserIDByPostID(studyPostID int64) (int64, *errors.RestErr) {
	studyPost, err := s.studyPostRepo.GetPost(studyPostID)
	if err != nil {
		return 0, err
	}

	return studyPost.UserID, nil
}
func (s *studyPostApp) GetPost(studyPostID int64) (*entity.StudyPost, *errors.RestErr) {
	return s.studyPostRepo.GetPost(studyPostID)
}

func (s *studyPostApp) GetPostsInLatestOrder(limit, offset int64) (entity.StudyPosts, *errors.RestErr) {
	return s.studyPostRepo.GetPostsInLatestOrder(limit, offset)
}

func (s *studyPostApp) GetPostsByUserID(userID, limit, offset int64) (entity.StudyPosts, *errors.RestErr) {
	return s.studyPostRepo.GetPostsByUserID(userID, limit, offset)
}

func (s *studyPostApp) UpdatePost(studyPost *entity.StudyPost) (*entity.StudyPost, *errors.RestErr) {
	updatedPost, err := s.studyPostRepo.UpdatePost(studyPost)
	if err != nil {
		return nil, err
	}

	err = s.studyPostTechStackRepo.UpdateStudyPostTechStack(studyPost.ID, studyPost.TechStack)
	if err != nil {
		return nil, err
	}

	return updatedPost, nil
}

func (s *studyPostApp) DeletePost(studyPostID int64) *errors.RestErr {
	return s.studyPostRepo.DeletePost(studyPostID)
}
