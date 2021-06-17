package repository

import "github.com/code-wave/go-wave/infrastructure/errors"

type StudyPostTechStackRepository interface {
	SaveStudyPostTechStack(studyPostID uint64, techStack []string) *errors.RestErr
	//GetStudyPostTechStack(studyPostTechStack *entity.StudyPostTechStack) (*entity.StudyPostTechStack, *errors.RestErr)
	//UpdateStudyPostTechStack(studyPostTechStack *entity.StudyPostTechStack) (*entity.StudyPostTechStack, *errors.RestErr)
}
