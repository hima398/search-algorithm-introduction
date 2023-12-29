package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/hima398/search-algorithm-introduction/5/maze"
)

func randomAction(rd *rand.Rand, maze *maze.AlternateMazeState) int {
	legalActions := maze.LegalActions()
	return legalActions[rand.Intn(len(legalActions))]
}

func playGame(seed int) {
	rd := rand.New(rand.NewSource(int64(seed)))
	state := maze.New(rd)
	for !state.IsDone() {
		// 1P
		{
			fmt.Println("1p " + strings.Repeat("-", 28))

			action := randomAction(rd, state)
			fmt.Printf("action: %d\n", action)
			state.Advance(action)
			fmt.Printf("state: %v\n", state)
			if state.IsDone() {
				//結果を出力してループを抜ける
				switch state.GetWinningStatus() {
				case maze.Win:
					fmt.Println("winner: 2p")
				case maze.Lose:
					fmt.Println("winner: 1p")
				default:
					fmt.Println("Draw")
				}
				break
			}
		}

		// 2P
		{
			fmt.Println("2p " + strings.Repeat("-", 28))

			action := randomAction(rd, state)
			fmt.Printf("action: %d\n", action)
			state.Advance(action)
			fmt.Printf("state: %v\n", state)
			if state.IsDone() {
				//結果を出力してループを抜ける
				switch state.GetWinningStatus() {
				case maze.Win:
					fmt.Println("winner: 1p")
				case maze.Lose:
					fmt.Println("winner: 2p")
				default:
					fmt.Println("Draw")
				}
				break
			}
		}
	}
}

func main() {
	playGame(1)
}
