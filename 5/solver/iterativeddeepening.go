package solver

import (
	"fmt"
	"strings"
	"time"

	"github.com/hima398/search-algorithm-introduction/5/maze"
	"github.com/hima398/search-algorithm-introduction/5/util"
)

func alphaBetaScoreWithTimeThreshold(state *maze.AlternateMazeState, alpha, beta, depth int, timeKeeper *util.TimeKeeper) int {
	if timeKeeper.IsTimeOver() {
		return 0
	}
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
		score := -alphaBetaScoreWithTimeThreshold(nextState, -beta, -alpha, depth-1, timeKeeper)
		if timeKeeper.IsTimeOver() {
			return 0
		}
		if score > alpha {
			alpha = score
		}
		if alpha >= beta {
			return alpha
		}
	}
	return alpha
}

func alphaBetaActionWithTimeThreshold(state *maze.AlternateMazeState, depth int, timeKeeper *util.TimeKeeper) int {
	bestAction := -1
	alpha := -inf
	beta := inf
	var isUpdated bool
	for _, action := range state.LegalActions() {
		nextState := state.Copy()
		nextState.Advance(action)
		score := -alphaBetaScoreWithTimeThreshold(nextState, -beta, -alpha, depth, timeKeeper)
		if timeKeeper.IsTimeOver() {
			return 0
		}
		if score > alpha {
			bestAction = action
			alpha = score
			isUpdated = true
		}
	}
	if !isUpdated {
		fmt.Println(strings.Repeat("!", 40), "Best action is not updated.", strings.Repeat("!", 40))
	}
	return bestAction

}

func IterativeDeepeningAction(state *maze.AlternateMazeState, milliSecond int) int {
	timeKeeper := util.NewTimeKeeper(milliSecond)
	bestAction := -1
	var isUpdated bool
	for d := 1; ; d++ {
		action := alphaBetaActionWithTimeThreshold(state, d, timeKeeper)

		if timeKeeper.IsTimeOver() {
			break
		} else {
			bestAction = action
			isUpdated = true
		}
	}
	if !isUpdated {
		fmt.Println(strings.Repeat("!", 40), "Best action is not updated.", strings.Repeat("!", 40))
		now := time.Now().UnixMicro()
		fmt.Println(now)
		fmt.Println(timeKeeper.StartTime)
		fmt.Println(now - timeKeeper.StartTime)
		fmt.Println(timeKeeper.TimeThreshold)
		fmt.Println(timeKeeper.IsTimeOver())
	}
	return bestAction
}
