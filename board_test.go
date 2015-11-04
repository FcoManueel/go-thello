package main
import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestBoardInit(t *testing.T) {
    b := NewBoard(5, 9)
    assert.True(t, len(b) >= 9)
    assert.True(t, len(b[0]) >= 5)
}

func TestBoardMax(t *testing.T) {
    var b Board

    b = NewBoard(0, 0)
    assert.Equal(t, 0, b.MaxX())
    assert.Equal(t, 0, b.MaxY())

    b = NewBoard(1, 2)
    assert.Equal(t, 1, b.MaxX())
    assert.Equal(t, 2, b.MaxY())

    b = NewBoard(5, 6)
    assert.Equal(t, 5, b.MaxX())
    assert.Equal(t, 6, b.MaxY())
}

func TestBoardIsInside(t *testing.T) {
    var b = NewBoard(5, 6)
    assert.True(t, b.IsInside(Move{X: 1, Y: 1}))
    assert.True(t, b.IsInside(Move{X: 3, Y: 3}))
    assert.True(t, b.IsInside(Move{X: 5, Y: 6}))
    assert.False(t, b.IsInside(Move{X: 0, Y: 0}))
    assert.False(t, b.IsInside(Move{X: 0, Y: 1}))
    assert.False(t, b.IsInside(Move{X: 1, Y: 0}))
    assert.False(t, b.IsInside(Move{X: 6, Y: 6}))
    assert.False(t, b.IsInside(Move{X: 5, Y: 7}))
}