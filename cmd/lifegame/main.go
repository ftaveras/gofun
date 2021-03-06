package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/ftaveras/gofun/pkg/screen"
)

const (
	gameCols int = 68
	gameRows int = 21 //gameCols

)

var (
	gameInit = map[int][][]int{
		0: { //glider
			{1, 25},
			{2, 23}, {2, 25},
			{3, 13}, {3, 14}, {3, 21}, {3, 22}, {3, 35}, {3, 35}, {3, 36},
			{4, 12}, {4, 16}, {4, 21}, {4, 22}, {4, 35}, {4, 36},
			{5, 1}, {5, 2}, {5, 11}, {5, 17}, {5, 21}, {5, 22},
			{6, 1}, {6, 2}, {6, 11}, {6, 15}, {6, 17}, {6, 18}, {6, 23}, {6, 25},
			{7, 11}, {7, 17}, {7, 25},
			{8, 12}, {8, 16},
			{9, 13}, {9, 14},
		},
		1: {
			{1, 3},
			{2, 2}, {2, 3}, {2, 4},
		},
		2: {
			{1, 1}, {1, 2}, {1, 3},
			{2, 1},
			{3, 2},
		},
		3: {
			{1, 2}, {1, 5},
			{2, 1},
			{3, 1}, {3, 5},
			{4, 1}, {4, 2}, {4, 3}, {4, 4},
		},
		4: {
			{1, 2}, {1, 3},
			{2, 1}, {2, 2},
			{3, 2},
		},
		/*5: {
			{2, 0}, {2, 2},
		},*/
		6: { //R-Pentomino
			{1, 2}, {1, 3},
			{2, 1}, {2, 2},
			{3, 2},
		},
		/*10: {
			{1, 2},
			{2, 1},
		},*/
		/*17: {
			{1, 1},
			{2, 2},
		},*/
		/*18: {
			{1, 1},
			{2, 1},
		},*/
		/*19: { //Static
			{1, 1},
			{2, 1}, {2, 2},
		},*/
		/*30: { //static
			{1, 1}, {1, 2},
			{2, 0}, {2, 1},
		},*/
		/*40: {
			{1, 0}, {1, 2},
		},*/
		/*42: {
			{1, 0}, {1, 2},
			{2, 1},
		},*/
		58: {
			{1, 0}, {1, 1}, {1, 2},
			{2, 1},
		},
		/*61: {
			{1, 0}, {1, 1}, {1, 2},
			{2, 0}, {2, 2},
		},*/
		/*70: {
			{0, 2},
			{2, 0}, {2, 1},
		},*/
		/*84: {
			{0, 2},
			{1, 1},
			{2, 0},
		},*/
		/*85: {	//Static
			{0, 2},
			{1, 1},
			{2, 0}, {2, 2},
		},*/
		87: {
			{0, 2},
			{1, 1},
			{2, 0}, {2, 1}, {2, 2},
		},
		/*171: {	//Static
			{0, 1},
			{1, 0}, {1, 2},
			{2, 1}, {2, 2},
		},*/
	}
)

func gameValue(game [][]int, x, y int) int {
	r, c := gamePos(x, y)
	return game[r][c]
}

func gamePos(x, y int) (int, int) {
	return gameMod(x, gameRows), gameMod(y, gameCols)
}

func gameMod(i, g int) int {
	v := i % g
	if v < 0 {
		v = v + g
	}
	return v
}

func game(id int, state [][]int) {
	var gameState [][]int
	var copyState [][]int

	gameState = make([][]int, gameRows)
	copyState = make([][]int, gameRows)
	for i := 0; i < gameRows; i++ {
		gameState[i] = make([]int, gameCols)
		copyState[i] = make([]int, gameCols)
	}

	//Init
	//iv := gameInit[state]
	for _, x := range state {
		gameState[x[0]][x[1]] = 1
	}

	screen.Clear()
	screen.DrawBox(1, 1, gameCols+2, gameRows+1)  //game box
	screen.DrawBox(1, gameCols+4, 30, gameRows+1) //info box
	screen.PrintXY(fmt.Sprintf("Game: %d", id), 2, gameCols+6)
	screen.PrintXY("Steps: ", 3, gameCols+6)
	screen.PrintXY("Timer: ", 4, gameCols+6)
	screen.GotoXY(gameRows+3, 1) //move cursor bellow boxes

	//stateDraw(gameState)
	//return
	step := 0
	init := time.Now()
	var s int
	for {

		//copy current state
		for i := 0; i < gameRows; i++ {
			copy(copyState[i], gameState[i]) //we copy current state before we alter it
		}

		for x := 0; x < gameRows; x++ {
			for y := 0; y < gameCols; y++ {

				//calculate adjacent cells state
				s = gameValue(copyState, x-1, y-1) + gameValue(copyState, x-1, y) + gameValue(copyState, x-1, y+1) +
					gameValue(copyState, x+0, y-1) + /*gameValue(copyState, x, y)+ */ gameValue(copyState, x+0, y+1) +
					gameValue(copyState, x+1, y-1) + gameValue(copyState, x+1, y) + gameValue(copyState, x+1, y+1)

				//Una c??lula muerta con exactamente 3 c??lulas vecinas vivas "nace" (es decir, al turno siguiente estar?? viva).
				if copyState[x][y] == 0 && s == 3 {
					gameState[x][y] = 1
					continue
				}
				//Una c??lula viva con 2 o 3 c??lulas vecinas vivas sigue viva, en otro caso muere (por "soledad" o "superpoblaci??n").
				if copyState[x][y] == 1 && (s < 2 || s > 3) {
					gameState[x][y] = 0
				}
			}

		}

		//PrintState
		stateDraw(gameState)
		time.Sleep(time.Millisecond * 120)
		step = step + 1
		screen.PrintXY(step, 3, gameCols+14)
		screen.PrintXY(time.Since(init), 4, gameCols+14)

	}
}

func stateDraw(game [][]int) {
	for x := 0; x < gameRows; x++ {
		for y := 0; y < gameCols; y++ {
			if game[x][y] == 1 {
				screen.PrintXY("???", x+2, y+2)

			} else {
				screen.PrintXY(" ", x+2, y+2)
			}
		}
	}
}

func main() {

	//check user init state
	ui := os.Args[1:]
	if len(ui) > 0 {

		g, err := strconv.ParseInt(ui[0], 10, 32)
		if err != nil {
			panic("Invalid game code")
		}
		game(int(g), gameInit[int(g)])
		return
	}

	i := 0
	r := rand.Intn(len(gameInit)) //We play random game
	for g, v := range gameInit {
		if i == r {
			game(g, v)
			break
		}
		i++
	}
}
