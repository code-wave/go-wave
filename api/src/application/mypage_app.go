package application

import (
	"log"

	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/domain/repository"
	"github.com/code-wave/go-wave/infrastructure/errors"
)

var _ MypageAppInterface = &mypageApp{}

type mypageApp struct {
	userRepo      repository.UserRepository
	studyPostRepo repository.StudyPostRepository
}

type MypageAppInterface interface {
	GetMypageByUserIDAndStudyPostID(int64, int64, int64) (*entity.Mypage, *errors.RestErr)
}

func NewMypageApp(userRepo repository.UserRepository, studyPostRepo repository.StudyPostRepository) *mypageApp {
	return &mypageApp{
		userRepo:      userRepo,
		studyPostRepo: studyPostRepo,
	}
}

func (m *mypageApp) GetMypageByUserIDAndStudyPostID(userID, limit, offset int64) (*entity.Mypage, *errors.RestErr) {
	var mypage entity.Mypage
	user, userErr := m.userRepo.GetUserByID(userID)
	if userErr != nil {
		log.Println("get user by id error at get my page application " + userErr.Message)
		return nil, userErr
	}

	studyPosts, studyPostErr := m.studyPostRepo.GetPostsByUserID(userID, limit, offset)
	if studyPostErr != nil {
		log.Println("get user by id error at get my page application " + studyPostErr.Message)
		return nil, studyPostErr
	}

	mypage.User = entity.PublicUser{
		ID:       user.ID,
		Name:     user.Name,
		Nickname: user.Nickname,
	}

	var mypagePosts []entity.MypagePost
	for _, studyPost := range studyPosts {
		var mypage entity.MypagePost
		mypage.Title = studyPost.Title
		mypage.CreatedAt = studyPost.CreatedAt

		mypagePosts = append(mypagePosts, mypage)
	}

	mypage.StudyPosts = mypagePosts

	return &mypage, nil
}
