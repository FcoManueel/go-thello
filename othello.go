package main

type Othello struct {
    board   OthelloBoard
    players Players
    history Moves
}

func NewOthello() Othello {
    o := Othello{}
    o.board = NewOthelloBoard()
    return o
}

func (o *Othello) Board() OthelloBoard {
    return o.board
}
func (o *Othello) GetValidMoves() Moves {
    return Moves{}
}
func (o *Othello) IsFinished() bool {
    return false
}

func (o *Othello) NextPlayer() (playerIndex int) {
    return 0
}

func (o *Othello) ApplyMove(playerIndex int, move Move) OthelloBoard {
    return OthelloBoard{}
}