package main

type Board [][]int8

func NewBoard(xSize, ySize int) Board {
    // Create room for borders (to simplify logic)
    b := Board{}
    b = make([][]int8, ySize + 2)
    for i := range (b) {
        b[i] = make([]int8, xSize + 2)
    }
    return b
}

func (b Board) MaxX() int {
    if len(b) == 0 {
        return 0
    }else if len(b[0]) == 0 {
        return 0
    }
    return len(b[0]) - 2
}

func (b Board) MaxY() int {
    if len(b) == 0 {
        return 0
    }
    return len(b) - 2
}

func (b *Board) IsInside(m Move) bool {
    return m.X >= 1 && m.X <= b.MaxX() &&
    m.Y >= 1 && m.Y <= b.MaxY()
}
