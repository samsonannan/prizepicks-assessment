package logger

import (
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

// SugaredLogger is a global variable that holds an instance of a SugaredLogger
// from the OpenTelemetry Zap integration package.
var SugaredLogger *otelzap.SugaredLogger

// SetGlobalLogger initializes the global SugaredLogger variable with a new
// SugaredLogger instance based on the provided Zap configuration.
// The function takes a Zap configuration as input and returns an error if there
// is an issue building the logger.
// The SugaredLogger is configured with options for OpenTelemetry integration,
// such as adding trace IDs, stack traces, and caller information.
// It sets the minimum log level for the logger based on the configured Zap level.
func SetGlobalLogger(config *zap.Config) error {
	// Build a new logger based on the provided Zap configuration.
	loggr, err := config.Build()
	if err != nil {
		return err
	}

	// Create a new SugaredLogger instance with OpenTelemetry options and
	// set it as the global SugaredLogger.
	SugaredLogger = otelzap.New(loggr, otelzap.WithTraceIDField(true),
		otelzap.WithStackTrace(true), otelzap.WithCaller(true),
		otelzap.WithMinLevel(config.Level.Level())).Sugar()

	// Return nil to indicate that the global SugaredLogger was successfully
	// initialized.
	return nil
}
