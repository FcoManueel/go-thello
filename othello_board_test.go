package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOthelloBoardInit(t *testing.T) {
	b := NewOthelloBoard()
	assert.True(t, b.MaxX() == 8)
	assert.True(t, b.MaxY() == 8)
}

func TestOthelloBoardIsInside(t *testing.T) {
	var b = NewOthelloBoard()
	assert.True(t, b.IsInside(Move{X: 1, Y: 1}))
	assert.True(t, b.IsInside(Move{X: 4, Y: 4}))
	assert.True(t, b.IsInside(Move{X: 8, Y: 8}))
	assert.False(t, b.IsInside(Move{X: 0, Y: 0}))
	assert.False(t, b.IsInside(Move{X: 0, Y: 1}))
	assert.False(t, b.IsInside(Move{X: 1, Y: 0}))
	assert.False(t, b.IsInside(Move{X: 8, Y: 9}))
	assert.False(t, b.IsInside(Move{X: 9, Y: 8}))
}

func TestOthelloBoardIterators(t *testing.T) {
	var b = NewOthelloBoard()

	first := b.First()
	assert.NotNil(t, first)
	assert.Equal(t, Move{1, 1}, *first)

	second := b.Next()
	assert.NotNil(t, second)
	assert.Equal(t, Move{2, 1}, *second)

	b.First()
	for i := 0; i < 8; i++ {
		b.Next()
	}
	secondRow := b.iterator
	assert.NotNil(t, secondRow)
	assert.Equal(t, Move{1, 2}, *secondRow)
}
