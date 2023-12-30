package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/hima398/search-algorithm-introduction/5/maze"
	"github.com/hima398/search-algorithm-introduction/5/solver"
)

var w1, w2 int

func playGame(seed int) {
	rd := rand.New(rand.NewSource(int64(seed)))
	state := maze.New(rd)
	for !state.IsDone() {
		// 1P
		{
			fmt.Println("1p " + strings.Repeat("-", 40))

			//action := miniMaxAction(state, maze.EndTurn)
			action := solver.AlphaBetaAction(state, maze.EndTurn)
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

			//action := randomAction(rd, state)
			action := solver.MiniMaxAction(state, maze.EndTurn)
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

func main() {
	for seed := 0; seed < int(1e5); seed++ {
		playGame(seed)
	}
	fmt.Println(float64(w1*100) / 1e5)
}
