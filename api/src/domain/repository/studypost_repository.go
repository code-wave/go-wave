package repository

import "github.com/code-wave/go-wave/domain/entity"

type StudyPostRepository interface {
	SavePost(post *entity.StudyPost) (*entity.StudyPost, error)
	GetPost(id uint64) (*entity.StudyPost, error)
	GetPostsInLatestOrder(limit, offset uint64) (entity.StudyPosts, error)
	UpdatePost(post *entity.StudyPost) (*entity.StudyPost, error)
	DeletePost(post *entity.StudyPost) error
}
