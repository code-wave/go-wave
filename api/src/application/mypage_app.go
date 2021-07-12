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

func getStudyPostForMypage(mypagePost entity.MypagePost, studyPost entity.StudyPost, writer string) entity.MypagePost {
	mypagePost.Writer = writer
	mypagePost.Title = studyPost.Title
	mypagePost.CreatedAt = studyPost.CreatedAt

	return mypagePost
}

func (m *mypageApp) GetMypageByUserIDAndStudyPostID(userID, limit, offset int64) (*entity.Mypage, *errors.RestErr) {
	var mypage entity.Mypage

	user, userErr := m.userRepo.GetUserByID(userID)
	if userErr != nil {
		log.Println("get user by id error at get my page application " + userErr.Message)
		return nil, userErr
	}
	//Get mypage user info
	mypage.User = entity.PublicUser{
		ID:       user.ID,
		Name:     user.Name,
		Nickname: user.Nickname,
	}

	studyPosts, studyPostErr := m.studyPostRepo.GetPostsByUserID(userID, limit, offset)
	if studyPostErr != nil {
		log.Println("get user by id error at get my page application " + studyPostErr.Message)
		return nil, studyPostErr
	}

	//Get stduyPosts list about writed
	var writedPosts []entity.MypagePost
	for _, studyPost := range studyPosts {
		var writedPost entity.MypagePost

		writedPost = getStudyPostForMypage(writedPost, studyPost, user.Nickname)
		writedPosts = append(writedPosts, writedPost)
	}
	mypage.WritedStudyPosts = writedPosts

	//수정 필요(만약 chatroom이 존재하지 않는다면 error 처리하는 방법 고려)
	//Get chatrooms by clientID and get studyPosts lists by chatrooms.StudyPost.ID...

	//Get studypost lists about participated
	var participatedPosts []entity.MypagePost
	for _, studyPost := range studyPosts {
		//Get writer nickame about participated post
		postWriter, postWriterErr := m.userRepo.GetUserByID(studyPost.UserID)
		if postWriterErr != nil {
			log.Println("get writer info by id error at get my page application " + postWriterErr.Message)
			return nil, postWriterErr
		}
		var participatedPost entity.MypagePost

		participatedPost = getStudyPostForMypage(participatedPost, studyPost, postWriter.Nickname)
		participatedPosts = append(participatedPosts, participatedPost)
	}
	mypage.ParticipatedStudyPosts = participatedPosts

	return &mypage, nil
}
