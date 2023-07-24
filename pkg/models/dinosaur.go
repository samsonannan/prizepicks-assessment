package models

import (
	"errors"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent"
)

// DinosaursResponse represents the response structure for dinosaur-related operations.
type DinosaursResponse struct {
	Status bool            `mapstructure:"status" json:"status"` // Status indicates the success or failure of the request.
	Data   []*ent.Dinosaur `mapstructure:"data" json:"data"`     // Data contains a slice of pointers to 'ent.Dinosaur' objects, which hold the dinosaur-related data.
	Error  string          `mapstructure:"error" json:"error"`   // Error stores any error message in case of a failed request.
}

// DinosaursSuccessResponse creates a successful response object for dinosaur-related operations.
// It takes a slice of 'ent.Dinosaur' objects as input and returns a pointer to 'DinosaursResponse'.
func DinosaursSuccessResponse(dinosaurs []*ent.Dinosaur) *DinosaursResponse {
	// Construct and return the 'DinosaursResponse' object with the provided 'dinosaurs' slice.
	return &DinosaursResponse{
		Status: true,      // Set the 'Status' field to 'true' to indicate a successful response.
		Data:   dinosaurs, // Set the 'Data' field to the provided 'dinosaurs' slice, containing the dinosaur-related data.
		Error:  "",        // Set the 'Error' field to an empty string since there is no error in a successful response.
	}
}

// DinosaurRequest represents the request structure for dinosaur-related operations.
type DinosaurRequest struct {
	Name    string `mapstructure:"name" json:"name"`                 // Name holds the name of the dinosaur.
	Species string `mapstructure:"species" json:"species"`           // Species holds the species of the dinosaur.
	Group   string `mapstructure:"group" json:"group,omitempty"`     // Group holds the group of the dinosaur (optional).
	CageID  string `mapstructure:"cage_id" json:"cage_id,omitempty"` // CageID holds the ID of the cage where the dinosaur resides (optional).
}

// validateUUID checks if the provided cageId is a valid UUID.
// It takes a cageId of type interface{}, converts it to a string (if possible),
// and then attempts to parse it as a UUID using the uuid.Parse function.
// If the parsing fails, it returns an error indicating that the cageId is not a valid UUID.
func validateUUID(cageId interface{}) error {
	// Convert the cageId to a string (if possible).
	id, _ := cageId.(string)
	// Attempt to parse the id as a UUID using the uuid.Parse function.
	_, err := uuid.Parse(id)
	if err != nil {
		// If parsing fails, return an error indicating that the cageId is not a valid UUID.
		return errors.New("must be a valid value")
	}
	// Return nil if the cageId is a valid UUID.
	return nil
}

func CheckGroup(group interface{}) error {
	// Attempt to assert the input group to a string type using type assertion.
	g, ok := group.(string)
	if !ok {
		return errors.New("must be a valid value")
	}

	// Check if the input group is an empty string or equal to "HERBIVORE" or "CARNIVORE" (case-insensitive).
	// If any of these conditions are met, consider it a valid group for a dinosaur, and return nil (no error).
	if g == "" || strings.EqualFold(g, "HERBIVORE") || strings.EqualFold(g, "CARNIVORE") {
		return nil
	}

	// If the input group is not a valid value, return an error indicating that the status must be a valid value.
	return errors.New("must be a valid value")
}

// ValidateWithCageIdRequired checks the validity of the DinosaurRequest data,
// specifically for cases where the CageID field is required.
// It uses the validation library to perform various checks on the fields of the struct.
// For the 'Name' and 'Species' fields, it ensures they are not empty (required fields).
// For the 'Group' field, it validates that the value is either "HERBIVORE" or "CARNIVORE".
// For the 'CageID' field, it ensures it is not empty (required) and passes the custom validation function 'validateUUID'.
// If any of the validations fail, an error is returned.
func (req DinosaurRequest) ValidateWithCageIdRequired() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required),                                // Ensure 'Name' field is not empty (required).
		validation.Field(&req.Species, validation.Required),                             // Ensure 'Species' field is not empty (required).
		validation.Field(&req.Group, validation.By(CheckGroup)),                         // Ensure 'Group' field is one of the specified values.
		validation.Field(&req.CageID, validation.Required, validation.By(validateUUID)), // Ensure 'CageID' field is not empty (required) and a valid UUID.
	)
}

// ValidateIfUpdate checks the validity of the DinosaurRequest data,
// specifically for cases where the CageID field is required.
// It uses the validation library to perform various checks on the fields of the struct.
// For the 'Name' and 'Species' fields, it ensures they are not empty (required fields).
// For the 'Group' field, it validates that the value is either "HERBIVORE" or "CARNIVORE".
// For the 'CageID' field, it ensures it is not empty (required) and passes the custom validation function 'validateUUID'.
// If any of the validations fail, an error is returned.
func (req DinosaurRequest) ValidateIfUpdate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Group, validation.By(CheckGroup)),    // Ensure 'Group' field is one of the specified values.
		validation.Field(&req.CageID, validation.By(validateUUID)), // Ensure 'CageID' field is not empty (required) and a valid UUID.
	)
}

// DinosaurResponse represents the response structure for dinosaur-related operations.
type DinosaurResponse struct {
	Status bool          `mapstructure:"status" json:"status"` // Status indicates the success or failure of the request.
	Data   *ent.Dinosaur `mapstructure:"data" json:"data"`     // Data holds a pointer to an 'ent.Dinosaur' object, which represents dinosaur-related data.
	Error  string        `mapstructure:"error" json:"error"`   // Error stores any error message in case of a failed request.
}

// DinosaurSuccessResponse creates a successful response object for dinosaur-related operations.
// It takes a pointer to an 'ent.Dinosaur' object as input and returns a pointer to 'DinosaurResponse'.
func DinosaurSuccessResponse(dinosaur *ent.Dinosaur) *DinosaurResponse {
	// Construct and return the 'DinosaurResponse' object with the provided 'dinosaur' pointer.
	return &DinosaurResponse{
		Status: true,     // Set the 'Status' field to 'true' to indicate a successful response.
		Data:   dinosaur, // Set the 'Data' field to the provided 'dinosaur' pointer, containing the dinosaur-related data.
		Error:  "",       // Set the 'Error' field to an empty string since there is no error in a successful response.
	}
}

// DinosaurErrorResponse creates an error response object for dinosaur-related operations.
// It takes an error as input and returns a pointer to 'DinosaurResponse'.
// The 'Status' field is set to 'false' to indicate a failed response.
// The 'Data' field is set to 'nil' since there is no relevant data in case of an error response.
// The 'Error' field is set to the error message obtained from the 'err' parameter.
func DinosaurErrorResponse(err error) *DinosaurResponse {
	return &DinosaurResponse{
		Status: false,       // Set the 'Status' field to 'false' to indicate a failed response.
		Data:   nil,         // Set the 'Data' field to 'nil' since there is no relevant data in an error response.
		Error:  err.Error(), // Set the 'Error' field to the error message obtained from the 'err' parameter.
	}
}
