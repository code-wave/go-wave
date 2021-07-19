package persistence

import (
	"database/sql"
	"fmt"
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
func (t *techStackRepo) SaveTechStack(techName string) *errors.RestErr {
	stmt, err := t.db.Prepare(`
		INSERT INTO tech_stack (tech_name)
		VALUES ($1);
	`)
	if err != nil {
		return errors.NewInternalServerError("database error " + err.Error())
	}

	_, err = stmt.Exec(techName)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (t *techStackRepo) GetTechStack(id int64) (*entity.TechStack, *errors.RestErr) {
	stmt, err := t.db.Prepare(`
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

func (t *techStackRepo) GetAllTechStack() (entity.TechStacks, *errors.RestErr) {
	stmt, err := t.db.Prepare(`
		SELECT tech_name
		FROM tech_stack;
	`)
	if err != nil {
		return nil, errors.NewInternalServerError("database error" + err.Error())
	}

	rows, err := stmt.Query()
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

func (t *techStackRepo) GetAllTechStackByStudyPostID(studyPostID int64) (entity.TechStacks, *errors.RestErr) {
	stmt, err := t.db.Prepare(`
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

func (t *techStackRepo) DeleteTechStack(techName string) *errors.RestErr {
	stmt, err := t.db.Prepare(`
		DELETE FROM tech_stack
		WHERE tech_name=$1
	`)
	if err != nil {
		return errors.NewInternalServerError("database error" + err.Error())
	}

	res, err := stmt.Exec(techName)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.NewBadRequestError(err.Error())
		}
		return errors.NewInternalServerError("execute error" + err.Error())
	}

	n, err := res.RowsAffected()
	if err != nil {
		return errors.NewInternalServerError("rows affected error" + err.Error())
	}

	if n == 0 {
		return errors.NewBadRequestError("no rows to be deleted")
	}
	return nil
}

func (t *techStackRepo) CheckTechStack(techStack []string) *errors.RestErr {
	query := t.checkTechStackQuery(techStack)
	stmt, err := t.db.Prepare(query)
	if err != nil {
		return errors.NewInternalServerError("database error " + err.Error())
	}

	res, err := stmt.Exec()
	if err != nil {
		return errors.NewInternalServerError("database error " + err.Error())
	}

	n, err := res.RowsAffected()
	if err != nil {
		return errors.NewInternalServerError("database error " + err.Error())
	}

	if n < int64(len(techStack)) { // 개수가 적다면 tech_stack Table에 저장되어있지 않은 tech를 요청한 것
		return errors.NewBadRequestError("some tech name is not stored in the table")
	}

	return nil
}

// checkTechStackQuery ex) ... WHERE tech_name IN ('go', 'react')
func (t *techStackRepo) checkTechStackQuery(techStack []string) string {
	query := `SELECT * FROM tech_stack WHERE tech_name IN (`

	for i := 0; i < len(techStack); i++ {
		techName := techStack[i]
		query += fmt.Sprintf("'%s'", techName)
		if i < len(techStack)-1 {
			query += ", "
		}
	}
	query += ");"

	return query
}
