package enum

// Move is player move
type Move rune

const (
    // NoMove is default value where is no move
    NoMove Move = 0
    // X player move
    X Move = 'X'
    // O player move
    O Move = 'O'
)
