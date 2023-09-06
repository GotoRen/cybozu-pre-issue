// Package config configures environment parameters.
package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	FileName  string
	DebugMode bool
}

// Get gets environment variables.
func Get() (*Config, error) {
	var (
		cfg       Config
		missed    []string
		debugMode string
	)

	for _, prop := range []struct {
		field *string
		name  string
	}{
		{&cfg.FileName, "FILE_NAME"},
		{&debugMode, "DEBUG_MODE"},
	} {
		v := os.Getenv(prop.name)
		*prop.field = v

		if v == "" {
			missed = append(missed, prop.name)
		}
	}

	if len(missed) > 0 {
		return nil, fmt.Errorf("missing required environment variables: %v", missed)
	}

	isDebugMode, err := strconv.ParseBool(debugMode)
	if err != nil {
		return nil, fmt.Errorf("parse error debugMode: %w", err)
	}

	cfg.DebugMode = isDebugMode

	return &cfg, nil
}
