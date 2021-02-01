package enum

// Result is game status result
type Result byte

const (
    // InProgress is in-progress status
    InProgress Result = 0
    // Draw is draw status
    Draw Result = 1
    // WinX is X player wins status
    WinX Result = 2
    // WinO is O player wins status
    WinO Result = 3
)
