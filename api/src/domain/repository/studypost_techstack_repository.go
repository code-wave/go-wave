package repository

import (
	"github.com/code-wave/go-wave/infrastructure/errors"
)

type StudyPostTechStackRepository interface {
	SaveStudyPostTechStack(studyPostID int64, techStack []string) *errors.RestErr
	//GetStudyPostTechStack(studyPostTechStack *entity.StudyPostTechStack) (*entity.StudyPostTechStack, *errors.RestErr)
	UpdateStudyPostTechStack(studyPostID int64, techStack []string) *errors.RestErr
}
