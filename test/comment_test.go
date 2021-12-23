package test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetComments(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/comments")

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestPostComment(t *testing.T) {
	client := resty.New()

	resp, err := client.R().SetBody(`{"slug": "/hello", "author": "eddie", "body": "loveyou"}`).Post(BASE_URL + "/api/comments")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}
