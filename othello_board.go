package main

type OthelloBoard [][]int8

func NewOthelloBoard() OthelloBoard {
    return OthelloBoard(NewPaddedSlice(8, 8))
}

func (b OthelloBoard) MaxX() int {
    if len(b) == 0 {
        return 0
    }else if len(b[0]) == 0 {
        return 0
    }
    return len(b[0]) - 2
}

func (b OthelloBoard) MaxY() int {
    if len(b) == 0 {
        return 0
    }
    return len(b) - 2
}

func (b *OthelloBoard) IsInside(m Move) bool {
    return m.X >= 1 && m.X <= b.MaxX() &&
    m.Y >= 1 && m.Y <= b.MaxY()
}
