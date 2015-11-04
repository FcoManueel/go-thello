package main
import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestBoardInit(t *testing.T) {
    b := NewPaddedSlice(5, 9)
    assert.True(t, len(b) >= 9)
    assert.True(t, len(b[0]) >= 5)
}

func TestBoardMax(t *testing.T) {
    var b OthelloBoard

    b = NewPaddedSlice(0, 0)
    assert.Equal(t, 0, b.MaxX())
    assert.Equal(t, 0, b.MaxY())

    b = NewPaddedSlice(1, 2)
    assert.Equal(t, 1, b.MaxX())
    assert.Equal(t, 2, b.MaxY())

    b = NewPaddedSlice(5, 6)
    assert.Equal(t, 5, b.MaxX())
    assert.Equal(t, 6, b.MaxY())
}