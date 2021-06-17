package interfaces

import (
	"encoding/json"
	"fmt"
	"github.com/code-wave/go-wave/application"
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/infrastructure/errors"
	"github.com/code-wave/go-wave/infrastructure/helpers"
	"net/http"
)

type StudyPost struct {
	sp application.StudyPostInterface
}

func NewStudyPost(sp application.StudyPostInterface) *StudyPost {
	return &StudyPost{
		sp,
	}
}

func (s *StudyPost) GetPost(w http.ResponseWriter, r *http.Request) {
	helpers.SetJsonHeader(w)

	postID, err := helpers.ExtractUintParam(r, "post_id")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	studyPost, restErr := s.sp.GetPost(postID)
	if restErr != nil {
		http.Error(w, "error", 500) // TODO: 나중에 더 디테일하게
		return
	}

	err = json.NewEncoder(w).Encode(&studyPost)
	if err != nil {
		http.Error(w, "error", 500)
		return
	}
}

func (s *StudyPost) GetPosts(w http.ResponseWriter, r *http.Request) {
	helpers.SetJsonHeader(w)

	limit, err := helpers.ExtractUintParam(r, "limit")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	offset, err := helpers.ExtractUintParam(r, "offset")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	studyPosts, restErr := s.sp.GetPostsInLatestOrder(limit, offset)
	if restErr != nil {
		http.Error(w, "error", 500)
		return
	}

	err = json.NewEncoder(w).Encode(&studyPosts)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (s *StudyPost) SavePost(w http.ResponseWriter, r *http.Request) {
	helpers.SetJsonHeader(w)

	var studyPost entity.StudyPost

	err := json.NewDecoder(r.Body).Decode(&studyPost)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid json body") // TODO: 에러명 나중에 수정 (아래 에러들도)
		w.WriteHeader(restErr.Status)
		w.Write(restErr.ResponseJSON().([]byte))
		return
	}

	fmt.Println("studypost: ", studyPost) // check용 나중에 삭제

	restErr := studyPost.Validate()
	if restErr != nil {
		w.WriteHeader(restErr.Status)
		w.Write(restErr.ResponseJSON().([]byte))
		return
	}

	newPost, restErr := s.sp.SavePost(&studyPost) // TODO: 여기 할 차례
	if restErr != nil {
		http.Error(w, "save post error", 500)
	}

	err = json.NewEncoder(w).Encode(&newPost)
	if err != nil {
		http.Error(w, "encode error", 500)
	}
}
