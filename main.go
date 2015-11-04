package main

type Player struct{}
type Players []Player

type Move struct {
    X int
    Y int
}
type Moves []Move

type Game struct {
    board   Board
    players Players
    history Moves
}

type Playable interface {
    Init()
    Board()
    GetValidMoves() Moves
    IsFinished() bool
    NextPlayer() (playerIndex int)
    ApplyMove(playerIndex int, move Move) Board
}

type PlayerInterface interface {
    SelectMove(Game) Move
}

func (g *Game) Init() {}
func (g *Game) Board() {}
func (g *Game) GetValidMoves() Moves {
    return Moves{}
}
func (g *Game) IsFinished() bool {
    return false
}

func (g *Game) NextPlayer() (playerIndex int) {
    return 0
}

func (g *Game) ApplyMove(playerIndex int, move Move) Board {
    return Board{}
}

func main() {

}