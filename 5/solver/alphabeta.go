package solver

import "github.com/hima398/search-algorithm-introduction/5/maze"

func alphaBetaScore(state *maze.AlternateMazeState, alpha, beta, depth int) int {
	if state.IsDone() || depth == 0 {
		return state.GetScore()
	}
	legalActions := state.LegalActions()
	if len(legalActions) == 0 {
		return state.GetScore()
	}
	for _, action := range legalActions {
		nextState := state.Copy()
		nextState.Advance(action)
		score := -alphaBetaScore(nextState, -beta, -alpha, depth-1)
		if score > alpha {
			alpha = score
		}
		if alpha >= beta {
			return alpha
		}
	}
	return alpha
}

func AlphaBetaAction(state *maze.AlternateMazeState, depth int) int {
	bestAction := -1
	alpha := -inf
	beta := inf
	for _, action := range state.LegalActions() {
		nextState := state.Copy()
		nextState.Advance(action)
		score := -alphaBetaScore(nextState, -beta, -alpha, depth)
		if score > alpha {
			bestAction = action
			alpha = score
		}
	}
	return bestAction

}
