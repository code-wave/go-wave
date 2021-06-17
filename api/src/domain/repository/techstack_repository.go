package repository

import (
	"github.com/code-wave/go-wave/infrastructure/errors"
)

type TechStackRepository interface {
	SaveTechStack(techName string) *errors.RestErr
}
