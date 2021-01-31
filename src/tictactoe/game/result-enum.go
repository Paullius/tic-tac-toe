package game

// ResultEnum is result enum
type ResultEnum byte

const (
    // InProgress is for in-progress status
    InProgress ResultEnum = 0
    // Draw is for draw status
    Draw ResultEnum = 1
    // WinX is for X win status
    WinX ResultEnum = 2
    // WinO is for O win status
    WinO ResultEnum = 3
)
