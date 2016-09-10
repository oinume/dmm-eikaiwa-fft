package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTeachersFromIdOrUrl(t *testing.T) {
	a := assert.New(t)
	teachers, err := NewTeachersFromIdsOrUrl("1,2")
	a.NoError(err)
	a.Equal(2, len(teachers))

	teachers2, err := NewTeachersFromIdsOrUrl("1,2,3,")
	a.NoError(err)
	a.Equal(3, len(teachers2))

	teachers3, err := NewTeachersFromIdsOrUrl("")
	a.Error(err)
	a.Equal(0, len(teachers3))
}