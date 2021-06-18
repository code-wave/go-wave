package application

import (
	"github.com/code-wave/go-wave/domain/repository"
	"github.com/code-wave/go-wave/infrastructure/errors"
	"github.com/code-wave/go-wave/infrastructure/helpers"
)

type techStackApp struct {
	techStackRepo repository.TechStackRepository
}

var _ TechStackInterface = &techStackApp{}

type TechStackInterface interface {
	SaveTechStack(techName string) *errors.RestErr
}

func NewTechStackApp(techStackRepo repository.TechStackRepository) *techStackApp {
	return &techStackApp{
		techStackRepo: techStackRepo,
	}
}

func (t *techStackApp) SaveTechStack(techName string) *errors.RestErr {
	err := helpers.CheckStringMinChar(techName, 1)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	return t.techStackRepo.SaveTechStack(techName)
}
