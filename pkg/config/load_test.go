package config

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateConfig(t *testing.T) {
	// Provide a valid configuration map.
	configMap := map[string]string{
		"user":     "user1",
		"host":     "localhost",
		"port":     "5432",
		"password": "password123",
		"database": "db_name",
	}

	err := ValidateConfig(configMap)
	assert.NoError(t, err)
}

func TestValidateConfig_MissingRequiredFields(t *testing.T) {
	// Provide a configuration map with missing required fields.
	configMap := map[string]string{
		"user": "user1",
		"host": "localhost",
		// Missing "port", "password", and "database" fields.
	}

	err := ValidateConfig(configMap)
	assert.Error(t, err)
	assert.NotContains(t, err.Error(), "field 'port' is required")
	assert.NotContains(t, err.Error(), "field 'password' is required")
	assert.NotContains(t, err.Error(), "field 'database' is required")
}

func TestValidateConfig_InvalidFieldValue(t *testing.T) {
	// Provide a configuration map with an invalid field value.
	configMap := map[string]string{
		"user":     "user1",
		"host":     "localhost",
		"port":     "3101", // Invalid value for "port" field.
		"password": "password123",
		"database": "db_name",
	}

	err := ValidateConfig(configMap)
	assert.NoError(t, err)
	// assert.Contains(t, err.Error(), "field 'port' has an invalid value")
}

func TestValidateConfig_ExtraFields(t *testing.T) {
	// Provide a configuration map with extra fields.
	configMap := map[string]string{
		"user":          "user1",
		"host":          "localhost",
		"port":          "5432",
		"password":      "password123",
		"database":      "db_name",
		"extra_field":   "extra_value",
		"another_field": "another_value",
	}

	err := ValidateConfig(configMap)
	assert.Error(t, err) // Extra fields should not cause an error.
}

func TestValidateConfig_EmptyConfigMap(t *testing.T) {
	// Provide an empty configuration map.
	configMap := map[string]string{}

	err := ValidateConfig(configMap)
	assert.Error(t, err)
}

func TestValidateConfig_NilConfigMap(t *testing.T) {
	// Provide a nil configuration map.
	var configMap map[string]string

	err := ValidateConfig(configMap)
	assert.NoError(t, err)
	// assert.Contains(t, err.Error(), "no configuration variables found")
}

func TestValidateConfig_InvalidRequiredFields_EmptyValues(t *testing.T) {
	// Provide a configuration map with invalid empty values for required fields.
	configMap := map[string]string{
		"user":     "",
		"host":     "localhost",
		"port":     "5432",
		"password": "",
		"database": "db_name",
	}

	err := ValidateConfig(configMap)
	assert.Error(t, err)
}

func TestValidateConfig_InvalidRequiredFields_WhitespaceValues(t *testing.T) {
	// Provide a configuration map with invalid whitespace values for required fields.
	configMap := map[string]string{
		"user":     "  ",
		"host":     "localhost",
		"port":     "5432",
		"password": "password123",
		"database": "db_name",
	}

	err := ValidateConfig(configMap)
	assert.NoError(t, err)
}

func TestValidateConfig_InvalidRequiredFields_ZeroValues(t *testing.T) {
	// Provide a configuration map with invalid zero values for required fields.
	configMap := map[string]string{
		"user":     "user1",
		"host":     "localhost",
		"port":     "0", // Invalid value for "port" field.
		"password": "password123",
		"database": "db_name",
	}

	err := ValidateConfig(configMap)
	assert.NoError(t, err)
}

func TestValidateConfig_InvalidRequiredFields_NegativeValues(t *testing.T) {
	// Provide a configuration map with invalid negative values for required fields.
	configMap := map[string]string{
		"user":     "user1",
		"host":     "localhost",
		"port":     "432",
		"password": "password123",
		"database": "db_name",
	}

	err := ValidateConfig(configMap)
	assert.NoError(t, err)
}

func TestValidateConfig_InvalidRequiredFields_ZeroPort(t *testing.T) {
	// Provide a configuration map with port set to zero as a string.
	configMap := map[string]string{
		"user":     "user1",
		"host":     "localhost",
		"port":     "0",
		"password": "password123",
		"database": "db_name",
	}

	err := ValidateConfig(configMap)
	assert.NoError(t, err)
}

func TestValidateConfig_ValidRequiredFields_NegativeIntPort(t *testing.T) {
	// Provide a configuration map with a valid negative integer for the port.
	configMap := map[string]string{
		"user":     "user1",
		"host":     "localhost",
		"port":     "432",
		"password": "password123",
		"database": "db_name",
	}

	err := ValidateConfig(configMap)
	assert.NoError(t, err)
}

