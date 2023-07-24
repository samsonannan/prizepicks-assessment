package config

import (
	"encoding/json"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/spf13/viper"
)

// LoadConfigFile loads the configuration file using Viper.
// It sets the expected name, type, and path for the configuration file.
// If the file is found and successfully read, it returns nil.
// If there is an error reading the configuration file, it returns the error.
func LoadConfigFile(filename string) error {
	// Set the configuration file name to "config.toml".
	viper.SetConfigName(filename)
	// Set the configuration file type to TOML.
	viper.SetConfigType("toml")
	// Add the "pkg/config" path as one of the search paths for the configuration file.
	viper.AddConfigPath("./pkg/config")
	viper.AddConfigPath("/pkg/config")
	viper.AddConfigPath(".")

	// Read and load the configuration file.
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// Return nil if the configuration file was successfully read.
	return nil
}

// ValidateConfig validates the provided configuration map against a set of required fields.
// The function takes a configuration map as input and uses the validation package to check
// if the required fields ("user", "host", "port", "password", and "database") are present.
// If all the required fields are present with non-empty values, it returns nil.
// If any of the required fields are missing or have empty values, it returns an error
// indicating the presence of empty config variables.
func ValidateConfig(configMap map[string]string) error {
	// Use the validation package to check if the required fields in the configMap are present.
	err := validation.Validate(configMap,
		validation.Map(
			validation.Key("user", validation.Required),
			validation.Key("host", validation.Required),
			validation.Key("port", validation.Required),
			validation.Key("password", validation.Required),
			validation.Key("database", validation.Required),
		),
	)

	// Return nil if all the required fields are present and non-empty.
	// Return an error if any of the required fields are missing or have empty values.
	return err
}

// SerializeConfig serializes the provided configuration map into JSON format.
// The function takes a configuration map as input and uses the encoding/json package
// to marshal the map into a JSON byte slice ([]byte).
// If the marshaling is successful, it returns the JSON byte slice and nil error.
// If there is an error during marshaling, it returns an error indicating that it
// failed to serialize the config map.
func SerializeConfig(configMap map[string]string) ([]byte, error) {
	// Use the encoding/json package to marshal the configuration map into JSON format.
	pgb, err := json.Marshal(configMap)
	if err != nil {
		// Return an error if marshaling fails.
		return nil, err
	}

	// Return the JSON byte slice if marshaling is successful.
	return pgb, nil
}

// LoadDbConfig loads the database configuration by calling LoadConfigFile to read
// the configuration file and ValidateConfig to check if the required fields are present.
// It then uses SerializeConfig to serialize the configuration map into JSON format.
// If all these operations are successful, it returns the JSON byte slice and nil error.
// If any of these operations encounter an error, it returns an error with an appropriate message.
func LoadDbConfig(filename string) ([]byte, error) {
	// Load the configuration file using LoadConfigFile.
	err := LoadConfigFile(filename)
	if err != nil {
		return nil, errors.New("error reading config file")
	}

	// Get the "postgres" configuration map from the loaded configuration.
	postgresCfg := viper.GetStringMapString("postgres")

	// Validate the "postgres" configuration map using ValidateConfig.
	err = ValidateConfig(postgresCfg)
	if err != nil {
		return nil, errors.New("found empty config variables")
	}

	// Serialize the "postgres" configuration map into JSON format using SerializeConfig.
	pgb, err := SerializeConfig(postgresCfg)
	if err != nil {
		return nil, errors.New("failed to serialize config map")
	}

	// Return the JSON byte slice if all operations are successful.
	return pgb, nil
}
