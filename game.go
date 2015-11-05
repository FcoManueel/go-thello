package main

import "fmt"

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

type TileColor int8

type Player struct {
	tileColor TileColor
}

type Players []Player

type PlayerInterface interface {
	SelectMove(Game) Move
}

type Move struct {
	X int
	Y int
}

func (m Move) String() string {
	return fmt.Sprintf("[%d, %d]", m.X, m.Y)
}

type Moves []Move
