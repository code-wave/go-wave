package persistence

import (
	"database/sql"
	"github.com/code-wave/go-wave/domain/repository"
	"github.com/code-wave/go-wave/infrastructure/errors"
)

type techStackRepo struct {
	db *sql.DB
}

func NewTechStackRepo(db *sql.DB) *techStackRepo {
	return &techStackRepo{db}
}

var _ repository.TechStackRepository = &techStackRepo{}

// SaveTechStack 나중에 추가로 필요한 기술들 외부에서 추가 가능하게 하기 위함 예를 들어 기존 테이블에 TypeScript가 없다면 추가 가능
func (s *techStackRepo) SaveTechStack(techName string) *errors.RestErr {
	stmt, err := s.db.Prepare(`
		INSERT INTO tech_stack (name)
		VALUES ($1);
	`)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}

	_, err = stmt.Exec(techName)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (s *techStackRepo) GetTechStackByStudyPostID(studyPostID uint64) ([]string, *errors.RestErr) {
	stmt, err := s.db.Prepare(`
		SELECT tech_name 
		FROM tech_stack 
		WHERE id IN (SELECT tech_stack_id FROM study_post_tech_stack WHERE study_post_id=$1)
	`)
	if err != nil {
		return nil, errors.NewInternalServerError("database error")
	}

	rows, err := stmt.Query(studyPostID)
	if err != nil {
		return nil, errors.NewInternalServerError("query error")
	}

	var techStack []string

	for rows.Next() {
		var techName string
		err := rows.Scan(&techName)
		if err != nil {
			return nil, errors.NewInternalServerError("scan error")
		}

		techStack = append(techStack, techName)
	}

	return techStack, nil
}
