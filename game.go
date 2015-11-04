package main

type Game interface {
    Board() Board
    GetValidMoves() Moves
    IsFinished() bool
    NextPlayer() (playerIndex int)
    ApplyMove(playerIndex int, move Move) Board
}

type Board interface {
    MaxX() int
    MaxY() int
    IsInside(Move) bool
}

type Player struct{}

type Players []Player

type PlayerInterface interface {
    SelectMove(Game) Move
}

type Move struct {
    X int
    Y int
}

type Moves []Move

// NewPaddedSlice is a helper function used to create a board with room for borders (simplifies logic)
func NewPaddedSlice(xSize, ySize int) [][]int8 {
    var b [][]int8
    b = make([][]int8, ySize + 2)
    for i := range (b) {
        b[i] = make([]int8, xSize + 2)
    }
    return b
}