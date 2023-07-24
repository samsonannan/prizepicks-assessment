package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUUID_ValidUUID(t *testing.T) {
	// Test valid UUIDs
	validUUIDs := []interface{}{
		"6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		"6ba7b8109dad11d180b400c04fd430c8",
	}

	for _, id := range validUUIDs {
		err := validateUUID(id)
		assert.Nil(t, err, "Expected no error for valid UUID")
	}
}

func TestValidateUUID_InvalidUUID(t *testing.T) {
	// Test invalid UUIDs
	invalidUUIDs := []interface{}{
		"invalid",
		"some-uuid-input-case-42042",
	}

	for _, id := range invalidUUIDs {
		err := validateUUID(id)
		// assert.NotNil(t, err, "Expected error for invalid UUID")
		assert.Error(t, err, "could not parse cage id as UUID", "Expected specific error message")
	}
}

func TestValidateUUID_Nil(t *testing.T) {
	// Test nil input
	err := validateUUID(nil)
	// assert.Nil(t, err, "Expected error for nil input")
	assert.Error(t, err, "could not parse cage id as UUID", "Expected specific error message")
}

func TestDinosaurRequest_ValidateWithCageIdRequired_ValidInput(t *testing.T) {
	// Test valid inputs for DinosaurRequest with CageID required
	validRequests := []*DinosaurRequest{
		{Name: "T-Rex", Species: "Tyrannosaurus", Group: "CARNIVORE", CageID: "6ba7b810-9dad-11d1-80b4-00c04fd430c8"},
		{Name: "Triceratops", Species: "Triceratops", Group: "HERBIVORE", CageID: "6ba7b810-9dad-11d1-80b4-00c04fd430c8"},
	}

	for _, req := range validRequests {
		err := req.ValidateWithCageIdRequired()
		assert.Nil(t, err, "Expected no error for valid input")
	}
}

func TestDinosaurRequest_ValidateWithCageIdRequired_InvalidName(t *testing.T) {
	// Test invalid 'Name' field (empty)
	invalidNameRequests := []*DinosaurRequest{
		{Name: "", Species: "Tyrannosaurus", Group: "CARNIVORE", CageID: "6ba7b810-9dad-11d1-80b4-00c04fd430c8"},
	}

	for _, req := range invalidNameRequests {
		err := req.ValidateWithCageIdRequired()
		assert.NotNil(t, err, "Expected error for invalid 'Name'")
		assert.NotContains(t, err.Error(), "Name is not required", "Expected specific error message for 'Name'")
	}
}

func TestDinosaurRequest_ValidateWithCageIdRequired_InvalidSpecies(t *testing.T) {
	// Test invalid 'Species' field (empty)
	invalidSpeciesRequests := []*DinosaurRequest{
		{Name: "T-Rex", Species: "", Group: "CARNIVORE", CageID: "6ba7b810-9dad-11d1-80b4-00c04fd430c8"},
	}

	for _, req := range invalidSpeciesRequests {
		err := req.ValidateWithCageIdRequired()
		assert.NotNil(t, err, "Expected error for invalid 'Species'")
		assert.NotContains(t, err.Error(), "Species is not required", "Expected specific error message for 'Species'")
	}
}

func TestDinosaurRequest_ValidateWithCageIdRequired_InvalidGroup(t *testing.T) {
	// Test invalid 'Group' field (not "HERBIVORE" or "CARNIVORE")
	invalidGroupRequests := []*DinosaurRequest{
		{Name: "T-Rex", Species: "Tyrannosaurus", Group: "CARNIVORES", CageID: "6ba7b810-9dad-11d1-80b4-00c04fd430c8"},
	}

	for _, req := range invalidGroupRequests {
		err := req.ValidateWithCageIdRequired()
		assert.NotNil(t, err, "Expected error for invalid 'Group'")
		assert.NotContains(t, err.Error(), "Group must be one of the specified values", "Expected specific error message for 'Group'")
	}
}

func TestDinosaurRequest_ValidateWithCageIdRequired_InvalidCageID(t *testing.T) {
	// Test invalid 'CageID' field (empty or invalid UUID)
	invalidCageIDRequests := []*DinosaurRequest{
		{Name: "T-Rex", Species: "Tyrannosaurus", Group: "CARNIVORE", CageID: ""},
		{Name: "T-Rex", Species: "Tyrannosaurus", Group: "CARNIVORE", CageID: "invalidUUID"},
	}

	for _, req := range invalidCageIDRequests {
		err := req.ValidateWithCageIdRequired()
		assert.NotNil(t, err, "Expected error for invalid 'CageID'")
		assert.NotContains(t, err.Error(), "CageID is not required", "Expected specific error message for 'CageID'")
	}
}
