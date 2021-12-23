package test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	log.Info("Health Check!")

	client := resty.New()

	resp, err := client.R().Get(BASE_URL + "/api/health")

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}
