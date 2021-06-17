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
	techIDList, restErr := s.getAllTechID(techStack)
	if restErr != nil {
		return errors.NewInternalServerError(restErr.Message)
	}

	query := s.insertAllTechStackQuery(studyPostID, techIDList)
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}

	_, err = stmt.Exec()
	if err != nil {
		return errors.NewInternalServerError("execute error")
	}

	return nil
}

// getAllTechID tech와 매칭되는 테이블의 id값을 모두 가져온 후 리스트로 반환
func (s *studyPostTechStackRepo) getAllTechID(techStack []string) ([]uint64, *errors.RestErr) {
	query := s.getAllTechStackIdQuery(len(techStack))
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, errors.NewInternalServerError("database error")
	}

	var tStack []interface{} // stmt.Query에 unpacking 하려면 []interface{}로 바꿔 넣어야함

	for i := 0; i < len(techStack); i++ {
		tStack = append(tStack, techStack[i])
	}

	rows, err := stmt.Query(tStack...)
	if err != nil {
		return nil, errors.NewInternalServerError("query error")
	}

	var techID uint64
	var techIDList []uint64

	for rows.Next() {
		err := rows.Scan(&techID)
		if err != nil {
			return nil, errors.NewInternalServerError("row scan error")
		}
		techIDList = append(techIDList, techID)
	}

	return techIDList, nil
}

// getAllTechStackIdQuery 매칭되는 techID를 한번에 모두 뽑기 위한 쿼리 ex) ... WHERE tech_stack=$1 or tech_stack=$2 ...
func (s *studyPostTechStackRepo) getAllTechStackIdQuery(n int) string {
	query := `SELECT id FROM tech_stack WHERE `

	for i := 0; i < n; i++ {
		query += "tech_stack=$" + strconv.Itoa(i+1)
		if i < n-1 {
			query += " or "
		}
	}
	return query
}

// insertAllTechStackQuery 여러개의 values를 한꺼번에 저장하기 위함 ex) ... VALUES (1, 2), (1, 3) ...
func (s *studyPostTechStackRepo) insertAllTechStackQuery(studyPostID uint64, techStackIDList []uint64) string {
	query := `INSERT INTO studypost_techstack VALUES `

	for i := 0; i < len(techStackIDList); i++ {
		query += fmt.Sprintf("(%d %d)", studyPostID, techStackIDList[i])
		if i < len(techStackIDList)-1 {
			query += ", "
		}
	}
	return query
}
