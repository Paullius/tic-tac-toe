package enum

// Result is game status result
type Result byte

const (
    // InProgress is for in-progress status
    InProgress Result = 0
    // Draw is for draw status
    Draw Result = 1
    // WinX is for X win status
    WinX Result = 2
    // WinO is for O win status
    WinO Result = 3
)
