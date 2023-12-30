package solver

import (
	"math/rand"

	"github.com/hima398/search-algorithm-introduction/5/maze"
)

func RandomAction(rd *rand.Rand, maze *maze.AlternateMazeState) int {
	legalActions := maze.LegalActions()
	return legalActions[rand.Intn(len(legalActions))]
}
