package users

import (
	"github.com/hostport/wildduck-wrapper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUsersNew(t *testing.T) {
	wildduck.SecretKey = "supersecret"
	wildduck.Endpoint = "http://10.0.1.20:8080"
	user := &wildduck.UserParams{Username: "anders@test.com", Password: "test123!"}
	err := New(user)
	assert.Nil(t, err)
	assert.Greater(t, len(user.Id), 0)
}

func TestUsersGetAll(t *testing.T) {
	wildduck.SecretKey = "supersecret"
	wildduck.Endpoint = "http://10.0.1.20:8080"
	res, err := GetAll()
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
