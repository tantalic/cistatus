package main

import (
	"os"
	"time"

	"github.com/pkg/errors"
)

const (
	GITLAB_API_BASE_URL           = "GITLAB_API_BASE_URL"
	GITLAB_API_TOKEN              = "GITLAB_API_TOKEN"
	GITLAB_REFRESH_PERIOD         = "GITLAB_REFRESH_PERIOD"
	GITLAB_REFRESH_PERIOD_DEFAULT = "10s"

	CI_STATUS_HTTP_SERVER_ADDRESS         = "CI_STATUS_HTTP_SERVER_ADDRESS"
	CI_STATUS_HTTP_SERVER_ADDRESS_DEFAULT = ":80"

	CI_STATUS_HTTP_SERVER_JWT_ALGORITHM         = "CI_STATUS_HTTP_SERVER_JWT_ALGORITHM"
	CI_STATUS_HTTP_SERVER_JWT_ALGORITHM_DEFAULT = "HS512"
	CI_STATUS_HTTP_SERVER_JWT_SECRET            = "CI_STATUS_HTTP_SERVER_JWT_SECRET"
)

type config struct {
	GitLabBaseURL         string
	GitLabAPIToken        string
	GitLabRefreshInterval time.Duration

	HTTPAddress  string
	JWTAlgorithm string
	JWTSecret    []byte
}

func configFromEnv() (config, error) {
	var c config

	c.GitLabBaseURL = os.Getenv(GITLAB_API_BASE_URL)
	if c.GitLabBaseURL == "" {
		return c, errors.Errorf("%s environment variable is required", GITLAB_API_BASE_URL)
	}

	c.GitLabAPIToken = os.Getenv(GITLAB_API_TOKEN)
	if c.GitLabAPIToken == "" {
		return c, errors.Errorf("%s environment variable is required", GITLAB_API_TOKEN)
	}

	refreshPeriod := os.Getenv(GITLAB_REFRESH_PERIOD)
	if refreshPeriod == "" {
		refreshPeriod = GITLAB_REFRESH_PERIOD_DEFAULT
	}

	var err error
	c.GitLabRefreshInterval, err = time.ParseDuration(refreshPeriod)
	if err != nil {
		return c, errors.Wrapf(err, "%s environment variable is invalid", GITLAB_REFRESH_PERIOD)
	}

	c.HTTPAddress = os.Getenv(CI_STATUS_HTTP_SERVER_ADDRESS)
	if c.HTTPAddress == "" {
		c.HTTPAddress = CI_STATUS_HTTP_SERVER_ADDRESS_DEFAULT
	}

	c.JWTAlgorithm = os.Getenv(CI_STATUS_HTTP_SERVER_JWT_ALGORITHM)
	if c.JWTAlgorithm == "" {
		c.JWTAlgorithm = CI_STATUS_HTTP_SERVER_JWT_ALGORITHM_DEFAULT
	}

	secret := os.Getenv(CI_STATUS_HTTP_SERVER_JWT_SECRET)
	if secret != "" {
		c.JWTSecret = []byte(secret)
	}

	return c, nil
}
