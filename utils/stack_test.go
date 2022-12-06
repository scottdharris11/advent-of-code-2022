package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := Stack{}
	s.Push('A')
	s.Push('B')
	s.Push('C')
	assert.Equal(t, 'C', s.Pop())
	s.Push('D')
	assert.Equal(t, 'D', s.Pop())
	assert.Equal(t, 'B', s.Peek())
	assert.Equal(t, 'B', s.Pop())
	assert.Equal(t, 'A', s.Pop())
	assert.Nil(t, s.Pop())
	assert.Nil(t, s.Peek())
}
