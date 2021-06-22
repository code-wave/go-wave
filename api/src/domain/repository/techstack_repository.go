package repository

import (
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/infrastructure/errors"
)

type TechStackRepository interface {
	SaveTechStack(techName string) *errors.RestErr
	GetTechStack(id int64) (*entity.TechStack, *errors.RestErr)
	GetAllTechStack() (entity.TechStacks, *errors.RestErr)
	GetAllTechStackByStudyPostID(studyPostID int64) (entity.TechStacks, *errors.RestErr)
	DeleteTechStack(techName string) *errors.RestErr
	CheckTechStack(techStack []string) *errors.RestErr
}
