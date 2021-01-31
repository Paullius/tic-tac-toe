package api

import (
	"time"

	"github.com/Paullius/tic-tac-toe/game"
)

// LocalCache is simple local cache
type LocalCache map[string]*game.Game
const maxCacheEntries = 50

var cache LocalCache = LocalCache{}

// CleanUp cleans up old instances
func (lc *LocalCache) CleanUp() {
    currentTime := time.Now().UTC()
    for key, c := range *lc {
        // game session can be max 3 hours
        endGameTime := c.StartTime.Add(3*time.Hour)
        if endGameTime.Before(currentTime) {
            delete(*lc, key)
        }
    }

    // remove from cache if above treshold 
    // TODO: change algorithm to use linkedlist to avoid iteration every time
    if len(*lc) > maxCacheEntries {
        var oldestEntry *game.Game
        for _, g := range *lc {
            if oldestEntry == nil {
                oldestEntry = g
            } else if g.StartTime.Before(oldestEntry.StartTime) {
                oldestEntry = g
            }
        }
        delete(*lc,oldestEntry.ID)
    }
}
