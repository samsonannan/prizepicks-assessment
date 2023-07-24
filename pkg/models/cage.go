package models

import (
	"errors"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent"
)

// CagesResponse represents the response structure for cage-related operations.
type CagesResponse struct {
	Status bool        `mapstructure:"status" json:"status"` // Status indicates the success or failure of the request.
	Data   []*ent.Cage `mapstructure:"data" json:"data"`     // Data contains a slice of pointers to 'ent.Cage' objects, which hold the cage-related data.
	Error  string      `mapstructure:"error" json:"error"`   // Error stores any error message in case of a failed request.
}

// CagesSuccessResponse creates a successful response object for cage-related operations.
// It takes a slice of 'ent.Cage' objects as input and returns a pointer to 'CagesResponse'.
func CagesSuccessResponse(Cages []*ent.Cage) *CagesResponse {
	// Construct and return the 'CagesResponse' object with the provided 'Cages' slice.
	return &CagesResponse{
		Status: true,  // Set the 'Status' field to 'true' to indicate a successful response.
		Data:   Cages, // Set the 'Data' field to the provided 'Cages' slice, containing the cage-related data.
		Error:  "",    // Set the 'Error' field to an empty string since there is no error in a successful response.
	}
}

// CageResponse represents the response structure for cage-related operations.
type CageResponse struct {
	Status bool      `mapstructure:"status" json:"status"` // Status indicates the success or failure of the request.
	Data   *ent.Cage `mapstructure:"data" json:"data"`     // Data holds a pointer to an 'ent.Cage' object, which represents cage-related data.
	Error  string    `mapstructure:"error" json:"error"`   // Error stores any error message in case of a failed request.
}

// CageRequest represents the request structure for cage-related operations.
type CageRequest struct {
	Capacity *int64 `mapstructure:"capacity" json:"capacity"` // Capacity holds a pointer to an 'int64' value, representing the capacity of a cage.
	Status   string `mapstructure:"status" json:"status"`     // Status holds a 'string' value, representing the status of a cage.
}

// CheckStatus validates the provided status value to ensure it is a valid status for a cage.
// The function takes an interface{} as input, which allows for flexibility in accepting different types.
// It attempts to assert the input status to a string type using type assertion, ignoring the second
// return value, which is a boolean indicating the success of the assertion.
// If the input status is an empty string or equal to "ACTIVE" or "DOWN" (case-insensitive),
// the function considers it as a valid status for a cage, and it returns nil (no error).
// If the input status is not a valid value, the function returns an error indicating that the
// status must be a valid value.
// Note: The function is lenient in accepting status values with different letter cases (e.g., "active" or "dOwN").
func CheckStatus(status interface{}) error {
	// Attempt to assert the input status to a string type using type assertion.
	s, _ := status.(string)

	// Check if the input status is an empty string or equal to "ACTIVE" or "DOWN" (case-insensitive).
	// If any of these conditions are met, consider it a valid status for a cage, and return nil (no error).
	if s == "" || strings.EqualFold(s, "ACTIVE") || strings.EqualFold(s, "DOWN") {
		return nil
	}

	// If the input status is not a valid value, return an error indicating that the status must be a valid value.
	return errors.New("must be a valid value")
}

// Validate checks the validity of the CageRequest data.
// It uses the validation library to perform various checks on the fields of the struct.
// For the 'Capacity' field, it ensures that it is not nil and is greater than or equal to 0.
// For the 'Status' field, it uses the custom validation function 'CheckStatus'.
// If any of the validations fail, an error is returned.
func (req CageRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Capacity, validation.Required, validation.Min(0)),
		validation.Field(&req.Status, validation.By(CheckStatus)),
	)
}

// ValidateIfUpdate checks the validity of the CageRequest data when updating.
// It uses the validation library to perform various checks on the fields of the struct.
// For the 'Capacity' field, it ensures that it is greater than or equal to 0 (optional check for updates).
// For the 'Status' field, it uses the custom validation function 'CheckStatus' (optional check for updates).
// If any of the validations fail, an error is returned.
func (req CageRequest) ValidateIfUpdate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Capacity, validation.Min(0)),
		validation.Field(&req.Status, validation.By(CheckStatus)),
	)
}

// CageSuccessResponse creates a successful response object for cage-related operations.
// It takes a pointer to an 'ent.Cage' object as input and returns a pointer to 'CageResponse'.
func CageSuccessResponse(Cage *ent.Cage) *CageResponse {
	// Construct and return the 'CageResponse' object with the provided 'Cage' pointer.
	return &CageResponse{
		Status: true, // Set the 'Status' field to 'true' to indicate a successful response.
		Data:   Cage, // Set the 'Data' field to the provided 'Cage' pointer, containing the cage-related data.
		Error:  "",   // Set the 'Error' field to an empty string since there is no error in a successful response.
	}
}

// CageErrorResponse creates an error response object for cage-related operations.
// It takes an error as input and returns a pointer to 'CageResponse'.
// The 'Status' field is set to 'false' to indicate a failed response.
// The 'Data' field is set to 'nil' since there is no relevant data in case of an error response.
// The 'Error' field is set to the error message obtained from the 'err' parameter.
func CageErrorResponse(err error) *CageResponse {
	return &CageResponse{
		Status: false,       // Set the 'Status' field to 'false' to indicate a failed response.
		Data:   nil,         // Set the 'Data' field to 'nil' since there is no relevant data in an error response.
		Error:  err.Error(), // Set the 'Error' field to the error message obtained from the 'err' parameter.
	}
}
