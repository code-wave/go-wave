package interfaces

import (
	"encoding/json"
	"github.com/code-wave/go-wave/application"
	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/infrastructure/helpers"
	"net/http"
)

type TechStack struct {
	ts application.TechStackInterface
}

func NewTechStackHandler(ts application.TechStackInterface) *TechStack {
	return &TechStack{
		ts: ts,
	}
}

// SaveTechStack post로 tech_name 받아서 저장
func (t *TechStack) SaveTechStack(w http.ResponseWriter, r *http.Request) {
	helpers.SetJsonHeader(w)

	var techStack entity.TechStack

	json.NewDecoder(r.Body).Decode(&techStack)
	err := techStack.Validate()
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	err = t.ts.SaveTechStack(techStack.TechName)
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetTechStack id에 매칭되는 tech_name return
func (t *TechStack) GetTechStack(w http.ResponseWriter, r *http.Request) {
	helpers.SetJsonHeader(w)

	id, err := helpers.ExtractIntParam(r, "tech_stack_id")
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	techStack, err := t.ts.GetTechStack(id)
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	tJson, err := techStack.ResponseJSON()
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(tJson)
}

func (t *TechStack) GetAllTechStack(w http.ResponseWriter, r *http.Request) {
	helpers.SetJsonHeader(w)

	techStacks, err := t.ts.GetAllTechStack()
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	tJson, err := techStacks.ResponseJSON()
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(tJson)
}

func (t *TechStack) GetAllTechStackByStudyPostID(w http.ResponseWriter, r *http.Request) {
	helpers.SetJsonHeader(w)

	studyPostID, err := helpers.ExtractIntParam(r, "study_post_id")
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	techStacks, err := t.ts.GetAllTechStackByStudyPostID(studyPostID)
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	tJson, err := techStacks.ResponseJSON()
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(tJson)
}

func (t *TechStack) DeleteTechStack(w http.ResponseWriter, r *http.Request) {
	helpers.SetJsonHeader(w)

	techName := helpers.ExtractStringParam(r, "tech_name")

	err := t.ts.DeleteTechStack(techName)
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	w.WriteHeader(http.StatusOK)
}
