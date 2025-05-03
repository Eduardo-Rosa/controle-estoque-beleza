package infra

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetDB(t *testing.T) {
	db := GetDB()
	assert.NotNil(t, db)
	err := db.Ping()
	assert.Nil(t, err)
}