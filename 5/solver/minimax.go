package solver

import "github.com/hima398/search-algorithm-introduction/5/maze"

func miniMaxScore(state *maze.AlternateMazeState, depth int) int {
	if state.IsDone() || depth == 0 {
		return state.GetScore()
	}
	legalActions := state.LegalActions()
	if len(legalActions) == 0 {
		return state.GetScore()
	}
	bestScore := -inf
	for _, action := range legalActions {
		nextState := state.Copy()
		nextState.Advance(action)
		score := -miniMaxScore(nextState, depth-1)
		if score > bestScore {
			bestScore = score
		}
	}
	return bestScore
}

func MiniMaxAction(state *maze.AlternateMazeState, depth int) int {
	bestAction := -1
	bestScore := -inf
	for _, action := range state.LegalActions() {
		nextState := state.Copy()
		nextState.Advance(action)
		score := -miniMaxScore(nextState, depth)
		if score > bestScore {
			bestAction = action
			bestScore = score
		}
	}
	return bestAction
}
