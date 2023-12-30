package main

import (
	"fmt"
	"strings"

	"github.com/hima398/search-algorithm-introduction/5/maze"
	"github.com/hima398/search-algorithm-introduction/5/solver"
)

var w1, w2 int

func playGame(seed int) {
	//rd := rand.New(rand.NewSource(int64(seed)))
	state := maze.New(seed)
	for !state.IsDone() {
		// 1P
		{
			fmt.Println("1p " + strings.Repeat("-", 40))

			//action := solver.AlphaBetaAction(state, maze.EndTurn)
			action := solver.IterativeDeepeningAction(state, 100)
			fmt.Printf("action: %d\n", action)
			state.Advance(action)
			fmt.Printf("state: %v\n", state)
			if state.IsDone() {
				//結果を出力してループを抜ける
				switch state.GetWinningStatus() {
				case maze.Win:
					fmt.Println("winner: 2p")
					w2++
				case maze.Lose:
					fmt.Println("winner: 1p")
					w1++
				default:
					fmt.Println("Draw")
				}
				break
			}
		}

		// 2P
		{
			fmt.Println("2p " + strings.Repeat("-", 40))

			//action := solver.MiniMaxAction(state, maze.EndTurn)
			action := solver.IterativeDeepeningAction(state, 1)

			fmt.Printf("action: %d\n", action)
			state.Advance(action)
			fmt.Printf("state: %v\n", state)
			if state.IsDone() {
				//結果を出力してループを抜ける
				switch state.GetWinningStatus() {
				case maze.Win:
					fmt.Println("winner: 1p")
					w1++
				case maze.Lose:
					fmt.Println("winner: 2p")
					w2++
				default:
					fmt.Println("Draw")
				}
				break
			}
		}
	}
}

func testFirstPlayerWinRate(gameNumber int) {
	var firstPlayerWinRate float64
	for i := 0; i < gameNumber; i++ {
		baseState := maze.New(i)
		for j := 0; j < 2; j++ {
			state := baseState.Copy()
			for {
				action1 := solver.IterativeDeepeningAction(state, 200)
				state.Advance(action1)
				if state.IsDone() {
					break
				}
				action2 := solver.IterativeDeepeningAction(state, 5)
				state.Advance(action2)
				if state.IsDone() {
					break
				}
			}
			winRatePoint := state.GetFirstPlayerScoreForWinRate()
			if j == 1 {
				winRatePoint = 1 - winRatePoint
			}
			if winRatePoint >= 0 {
				fmt.Println(state)
			}
			firstPlayerWinRate += winRatePoint
		}
		fmt.Printf("i %d w %f\n", i, firstPlayerWinRate/(float64(i+1)*2.0))
	}
}

func main() {
	testFirstPlayerWinRate(100)
}
