package persistence

import (
	"database/sql"
	"fmt"
	"github.com/code-wave/go-wave/domain/repository"
	"github.com/code-wave/go-wave/infrastructure/errors"
)

type studyPostTechStackRepo struct {
	db *sql.DB
}

func NewStudyPostTechStackRepo(db *sql.DB) *studyPostTechStackRepo {
	return &studyPostTechStackRepo{db}
}

var _ repository.StudyPostTechStackRepository = &studyPostTechStackRepo{}

// SaveStudyPostTechStack (studyPostID, techStackID)의 형태로 인자로 받는 techStack 배열만큼 한번에 저장
func (s *studyPostTechStackRepo) SaveStudyPostTechStack(studyPostID int64, techStack []string) *errors.RestErr {
	query := s.insertAllTechStackQuery(studyPostID, techStack)

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return errors.NewInternalServerError("database error" + err.Error())
	}

	_, err = stmt.Exec()
	if err != nil {
		return errors.NewInternalServerError("execute error " + err.Error())
	}

	return nil
}

// insertAllTechStackQuery 여러개의 values를 한꺼번에 저장하기 위함
func (s *studyPostTechStackRepo) insertAllTechStackQuery(studyPostID int64, techStack []string) string { // TODO: 쿼리 만드는 함수
	query := fmt.Sprintf("INSERT INTO study_post_tech_stack SELECT %d, id FROM tech_stack WHERE tech_name IN (", studyPostID)
	for i := 0; i < len(techStack); i++ {
		query += "'" + techStack[i] + "'"
		if i < len(techStack)-1 {
			query += ", "
		}
	}
	query += ");"

	return query
}

func (s *studyPostTechStackRepo) GetAllTechStackQuery() {}

func (s *studyPostTechStackRepo) UpdateStudyPostTechStack(studyPostID int64, techStack []string) *errors.RestErr {
	stmt, err := s.db.Prepare(`
		DELETE FROM study_post_tech_stack
		WHERE study_post_id=$1
	`)
	if err != nil {
		return errors.NewInternalServerError("database error " + err.Error())
	}

	_, err = stmt.Exec(studyPostID)
	if err != nil {
		return errors.NewInternalServerError("execute error")
	}

	return s.SaveStudyPostTechStack(studyPostID, techStack)
}

//func (s *studyPostTechStackRepo) updateAllTechStackQuery(studyPostID int64, techStack []string) string {
//	// TODO: query
//	query := fmt.Sprintf("UPDATE study_post_tech_stack SET tech_stack_id=, id FROM tech_stack WHERE tech_name IN (", studyPostID)
//}
