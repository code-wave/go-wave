package entity

import "github.com/code-wave/go-wave/infrastructure/errors"

type StudyPostTechStack struct {
	StudyPostID int64
	TechStackID int64
}

func (s *StudyPostTechStack) Validate() *errors.RestErr {
	if s.StudyPostID <= 0 {
		return errors.NewBadRequestError("wrong study_post id")
	}
	if s.TechStackID <= 0 {
		return errors.NewBadRequestError("wrong tech_stack id")
	}
	return nil
}
