package handler

import "fmt"

func ErrParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required)", name, typ)
}

// Create request Opening
type CreateOpeningRequest struct {
	Name             string  `json:"name"`
	Publication_data string  `json:"publication_data"`
	Role             string  `json:"role"`
	Company          string  `json:"company"`
	Locate           string  `json:"locate"`
	Remote           *bool   `json:"remote"`
	Link             string  `json:"link"`
	Salary           float32 `json:"salary"`
	Description      string  `json:"description"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Name == "" && r.Publication_data == "" &&
		r.Role == "" && r.Company == "" &&
		r.Locate == "" && r.Remote == nil &&
		r.Link == "" && r.Salary <= 0 &&
		r.Description == "" {
		return fmt.Errorf("body is empty or malformed")
	}
	if r.Name == "" {
		return ErrParamIsRequired("Name", "string")
	}
	if r.Publication_data == "" {
		return ErrParamIsRequired("Publication_data", "string")
	}
	if r.Role == "" {
		return ErrParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return ErrParamIsRequired("Company", "string")
	}
	if r.Locate == "" {
		return ErrParamIsRequired("Locate", "string")
	}
	if r.Link == "" {
		return ErrParamIsRequired("Link", "string")
	}

	if r.Salary <= 0.0 {
		return ErrParamIsRequired("Salary", "float32")
	}
	if r.Remote == nil {
		return ErrParamIsRequired("Remote", "bool")
	}
	return nil
}
