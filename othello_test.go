package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestOthelloInit(t *testing.T) {
    o := NewOthello()

    var w = whiteTile
    var b = blackTile
    var p = int8(0) // board padding

    var expectedBoard OthelloBoard

    expectedBoard = OthelloBoard{
        {p, p, p, p, p, p, p, p, p, p},
        {p, 0, 0, 0, 0, 0, 0, 0, 0, p},
        {p, 0, 0, 0, 0, 0, 0, 0, 0, p},
        {p, 0, 0, 0, 0, 0, 0, 0, 0, p},
        {p, 0, 0, 0, w, b, 0, 0, 0, p},
        {p, 0, 0, 0, b, w, 0, 0, 0, p},
        {p, 0, 0, 0, 0, 0, 0, 0, 0, p},
        {p, 0, 0, 0, 0, 0, 0, 0, 0, p},
        {p, 0, 0, 0, 0, 0, 0, 0, 0, p},
        {p, p, p, p, p, p, p, p, p, p},
    }
    assert.EqualValues(t, expectedBoard, o.board, "New board layout is wrong")
}
func TestOthelloBoard(t *testing.T) {}
func TestOthelloGetValidMoves(t *testing.T) {
    o := NewOthello()
    moves := o.GetValidMoves()
    assert.Len(t, moves, 4)
    assert.Contains(t, moves, Move{X: 3, Y: 4})
    assert.Contains(t, moves, Move{X: 4, Y: 3})
    assert.Contains(t, moves, Move{X: 5, Y: 6})
    assert.Contains(t, moves, Move{X: 6, Y: 5})
}

func TestOthelloIsFinished(t *testing.T) {
    o := NewOthello()
    assert.False(t, o.IsFinished(), "A new othello game is not finished")

    o.history = make(Moves, maxPossibleMoves)
    assert.True(t, o.IsFinished(), "Game finishes if game history reaches maxPossibleMoves")

    o = NewOthello()
    o.board[4][4], o.board[4][5], o.board[5][4], o.board[5][5] = blackTile, blackTile, blackTile, blackTile
    assert.True(t, o.IsFinished(), "Game finishes if all tiles in the board are black")

    o.board[4][4], o.board[4][5], o.board[5][4], o.board[5][5] = whiteTile, whiteTile, whiteTile, whiteTile
    assert.True(t, o.IsFinished(), "Game finishes if all tiles in the board are white")
}

func TestOthelloNextPlayer(t *testing.T) {
    o := NewOthello()
    assert.Equal(t, 0, o.currentPlayer)
    o.NextPlayer()
    assert.Equal(t, 1, o.currentPlayer)
    o.NextPlayer()
    assert.Equal(t, 0, o.currentPlayer)
}
func xTestOthelloApplyMove(t *testing.T) {
    var w = whiteTile
    var b = blackTile
    var p = int8(0) // board padding
    o := NewOthello()
    o.ApplyMove(0, Move{X:3, Y:4})

    var expectedBoard1 = OthelloBoard{
        {p, p, p, p, p, p, p, p, p, p},
        {p, 0, 0, 0, 0, 0, 0, 0, 0, p},
        {p, 0, 0, 0, 0, 0, 0, 0, 0, p},
        {p, 0, 0, 0, 0, 0, 0, 0, 0, p},
        {p, 0, 0, b, b, b, 0, 0, 0, p},
        {p, 0, 0, 0, b, w, 0, 0, 0, p},
        {p, 0, 0, 0, 0, 0, 0, 0, 0, p},
        {p, 0, 0, 0, 0, 0, 0, 0, 0, p},
        {p, 0, 0, 0, 0, 0, 0, 0, 0, p},
        {p, p, p, p, p, p, p, p, p, p},
    }
    assert.EqualValues(t, expectedBoard1, o.board)

    o.ApplyMove(0, Move{X:2, Y:3})
    assert.EqualValues(t, expectedBoard1, o.board, "Board should only be affected once in a row per player")

    var expectedBoard2 = OthelloBoard{
        {p, p, p, p, p, p, p, p, p, p},
        {p, 0, 0, 0, 0, 0, 0, 0, 0, p},
        {p, 0, 0, 0, 0, 0, 0, 0, 0, p},
        {p, 0, 0, w, 0, 0, 0, 0, 0, p},
        {p, 0, 0, b, w, b, 0, 0, 0, p},
        {p, 0, 0, 0, b, w, 0, 0, 0, p},
        {p, 0, 0, 0, 0, 0, 0, 0, 0, p},
        {p, 0, 0, 0, 0, 0, 0, 0, 0, p},
        {p, 0, 0, 0, 0, 0, 0, 0, 0, p},
        {p, p, p, p, p, p, p, p, p, p},
    }
    o.ApplyMove(1, Move{X:2, Y:3})
    assert.EqualValues(t, expectedBoard2, o.board)
}