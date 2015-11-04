package main
import (
    "testing"
    "github.com/stretchr/testify/assert"
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