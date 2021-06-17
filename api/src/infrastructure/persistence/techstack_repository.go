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
