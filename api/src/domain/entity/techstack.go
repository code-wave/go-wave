package entity

import (
	"encoding/json"
	"github.com/code-wave/go-wave/infrastructure/errors"
	"github.com/code-wave/go-wave/infrastructure/helpers"
)

type TechStacks []TechStack

type TechStack struct {
	ID       int64  `json:"id,omitempty"`
	TechName string `json:"tech_name"`
}

func (t *TechStack) Validate() *errors.RestErr {
	err := helpers.CheckStringMinChar(t.TechName, 1)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	return nil
}

func (t *TechStack) ResponseJSON() ([]byte, *errors.RestErr) {
	tJson, err := json.Marshal(t)
	if err != nil {
		return nil, errors.NewInternalServerError("marshal error" + err.Error())
	}
	return tJson, nil
}

func (t *TechStacks) ResponseJSON() ([]byte, *errors.RestErr) {
	tJson, err := json.Marshal(t)
	if err != nil {
		return nil, errors.NewInternalServerError("marshal error" + err.Error())
	}
	return tJson, nil
}
