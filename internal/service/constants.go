package service

import "time"

// Environment is the enumeration for the various runtime environments
type Environment string

// general constants
const (
	// EnvironmentStaging the staging environment
	EnvironmentStaging Environment = "STAGING"
	// EnvironmentProduction the production environment
	EnvironmentProduction Environment = "PRODUCTION"
	// EnvironmentTest the test environment
	EnvironmentTest Environment = "TEST"
	// EnvironmentLocal local development environment
	EnvironmentLocal Environment = "LOCAL"
	// TimeFormat
	// This loses all nanosecond precision
	// TimeFormat = "2006-01-02T15:04:05Z0700"
	TimeFormat = time.RFC3339Nano
	// ErrorResponseKey is the key used when an error is returned from the api
	ErrorResponseKey = "error"
)
