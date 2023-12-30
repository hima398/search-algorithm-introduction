package maze

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Character struct {
	y, x  int
	score int
}

const Height = 5
const Width = 5
const EndTurn = 10

var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, 1, -1}

type AlternateMazeState struct {
	points     [Height][Width]int
	turn       int
	characters []Character
}

func New(seed int) *AlternateMazeState {
	generator := rand.New(rand.NewSource(int64(seed)))
	res := new(AlternateMazeState)
	res.characters = []Character{{Height / 2, Width/2 - 1, 0}, {Height / 2, Width/2 + 1, 0}}
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			if y == res.characters[0].y && x == res.characters[0].x {
				continue
			}
			if y == res.characters[1].y && x == res.characters[1].x {
				continue
			}
			res.points[y][x] = generator.Intn(10)
		}
	}
	return res
}

// ゲームが終了したかを返す
func (a *AlternateMazeState) IsDone() bool {
	return a.turn == EndTurn
}

// ゲームを1ターン進める
func (a *AlternateMazeState) Advance(action int) {
	if action < 0 {
		fmt.Println("!!!Warning!!!")
		fmt.Println(action)
		fmt.Println(a)
	}
	a.characters[0].x += dx[action]
	a.characters[0].y += dy[action]
	a.characters[0].score += a.points[a.characters[0].y][a.characters[0].x]
	a.points[a.characters[0].y][a.characters[0].x] = 0
	a.turn++
	a.characters[0], a.characters[1] = a.characters[1], a.characters[0]
}

// 現在のプレイヤーが可能な行動をすべて取得する
func (a *AlternateMazeState) LegalActions() []int {
	var res []int
	for k := 0; k < 4; k++ {
		ty, tx := a.characters[0].y+dy[k], a.characters[0].x+dx[k]
		if ty < 0 || ty >= Height || tx < 0 || tx >= Width {
			continue
		}
		res = append(res, k)
	}
	return res
}

type WinningStatus int

const (
	Win WinningStatus = iota
	Lose
	Draw
	None
)

// ゲームの勝敗を取得する
func (a *AlternateMazeState) GetWinningStatus() WinningStatus {
	if a.IsDone() {
		if a.characters[0].score > a.characters[1].score {
			return Win
		} else if a.characters[0].score < a.characters[1].score {
			return Lose
		} else {
			return Draw
		}
	} else {
		return None
	}
}

// プレイヤー視点の盤面評価をする
func (a *AlternateMazeState) GetScore() int {
	return a.characters[0].score - a.characters[1].score
}

// 内容をコピーしたインスタンスを作成する
func (a *AlternateMazeState) Copy() *AlternateMazeState {
	res := new(AlternateMazeState)
	for _, c := range a.characters {
		res.characters = append(res.characters, c)
	}
	for i := range a.points {
		for j := range a.points[i] {
			res.points[i][j] = a.points[i][j]
		}
	}
	res.turn = a.turn
	return res
}

// 先手プレイヤーの勝率計算のためのスコアを計算する
func (a *AlternateMazeState) GetFirstPlayerScoreForWinRate() float64 {
	switch a.GetWinningStatus() {
	case Win:
		if a.isFirstPlayer() {
			return 1
		} else {
			return 0
		}
	case Lose:
		if a.isFirstPlayer() {
			return 0
		} else {
			return 1
		}
	default:
		return 0.5
	}
}

// 現在のゲーム状況を文字列にする
func (a *AlternateMazeState) String() string {
	res := fmt.Sprintf("\nturn:\t%d\n", a.turn)
	for i := range a.characters {
		actualCharacterId := i
		if a.turn%2 == 1 {
			//奇数ターンの場合は初期配置の視点で見るとidが逆
			actualCharacterId = (i + 1) % 2
		}
		c := a.characters[actualCharacterId]
		res += fmt.Sprintf("score(%d)\t%d\t", i, c.score)
		res += fmt.Sprintf("\ty: %d x: %d\n", c.y, c.x)
	}
	for i := range a.points {
		for j := range a.points[i] {
			var isWritten bool
			for id := range a.characters {
				actualCharacterId := id
				if a.turn%2 == 1 {
					actualCharacterId = (id + 1) % 2
				}
				c := a.characters[actualCharacterId]
				if c.y == i && c.x == j {
					if id == 0 {
						res += "A"
					} else {
						res += "B"
					}
					isWritten = true
				}
			}
			if !isWritten {
				if a.points[i][j] > 0 {
					res += strconv.Itoa(a.points[i][j])
				} else {
					res += "."
				}
			}
		}
		res += "\n"
	}
	return res
}

// 現在のプレイヤーが先手であるか判定する
func (a *AlternateMazeState) isFirstPlayer() bool {
	return a.turn%2 == 0
}
