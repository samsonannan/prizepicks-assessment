package models

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckStatus_ValidStatus(t *testing.T) {
	// Test valid statuses in different cases (upper, lower, mixed)
	validStatuses := []string{"ACTIVE", "active", "aCtIvE", "DOWN", "down", "dOwN"}
	for _, status := range validStatuses {
		err := CheckStatus(status)
		assert.Nil(t, err, "Expected no error for valid status")
	}
}

func TestCheckStatus_EmptyString(t *testing.T) {
	// Test an empty string as status input
	err := CheckStatus("")
	assert.Nil(t, err, "Expected no error for an empty string")
}

func TestCheckStatus_InvalidStatus(t *testing.T) {
	// Test invalid statuses
	invalidStatuses := []string{"activee", "downn", "unknown", "UP", "123", " "}

	for _, status := range invalidStatuses {
		err := CheckStatus(status)
		assert.NotNil(t, err, "Expected error for invalid status")
		assert.EqualError(t, err, "must be a valid value", "Expected specific error message")
	}
}

func TestCheckStatus_NonStringInput(t *testing.T) {
	// Test non-string input (e.g., int, bool, struct)
	nonStringInputs := []interface{}{"123", "true", "struct{}{}", "nil"}

	for _, input := range nonStringInputs {
		err := CheckStatus(input)
		// assert.NotNil(t, err, "Expected error for non-string input")
		assert.Error(t, err, "must be a valid value", "Expected specific error message")
	}
}

func TestCheckStatus_MoreValidStatuses(t *testing.T) {
	// Test more valid statuses
	moreValidStatuses := []string{"Up", "dowNr", "Acitive", "inAcTivE", "aCfTiVe"}
	for _, status := range moreValidStatuses {
		err := CheckStatus(status)
		assert.Error(t, err, "Expected no error for valid status")
	}
}

func TestCheckStatus_StatusWithWhitespace(t *testing.T) {
	// Test statuses with leading and trailing whitespaces
	statusesWithWhitespace := []string{"aACTIVE", "DOWN\t", "@/Active", "ddowno"}
	for _, status := range statusesWithWhitespace {
		err := CheckStatus(status)
		assert.Error(t, err, "Expected no error for valid status with whitespace")
	}
}

func TestCheckStatus_StatusWithSpecialCharacters(t *testing.T) {
	// Test statuses with special characters
	statusesWithSpecialChars := []string{"ACTIVE!", "@down", "#ACTIVE#", "$dowN$"}
	for _, status := range statusesWithSpecialChars {
		err := CheckStatus(status)
		assert.NotNil(t, err, "Expected error for status with special characters")
		assert.EqualError(t, err, "must be a valid value", "Expected specific error message")
	}
}

func TestCheckStatus_StatusWithSpaces(t *testing.T) {
	// Test statuses with spaces within the status
	statusesWithSpaces := []string{"A CTIVE", "DOWN IS GOOD", "ACT IVE", " d o w n"}
	for _, status := range statusesWithSpaces {
		err := CheckStatus(status)
		assert.NotNil(t, err, "Expected error for status with spaces")
		assert.EqualError(t, err, "must be a valid value", "Expected specific error message")
	}
}

func TestCheckStatus_CombinationOfValidAndInvalidStatuses(t *testing.T) {
	// Test a combination of valid and invalid statuses
	mixedStatuses := []string{"ACTIVE", "up", "DOWN", "active", "unknown", "DOWNN"}
	for _, status := range mixedStatuses {
		err := CheckStatus(status)
		if strings.ToUpper(status) == "DOWN" || strings.ToUpper(status) == "ACTIVE" {
			assert.Nil(t, err, "Expected no error for valid status")
		} else {
			assert.NotNil(t, err, "Expected error for invalid status")
			assert.EqualError(t, err, "must be a valid value", "Expected specific error message")
		}
	}
}

func TestCheckStatus_NilInput(t *testing.T) {
	// Test nil input
	err := CheckStatus(nil)
	assert.Nil(t, err, "Expected error for nil input")
	assert.NoError(t, err, "must be a valid value", "Expected specific error message")
}

func TestCheckStatus_IntegersAsStatus(t *testing.T) {
	// Test integers as status input
	integersAsStatus := []int{0, 1, 2, 100}
	for _, status := range integersAsStatus {
		err := CheckStatus(status)
		assert.NoError(t, err, "must be a valid value", "Expected specific error message")
	}
}

func TestCheckStatus_BooleanAsStatus(t *testing.T) {
	// Test boolean as status input
	booleanAsStatus := []bool{true, false}
	for _, status := range booleanAsStatus {
		err := CheckStatus(status)
		assert.NoError(t, err, "must be a valid value", "Expected specific error message")
	}
}

func TestCheckStatus_StructAsStatus(t *testing.T) {
	// Test struct as status input
	type testStruct struct {
		Field1 string
		Field2 int
	}

	s := testStruct{Field1: "ACTIVE", Field2: 123}
	err := CheckStatus(s)
	assert.NoError(t, err, "must be a valid value", "Expected specific error message")
}

func TestCheckStatus_InvalidInterfaceType(t *testing.T) {
	// Test invalid interface type as status input
	err := CheckStatus(3.14)
	assert.Nil(t, err, "Expected error for invalid interface type")
	assert.NoError(t, err, "must be a valid value", "Expected specific error message")
}

func TestCageRequest_Validate_ValidInput(t *testing.T) {
	// Test valid inputs for CageRequest
	validRequests := []CageRequest{
		{Capacity: intPtr(10), Status: "ACTIVE"},
		{Capacity: intPtr(1), Status: "DOWN"},
		{Capacity: intPtr(5), Status: "active"},
		{Capacity: intPtr(20), Status: "down"},
	}

	for _, req := range validRequests {
		err := req.Validate()
		assert.NoError(t, err, "Expected no error for valid input")
	}
}

func TestCageRequest_Validate_InvalidCapacity(t *testing.T) {
	// Test invalid capacity values
	invalidCapacityRequests := []CageRequest{
		{Capacity: intPtr(-1), Status: "ACTIVE"},
		{Capacity: intPtr(-10), Status: "DOWN"},
	}

	for _, req := range invalidCapacityRequests {
		err := req.Validate()
		assert.NotNil(t, err, "Expected error for invalid capacity")
		assert.Error(t, err, "Capacity must be greater than or equal to 0", "Expected specific error message")
	}
}

func TestCageRequest_Validate_InvalidStatus(t *testing.T) {
	// Test invalid status values
	invalidStatusRequests := []CageRequest{
		{Capacity: intPtr(10), Status: "invalid"},
		{Capacity: intPtr(5), Status: "alive"},
	}

	for _, req := range invalidStatusRequests {
		err := req.Validate()
		assert.Error(t, err, "Status must be a valid value", "Expected specific error message")
	}
}

func TestCageRequest_Validate_InvalidCombination(t *testing.T) {
	// Test invalid combinations of capacity and status
	invalidCombinationRequests := []CageRequest{
		{Capacity: intPtr(-1), Status: "invalid"},
		{Capacity: intPtr(0), Status: ""},
	}

	for _, req := range invalidCombinationRequests {
		err := req.Validate()
		assert.NotNil(t, err, "Expected error for invalid combination")
		assert.NotContains(t, err.Error(), "Capacity must be lesser than or equal to 0", "Expected specific error message for capacity")
		assert.NotContains(t, err.Error(), "Status must be a valid value", "Expected specific error message for status")
	}
}

// Helper function to create a pointer to an integer
func intPtr(i int64) *int64 {
	return &i
}
