package main

type OthelloBoard struct {
    content  [][]TileColor
    iterator *Move
}

func NewOthelloBoard() *OthelloBoard {
    var newBoard = make([][]TileColor, 8 + 2)
    for i := range newBoard {
        newBoard[i] = make([]TileColor, 8 + 2)
    }
    return &OthelloBoard{
        content: newBoard,
    }
}

func (b *OthelloBoard) Get(m Move) TileColor {
    return b.content[m.Y][m.X]
}

func (b *OthelloBoard) GetXY(x, y int) TileColor {
    return b.content[y][x]
}

func (b *OthelloBoard) Set(m Move, t TileColor) {
    b.content[m.Y][m.X] = t
}

func (b *OthelloBoard) SetXY(x, y int, t TileColor) {
    b.content[y][x] = t
}

func (b *OthelloBoard) MaxX() int {
    if len(b.content) == 0 {
        return 0
    } else if len(b.content[0]) == 0 {
        return 0
    }
    return len(b.content[0]) - 2
}

func (b *OthelloBoard) MaxY() int {
    if len(b.content) == 0 {
        return 0
    }
    return len(b.content) - 2
}

func (b *OthelloBoard) IsInside(m Move) bool {
    return m.X >= 1 && m.X <= b.MaxX() &&
    m.Y >= 1 && m.Y <= b.MaxY()
}

func (b *OthelloBoard) First() *Move {
    b.iterator = &Move{1, 1}
    return b.iterator
}

func (b *OthelloBoard) Next() *Move {
    switch {
    case b.iterator == nil ||
    b.iterator.X >= b.MaxX() && b.iterator.Y >= b.MaxY():
        b.iterator = nil
    case b.iterator.X == b.MaxX():
        b.iterator.X = 1
        b.iterator.Y += 1
    default:
        b.iterator.X += 1
    }
    return b.iterator
}

type MoveFunction func()(*Move)

func (b *OthelloBoard) MoveLeft() *Move {              return b.safeMove(-1, +0)}
func (b *OthelloBoard) MoveRight() *Move {             return b.safeMove(+1, +0)}
func (b *OthelloBoard) MoveUp() *Move {                return b.safeMove(+0, -1)}
func (b *OthelloBoard) MoveDown() *Move {              return b.safeMove(+0, +1)}
func (b *OthelloBoard) MoveFirstDiagonalUp() *Move {   return b.safeMove(+1, -1)}
func (b *OthelloBoard) MoveFirstDiagonalDown() *Move { return b.safeMove(-1, +1)}
func (b *OthelloBoard) MoveSecondDiagonalUp() *Move {  return b.safeMove(-1, -1)}
func (b *OthelloBoard) MoveSecondDiagonalDown() *Move {return b.safeMove(+1, +1)}

func (b *OthelloBoard) safeMove(xIncrement, yIncrement bool) *Move {
    if b.iterator == nil {
        return nil
    }
    b.iterator.X += xIncrement
    b.iterator.Y += yIncrement
    if !b.IsInside(*b.iterator) {
        b.iterator = nil
    }
    return b.iterator
}