package main

import "fmt"

const (
	maxPossibleMoves           = 64
	emptyTile        TileColor = 0
	blackTile        TileColor = 1
	whiteTile        TileColor = -1
)

type Othello struct {
	board         *OthelloBoard
	players       Players
	history       Moves
	currentPlayer int
}

func NewOthello() *Othello {
	o := &Othello{}
	o.board = NewOthelloBoard()
	o.board.SetXY(4, 4, whiteTile)
	o.board.SetXY(5, 5, whiteTile)
	o.board.SetXY(4, 5, blackTile)
	o.board.SetXY(5, 4, blackTile)

	o.currentPlayer = 0
	o.players = Players{
		Player{blackTile},
		Player{whiteTile},
	}

	o.history = make(Moves, 0, maxPossibleMoves)
	return o
}

func (o *Othello) Board() *OthelloBoard {
	return o.board
}

func (o *Othello) GetValidMoves(player int) Moves {
	var validMoves = Moves{}
	for m := o.board.First(); m != nil; m = o.board.Next() {
		fmt.Println(m)
		tileIsOccupied := o.board.Get(*m) != emptyTile
		if tileIsOccupied {
			continue
		}

		opponent := o.players[(player+1)%2]
		neighbors := o.getNeighbors(*m, true, opponent.tileColor)
		if len(neighbors) > 0 {
			validMoves = append(validMoves, *m)
		}
		//        for _ = range neighbors {
		//            if true { //o.canEat(m, neighbors[i]) {
		//                validMoves = append(validMoves, *m)
		//                break
		//            }
	}
	//    for j := 1; j <= o.board.MaxY(); j++ {
	//        for i := 1; i <= o.board.MaxX(); i++ {
	//            if o.board[j][i] != 0 {
	//                continue
	//            }
	//
	//            opponentToken := o.players[(o.currentPlayer+1)%2].tileColor
	//            neighbors := o.getNeighbors(i, j, true, opponentToken)
	//            for i := range neighbors {
	//                if o.canEat(Move{X: i, Y: j}, neighbors[i]) {
	//                    validMoves = append(validMoves, Move{i, j})
	//                    break
	//                }
	//            }
	//        }
	//    }
	return validMoves
}

// getNeighbors return all the surrounding Moves that fall inside the
// board (i.e. they are ignored if they are in the padding area).
// Also, if withToken = true the returned Moves should be filled with token
func (o *Othello) getNeighbors(m Move, withTile bool, tile TileColor) Moves {
	possibleNeighbors := Moves{
		Move{m.X + 1, m.Y}, // horizontal
		Move{m.X - 1, m.Y},
		Move{m.X, m.Y + 1}, // vertical
		Move{m.X, m.Y - 1},
		Move{m.X - 1, m.Y - 1}, // asc diagonal
		Move{m.X + 1, m.Y + 1},
		Move{m.X - 1, m.Y + 1}, // desc diagonal
		Move{m.X + 1, m.Y - 1},
	}
	var actualNeighbors = Moves{}
	for i := range possibleNeighbors {
		n := &possibleNeighbors[i]
		if o.board.IsInside(*n) &&
			(!withTile || o.board.Get(*n) == tile) {
			actualNeighbors = append(actualNeighbors, *n)
		}
	}

	return actualNeighbors
}

func (o *Othello) IsFinished() bool {
	if len(o.history) >= maxPossibleMoves {
		return true
	}
	var whiteTilesAlive = false
	var blackTilesAlive = false
	for m := o.board.First(); m != nil; m = o.board.Next() {
		switch o.board.Get(*m) {
		case whiteTile:
			whiteTilesAlive = true
		case blackTile:
			blackTilesAlive = true
		}
		if whiteTilesAlive && blackTilesAlive {
			return false
		}
	}
	return true
}

func (o *Othello) NextPlayer() (playerIndex int) {
	o.currentPlayer = (o.currentPlayer + 1) % 2
	return o.currentPlayer
}

func (o *Othello) ApplyMove(playerIndex int, move Move) *OthelloBoard {
	if playerIndex != o.currentPlayer {
		return o.board
	}
	// TODO add move logic here
	// if there's a score, here should be affected

	o.history = append(o.history, move)
	o.NextPlayer()
	return o.board
}
