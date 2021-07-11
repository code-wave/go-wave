package interfaces

import (
	"log"
	"net/http"

	"github.com/code-wave/go-wave/application"
	"github.com/code-wave/go-wave/infrastructure/helpers"
	"github.com/code-wave/go-wave/interfaces/middleware"
)

type MypageHandler struct {
	mypageApp application.MypageAppInterface
}

func NewMyPageHandler(mypageApp application.MypageAppInterface) *MypageHandler {
	return &MypageHandler{mypageApp: mypageApp}
}

func (mypageHandler *MypageHandler) GetMypage(w http.ResponseWriter, r *http.Request) {
	helpers.SetJsonHeader(w)

	authUserID := r.Context().Value(middleware.ContextKeyTokenUserID)
	log.Println("get userID from token middleware", authUserID)

	userID, err := helpers.ExtractIntParam(r, "user_id")
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	limit, err := helpers.ExtractIntParam(r, "limit")
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	offset, err := helpers.ExtractIntParam(r, "offset")
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	mypage, err := mypageHandler.mypageApp.GetMypageByUserIDAndStudyPostID(userID, limit, offset)
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	mypageRes, err := mypage.ResponseJSON()
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(mypageRes.([]byte))
}
