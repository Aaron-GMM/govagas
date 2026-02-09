package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
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
		return errParamIsRequired("Name", "string")
	}
	if r.Publication_data == "" {
		return errParamIsRequired("Publication_data", "string")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("Company", "string")
	}
	if r.Locate == "" {
		return errParamIsRequired("Locate", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("Link", "string")
	}

	if r.Salary <= 0.0 {
		return errParamIsRequired("Salary", "float32")
	}
	if r.Remote == nil {
		return errParamIsRequired("Remote", "bool")
	}
	return nil
}
