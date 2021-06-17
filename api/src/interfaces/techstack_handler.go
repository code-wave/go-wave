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

func NewTechStack(ts application.TechStackInterface) *TechStack {
	return &TechStack{
		ts: ts,
	}
}

// SaveTechStack post로 tech_name 받아서 저장
func (t *TechStack) SaveTechStack(w http.ResponseWriter, r *http.Request) {
	helpers.SetJsonHeader(w)

	var techStack entity.TechStack

	json.NewDecoder(r.Body).Decode(&techStack)

	err := t.ts.SaveTechStack(techStack.TechName)
	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.ResponseJSON().([]byte))
		return
	}

	w.WriteHeader(http.StatusOK)
}
