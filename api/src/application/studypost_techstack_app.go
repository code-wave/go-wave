package application

//import (
//	"github.com/code-wave/go-wave/domain/repository"
//	"github.com/code-wave/go-wave/infrastructure/errors"
//)
//
//type studyPostTechStackApp struct {
//	studyPostTechStackRepo repository.StudyPostTechStackRepository
//}
//
//var _ StudyPostTechStackInterface = &studyPostTechStackApp{}
//
//type StudyPostTechStackInterface interface {
//	SaveStudyPostTechStack(studyPostID uint64, techStack []string) *errors.RestErr
//}
//
//func (t *studyPostTechStackApp) SaveStudyPostTechStack(studyPostID uint64, techStack []string) *errors.RestErr {
//	return t.studyPostTechStackRepo.SaveStudyPostTechStack(studyPostID, techStack)
//}
