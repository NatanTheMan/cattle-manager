package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateCattleRequest struct {
	Name    string `json:"name"`
	Earring string `json:"earring"`
	Gender  string `json:"gender"`
}

func (r *CreateCattleRequest) Validate() error {
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Gender == "" {
		return errParamIsRequired("gender", "string")
	}
	if r.Earring == "" {
		return errParamIsRequired("erring", "string")
	}
	if len(r.Earring) < 6 || len(r.Earring) > 6 {
		return fmt.Errorf("earring must have 6 digits")
	}
	return nil
}
