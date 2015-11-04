package main

const (
    maxPossibleMoves = 64
    blackTile = int8(1)
    whiteTile = int8(-1)
)
type Othello struct {
    board         OthelloBoard
    players       Players
    history       Moves
    currentPlayer int
}

func NewOthello() *Othello {
    o := &Othello{}
    o.board = NewOthelloBoard()
    o.board[4][4], o.board[5][5] = whiteTile, whiteTile
    o.board[4][5], o.board[5][4] = blackTile, blackTile

    o.currentPlayer = 0
    o.players = Players{
        Player{blackTile},
        Player{whiteTile},
    }

    o.history = make(Moves, 0, maxPossibleMoves)
    return o
}

func (o Othello) Board() OthelloBoard {
    return o.board
}
func (o Othello) GetValidMoves() Moves {
    return Moves{}
}
func (o Othello) IsFinished() bool {
    if len(o.history) >= maxPossibleMoves {
        return true
    }
    var whiteTilesAlive = false
    var blackTilesAlive = false
    for j := range(o.board){
        for i := range(o.board[j]) {
            switch (o.board[j][i]) {
            case whiteTile:
                whiteTilesAlive = true
            case blackTile:
                blackTilesAlive = true
            }
            if whiteTilesAlive && blackTilesAlive {
                return false
            }
        }
    }
    return true
}

func (o *Othello) NextPlayer() (playerIndex int) {
    o.currentPlayer = (o.currentPlayer+1)%2
    return o.currentPlayer
}

func (o *Othello) ApplyMove(playerIndex int, move Move) OthelloBoard {
    if playerIndex != o.currentPlayer {
       return o.board
    }
    //TODO add move logic here

    o.NextPlayer()
    return o.board
}