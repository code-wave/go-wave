package persistence

import (
	"database/sql"
	"errors"
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/domain/repository"
	"github.com/code-wave/go-wave/infrastructure/helpers"
)

type studyPostRepo struct {
	db *sql.DB
}

func NewStudyPostRepo(db *sql.DB) *studyPostRepo {
	return &studyPostRepo{db}
}

var _ repository.StudyPostRepository = &studyPostRepo{}

func (s *studyPostRepo) SavePost(studyPost *entity.StudyPost) (*entity.StudyPost, error) {
	stmt, err := s.db.Prepare(`
		INSERT INTO study_post (user_id, title, topic, content, num_of_members, is_mento, price, start_date, end_date, is_online, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`)
	if err != nil {
		return nil, errors.New("insert error") //TODO: 나중에 errors 폴더로 처리?
	}

	currentTime := helpers.GetCurrentTimeForDB()

	_, err = stmt.Exec(studyPost.UserID, studyPost.Title, studyPost.Topic, studyPost.Content,
		studyPost.NumOfMembers, studyPost.IsMento, studyPost.Price, studyPost.StartDate, studyPost.EndDate,
		currentTime, currentTime)
	if err != nil {
		return nil, errors.New("execute error") // 얘도 위처럼 나중에 처리
	}

	return studyPost, nil
}

func (s *studyPostRepo) GetPost(id uint64) (*entity.StudyPost, error) {
	row := s.db.QueryRow(`
		SELECT id, user_id, title, topic, content, num_of_members, is_mento, price, start_date, 
		       end_date, is_online, created_at, updated_at
		FROM study_post
		WHERE id=$1
	`, id)

	err := row.Err()
	if err != nil {
		return nil, err // TODO: 나중에 에러처리 보완?
	}

	var studyPost entity.StudyPost

	err = row.Scan(&studyPost.ID, &studyPost.UserID, &studyPost.Title, &studyPost.Topic, &studyPost.NumOfMembers,
		&studyPost.IsMento, &studyPost.Price, &studyPost.StartDate, &studyPost.EndDate, &studyPost.IsOnline,
		&studyPost.CreatedAt, &studyPost.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &studyPost, nil
}

func (s *studyPostRepo) GetPostsInLatestOrder(limit uint64) (entity.StudyPosts, error) { // TODO: uint64 맞나? (프론트와 통신할때)
	rows, err := s.db.Query(`
		SELECT id, user_id, title, topic, content, num_of_members, is_mento, price, start_date, 
		       end_date, is_online, created_at, updated_at
		FROM study_post
		ORDER BY created_at DESC
		LIMIT $1
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var studyPosts entity.StudyPosts

	for rows.Next() {
		var studyPost entity.StudyPost

		err := rows.Scan(&studyPost.ID, &studyPost.UserID, &studyPost.Title, &studyPost.Topic, &studyPost.NumOfMembers,
			&studyPost.IsMento, &studyPost.Price, &studyPost.StartDate, &studyPost.EndDate, &studyPost.IsOnline,
			&studyPost.CreatedAt, &studyPost.UpdatedAt)
		if err != nil {
			return nil, err
		}
		studyPosts = append(studyPosts, studyPost)
	}

	if err = rows.Err(); err != nil { // 끝난 후에도 에러체크 한번
		return nil, err
	}

	return studyPosts, nil
}
