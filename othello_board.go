package main

type OthelloBoard struct {
	content  [][]TileColor
	iterator *Move
}

func NewOthelloBoard() *OthelloBoard {
	var newBoard = make([][]TileColor, 8+2)
	for i := range newBoard {
		newBoard[i] = make([]TileColor, 8+2)
	}
	return &OthelloBoard{
		content: newBoard,
	}
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
	if b.iterator == nil {
		return nil
	}
	var m = b.iterator
	if m.X >= b.MaxX() && m.Y >= b.MaxY() {
		b.iterator = nil
	} else if m.X == b.MaxX() {
		b.iterator = &Move{1, m.Y + 1}
	} else {
		b.iterator = &Move{m.X + 1, m.Y}
	}
	return b.iterator
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
