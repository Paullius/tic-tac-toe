package enum

// Mode is for defining game mode
type Mode rune

const (
    // PvP is Player vs Player
    PvP Mode = 0
    // PvAIv1 is for Player vs AI lvl1
    PvAIv1 Mode = 1
    // PvAIv2 is for Player vs AI lvl2
    PvAIv2 Mode = 2
)
