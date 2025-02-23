package handler

import (
	"fmt"
	"strconv"
	"strings"
)

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
	if len(r.Gender) > 1 {
		return fmt.Errorf("gender must have 1 digits")
	}
	if strings.ToUpper(r.Gender) != "M" && strings.ToUpper(r.Gender) != "F" {
		return fmt.Errorf("gender have to be 'M' or 'F'")
	}
	if r.Earring == "" {
		return errParamIsRequired("erring", "string")
	}
	if len(r.Earring) < 6 || len(r.Earring) > 6 {
		return fmt.Errorf("earring must have 6 digits")
	}
	_, err := strconv.Atoi(r.Earring)
	if err != nil {
		return fmt.Errorf("earring have to be a number")
	}
	return nil
}

type UpdateCattleRequest struct {
	Name string `json:"name"`
}

func (r *UpdateCattleRequest) Validate() error {
	if r.Name != "" {
		return nil
	}
	logger.Errorf("name is: %v", r.Name)
	return fmt.Errorf("name must be provided")
}
