package main

import "fmt"

const debugMode = false
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
		var savedIterator = o.board.iterator
		if o.PlayerCanCapture(player, *m) {
			validMoves = append(validMoves, *m)
		}
		o.board.iterator = savedIterator
	}
	return validMoves
}

func (o *Othello) PlayerCanCapture(playerIndex int, m Move) bool {
	var tileIsOccupied = o.board.Get(m) != emptyTile
	if tileIsOccupied {
		return false
	}

	player := o.players[playerIndex]
	opponent := o.players[(playerIndex+1)%2]
	neighbors := o.getNeighbors(m, true, opponent.tileColor)

	var move MoveFunction
	var err error

	for _, n := range neighbors {
		fmt.Printf("Checking eatability from %s to %s\n", m.String(), n.String())
		move, err = o.board.GetMoveFunction(m, n)
		if err != nil {
			logDebug("Found an error %+v", err.Error())
			continue
		}

		var playerCanCapture = true
		var tile *Move
		// A player can capture there are opponent tiles in line, followed one of his tiles
		for tile = move(&m); tile != nil; tile = move(tile) {
			if o.board.Get(*tile) != opponent.tileColor {
				logDebug("Found a tile of different color in %s\n\n", tile.String())
				playerCanCapture = o.board.Get(*tile) == player.tileColor
				break
			}
		}
		if tile != nil && playerCanCapture {
			return true
		}
	}
	return false
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
	for _, n := range possibleNeighbors {
		if o.board.IsInside(n) &&
			(!withTile || o.board.Get(n) == tile) {
			actualNeighbors = append(actualNeighbors, n)
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

func logDebug(fmtString string, args ...interface{}) {
	if debugMode {
		fmt.Printf(fmtString, args)
	}
}
