package persistence

import (
	"database/sql"
	"fmt"
	"github.com/code-wave/go-wave/domain/repository"
	"github.com/code-wave/go-wave/infrastructure/errors"
	"strconv"
)

type studyPostTechStackRepo struct {
	db *sql.DB
}

func NewStudyPostTechStackRepo(db *sql.DB) *studyPostTechStackRepo {
	return &studyPostTechStackRepo{db}
}

var _ repository.StudyPostTechStackRepository = &studyPostTechStackRepo{}

// SaveStudyPostTechStack (studyPostID, techStackID)의 형태로 인자로 받는 techStack 배열만큼 한번에 저장
func (s *studyPostTechStackRepo) SaveStudyPostTechStack(studyPostID uint64, techStack []string) *errors.RestErr {
	query := s.insertAllTechStackQuery(studyPostID, techStack)

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errors.NewInternalServerError("database error")
	}

	var tStack []interface{} // stmt.Query에 unpacking 하려면 []interface{}로 바꿔 넣어야함

	for i := 0; i < len(techStack); i++ {
		tStack = append(tStack, techStack[i])
	}

	_, err = stmt.Exec(tStack...) // TODO: interace 부분때매 걸림
	if err != nil {
		errors.NewInternalServerError("execute error")
	}

	return nil
}

// insertAllTechStackQuery 여러개의 values를 한꺼번에 저장하기 위함
func (s *studyPostTechStackRepo) insertAllTechStackQuery(studyPostID uint64, techStack []string) string {
	query := fmt.Sprintf("INSERT INTO study_post_tech_stack SELECT %d, id FROM tech_stack WHERE tech_name in (", studyPostID)

	for i := 0; i < len(techStack); i++ {
		query += "'$" + strconv.Itoa(i+1) + "'" // 'go' 이런형태로 ''가 있어야 하므로
		if i < len(techStack)-1 {
			query += ", "
		}
	}
	query += ");"

	return query
}

func (s *studyPostTechStackRepo) GetAllTechStackQuery() {}