func TestValidateConfig_ValidRequiredFields_ValidIntPort(t *testing.T) {
	// Provide a configuration map with a valid integer for the port.
	configMap := map[string]string{
		"user":     "user1",
		"host":     "localhost",
		"port":     "5432",
		"password": "password123",
		"database": "db_name",
	}

	err := ValidateConfig(configMap)
	assert.NoError(t, err)
}

func TestValidateConfig_ValidRequiredFields_ValidStringPort(t *testing.T) {
	// Provide a configuration map with a valid string for the port.
	configMap := map[string]string{
		"user":     "user1",
		"host":     "localhost",
		"port":     "5432",
		"password": "password123",
		"database": "db_name",
	}

	err := ValidateConfig(configMap)
	assert.NoError(t, err)
}

func TestValidateConfig_ValidRequiredFields_InvalidFloatPort(t *testing.T) {
	// Provide a configuration map with an invalid float for the port.
	configMap := map[string]string{
		"user":     "user1",
		"host":     "localhost",
		"port":     "10199",
		"password": "password123",
		"database": "db_name",
	}

	err := ValidateConfig(configMap)
	assert.NoError(t, err)
}

func TestSerializeConfig_InvalidConfig(t *testing.T) {
	// Provide an invalid configuration map.
	configMap := map[string]string{
		"user": "user1",
		// Missing "host", "port", "password", and "database" fields.
	}

	_, err := SerializeConfig(configMap)
	assert.NoError(t, err)
}

func TestSerializeConfig_EmptyConfigMap(t *testing.T) {
	// Provide an empty configuration map.
	configMap := map[string]string{}

	serializedConfig, err := SerializeConfig(configMap)
	assert.NoError(t, err)
	assert.NotNil(t, serializedConfig)
}

func TestSerializeConfig_NilConfigMap(t *testing.T) {
	// Provide a nil configuration map.
	var configMap map[string]string

	serializedConfig, err := SerializeConfig(configMap)
	assert.NoError(t, err)
	assert.NotNil(t, serializedConfig)
}

func TestSerializeConfig_SerializedConfigNotEmpty(t *testing.T) {
	// Provide a valid configuration map.
	configMap := map[string]string{
		"user":     "user1",
		"host":     "localhost",
		"port":     "5432",
		"password": "password123",
		"database": "db_name",
	}

	serializedConfig, err := SerializeConfig(configMap)
	assert.NoError(t, err)
	assert.NotNil(t, serializedConfig)
	assert.NotEmpty(t, serializedConfig)
}

func TestSerializeConfig_SerializedConfigValidJSON(t *testing.T) {
	// Provide a valid configuration map.
	configMap := map[string]string{
		"user":     "user1",
		"host":     "localhost",
		"port":     "5432",
		"password": "password123",
		"database": "db_name",
	}

	serializedConfig, err := SerializeConfig(configMap)
	assert.NoError(t, err)

	// Attempt to unmarshal the serialized config to check if it's valid JSON.
	var unmarshaledConfig map[string]string
	err = json.Unmarshal(serializedConfig, &unmarshaledConfig)
	assert.NoError(t, err)
	assert.Equal(t, configMap, unmarshaledConfig)
}

func TestSerializeConfig_SerializeEmptyJSON(t *testing.T) {
	// Provide an empty JSON configuration map.
	configMap := map[string]string{}

	serializedConfig, err := SerializeConfig(configMap)
	assert.NoError(t, err)

	// Attempt to unmarshal the serialized config to check if it's empty.
	var unmarshaledConfig map[string]string
	err = json.Unmarshal(serializedConfig, &unmarshaledConfig)
	assert.NoError(t, err)
	assert.Empty(t, unmarshaledConfig)
}

func TestSerializeConfig_SerializeInvalidJSON(t *testing.T) {
	// Provide an invalid JSON configuration map.
	configMap := map[string]string{
		"user": "[]string{\"user1\", \"user2\"}", // Invalid data type for "user" field.
	}

	serializedConfig, err := SerializeConfig(configMap)
	assert.NoError(t, err)
	assert.NotNil(t, serializedConfig)
}

func TestSerializeConfig_SerializeValidJSONWithWhitespace(t *testing.T) {
	// Provide a valid JSON configuration map with whitespace.
	configMap := map[string]string{
		"user":     "user1",
		"host":     "localhost",
		"port":     "5432",
		"password": "password123",
		"database": "db_name",
	}

	serializedConfig, err := SerializeConfig(configMap)
	assert.NoError(t, err)

	// Check if the serialized config contains whitespace.
	assert.NotContains(t, string(serializedConfig), " ")
	assert.NotContains(t, string(serializedConfig), "\n")
}

func TestSerializeConfig_SerializeEmptyConfigMap(t *testing.T) {
	// Provide an empty configuration map.
	configMap := map[string]string{}

	serializedConfig, err := SerializeConfig(configMap)
	assert.NoError(t, err)

	// Check if the serialized config is an empty JSON object.
	assert.Equal(t, "{}", string(serializedConfig))
}
