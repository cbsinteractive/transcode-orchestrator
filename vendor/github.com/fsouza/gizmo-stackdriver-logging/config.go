// Copyright 2018 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package logging provides the configuration object for a
// StackDriver-integrated logger.
//
// It supports loading the configuration values using envconfig.
package logging

import (
	"os"

	logrus_stack "github.com/Gurpartap/logrus-stack"
	"github.com/knq/sdhook"
	logrus_env "github.com/marzagao/logrus-env"
	"github.com/sirupsen/logrus"
)

// Config contains configuration for logging level and services integration.
type Config struct {
	Level string `envconfig:"LOGGING_LEVEL" default:"info"`

	// List of environment variables that should be included in all log
	// lines.
	EnvironmentVariables []string `envconfig:"LOGGING_ENVIRONMENT_VARIABLES"`

	// Send logs to StackDriver?
	SendToStackDriver bool `envconfig:"LOGGING_SEND_TO_STACKDRIVER"`

	// StackDriver error reporting options. When present, error logs are
	// going to be reported as errors on StackDriver.
	StackDriverErrorServiceName string `envconfig:"LOGGING_STACKDRIVER_ERROR_SERVICE_NAME"`
	StackDriverErrorLogName     string `envconfig:"LOGGING_STACKDRIVER_ERROR_LOG_NAME" default:"error_log"`

	// When StackDriverCredentialsFile is set, the logger will use the
	// Google logging API to send the logs. Otherwise the fluentd Agent is
	// used.
	StackDriverCredentialsFile string `envconfig:"LOGGING_STACKDRIVER_CREDENTIALS_FILE"`
}

// Logger returns a logrus logger with the features defined in the config.
func (c *Config) Logger() (*logrus.Logger, error) {
	level, err := logrus.ParseLevel(c.Level)
	if err != nil {
		return nil, err
	}

	logger := logrus.New()
	logger.Out = os.Stdout
	logger.Level = level
	logger.Hooks.Add(logrus_stack.StandardHook())
	logger.Hooks.Add(logrus_env.NewHook(c.EnvironmentVariables))

	if c.SendToStackDriver {
		var opts []sdhook.Option
		if c.StackDriverCredentialsFile != "" {
			opts = []sdhook.Option{sdhook.GoogleServiceAccountCredentialsFile(c.StackDriverCredentialsFile)}
		} else {
			opts = []sdhook.Option{sdhook.GoogleLoggingAgent()}
		}

		if c.StackDriverErrorServiceName != "" {
			opts = append(opts,
				sdhook.ErrorReportingService(c.StackDriverErrorServiceName),
				sdhook.ErrorReportingLogName(c.StackDriverErrorLogName),
			)
		}

		gcpLoggingHook, err := sdhook.New(opts...)
		if err != nil {
			return nil, err
		}

		logger.Hooks.Add(gcpLoggingHook)
	}

	return logger, nil
}
