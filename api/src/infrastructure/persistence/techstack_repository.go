package persistence

import (
	"database/sql"
	"github.com/code-wave/go-wave/domain/entity"
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
		INSERT INTO tech_stack (tech_name)
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

func (s *techStackRepo) GetTechStack(id int64) (*entity.TechStack, *errors.RestErr) {
	stmt, err := s.db.Prepare(`
		SELECT tech_name
		FROM tech_stack
		WHERE id=$1;
	`)
	if err != nil {
		return nil, errors.NewInternalServerError("database error: " + err.Error())
	}

	var techStack entity.TechStack

	err = stmt.QueryRow(id).Scan(&techStack.TechName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewBadRequestError(err.Error())
		}
		return nil, errors.NewInternalServerError("scan error " + err.Error())
	}

	return &techStack, nil
}

func (s *techStackRepo) GetAllTechStackByStudyPostID(studyPostID int64) (entity.TechStacks, *errors.RestErr) {
	stmt, err := s.db.Prepare(`
		SELECT tech_name
		FROM tech_stack
		WHERE id IN (SELECT tech_stack_id FROM study_post_tech_stack WHERE study_post_id=$1);
	`)
	if err != nil {
		return nil, errors.NewInternalServerError("database error" + err.Error())
	}

	rows, err := stmt.Query(studyPostID)
	if err != nil {
		return nil, errors.NewInternalServerError("query error" + err.Error())
	}

	var techStacks entity.TechStacks

	for rows.Next() {
		var techStack entity.TechStack
		err := rows.Scan(&techStack.TechName)
		if err != nil {
			return nil, errors.NewInternalServerError("scan error" + err.Error())
		}

		techStacks = append(techStacks, techStack)
	}

	return techStacks, nil
}
