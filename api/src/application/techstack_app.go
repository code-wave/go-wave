package application

import (
	"github.com/code-wave/go-wave/domain/entity"
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
	GetTechStack(id int64) (*entity.TechStack, *errors.RestErr)
	GetAllTechStack() (entity.TechStacks, *errors.RestErr)
	GetAllTechStackByStudyPostID(studyPostID int64) (entity.TechStacks, *errors.RestErr)
	DeleteTechStack(techName string) *errors.RestErr
}

func NewTechStackApp(techStackRepo repository.TechStackRepository) *techStackApp {
	return &techStackApp{
		techStackRepo: techStackRepo,
	}
}

func (t *techStackApp) SaveTechStack(techName string) *errors.RestErr {
	return t.techStackRepo.SaveTechStack(techName)
}

func (t *techStackApp) GetTechStack(id int64) (*entity.TechStack, *errors.RestErr) {
	return t.techStackRepo.GetTechStack(id)
}

func (t *techStackApp) GetAllTechStack() (entity.TechStacks, *errors.RestErr) {
	return t.techStackRepo.GetAllTechStack()
}

func (t *techStackApp) GetAllTechStackByStudyPostID(studyPostID int64) (entity.TechStacks, *errors.RestErr) {
	return t.techStackRepo.GetAllTechStackByStudyPostID(studyPostID)
}

func (t *techStackApp) DeleteTechStack(techName string) *errors.RestErr {
	err := helpers.CheckStringMinChar(techName, 1)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	return t.techStackRepo.DeleteTechStack(techName)
}
