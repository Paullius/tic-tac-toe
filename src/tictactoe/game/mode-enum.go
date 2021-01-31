package game

// ModeEnum is for defining game mode
type ModeEnum rune

const (
    // PvP is Player vs Player
    PvP ModeEnum = 0
    // PvAIv1 is for Player vs AI lvl1
    PvAIv1 ModeEnum = 1
    // PvAIv2 is for Player vs AI lvl2
    PvAIv2 ModeEnum = 2
)
