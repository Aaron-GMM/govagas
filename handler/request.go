package handler

import (
	"fmt"
	"reflect"
)

func ErrParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required)", name, typ)
}

// Create request Opening
type CreateOpeningRequest struct {
	Name            string  `json:"name"`
	PublicationData string  `json:"publication_data"`
	Role            string  `json:"role"`
	Company         string  `json:"company"`
	Locate          string  `json:"locate"`
	Remote          *bool   `json:"remote"`
	Link            string  `json:"link"`
	Salary          float32 `json:"salary"`
	Description     string  `json:"description"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Name == "" && r.PublicationData == "" &&
		r.Role == "" && r.Company == "" &&
		r.Locate == "" && r.Remote == nil &&
		r.Link == "" && r.Salary <= 0 &&
		r.Description == "" {
		return fmt.Errorf("body is empty or malformed")
	}
	if r.Name == "" {
		return ErrParamIsRequired("Name", "string")
	}
	if r.PublicationData == "" {
		return ErrParamIsRequired("PublicationData", "string")
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

type UpadeteOpeningRequest struct {
	Name            *string  `json:"name"`
	PublicationData *string  `json:"PublicationData"`
	Role            *string  `json:"role"`
	Company         *string  `json:"company"`
	Locate          *string  `json:"locate"`
	Remote          *bool    `json:"remote"`
	Link            *string  `json:"link"`
	Salary          *float32 `json:"salary"`
	Description     *string  `json:"description"`
}

func (req *UpadeteOpeningRequest) Validate() error {
	value := reflect.ValueOf(req).Elem()

	for i := 0; i < value.NumField(); i++ {
		if !value.Field(i).IsNil() {
			return nil
		}
	}
	return fmt.Errorf("at least one param is required")
}
