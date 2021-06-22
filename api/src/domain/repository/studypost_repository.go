package repository

import (
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/infrastructure/errors"
)

type StudyPostRepository interface {
	SavePost(studyPost *entity.StudyPost) (*entity.StudyPost, *errors.RestErr)
	GetPost(id int64) (*entity.StudyPost, *errors.RestErr)
	GetPostsInLatestOrder(limit, offset int64) (entity.StudyPosts, *errors.RestErr)
	GetPostsByUserID(userID, limit, offset int64) (entity.StudyPosts, *errors.RestErr)
	UpdatePost(studyPost *entity.StudyPost) (*entity.StudyPost, *errors.RestErr)
	DeletePost(studyPostID int64) *errors.RestErr
}
