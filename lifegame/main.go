package main

import "fmt"

const gameSize = 32

func gamePos(p int) int {
	p = p % gameSize
	if p < 0 {
		p = gameSize + p
	}
	return p
}

func main() {

	var gameState [gameSize][gameSize]int
	var newState [gameSize][gameSize]int
	//var x, y int
	var s int
	//Init
	//gameState[20][20] = 1
	gameState[21][21] = 1
	gameState[22][22] = 1
	gameState[22][23] = 1
	gameState[21][23] = 1
	gameState[20][23] = 1

	//println(gamePos(-1))
	//return

	for {
		newState = gameState

		for x := 0; x < gameSize; x++ {

			for y := 0; y < gameSize; y++ {

				//gameState[x][y] = x * y
				//fmt.Printf("Game: %v, x: %d, y: %d\n", gameState[x][y], x, y)
				s = gameState[gamePos(x-1)][gamePos(y-1)] + gameState[x][gamePos(y-1)] + gameState[gamePos(x+1)][gamePos(y-1)] +
					gameState[gamePos(x-1)][y] + (0) + gameState[gamePos(x+1)][y] +
					gameState[gamePos(x-1)][gamePos(y+1)] + gameState[x][gamePos(y+1)] + gameState[gamePos(x+1)][gamePos(y+1)]

				//Una célula muerta con exactamente 3 células vecinas vivas "nace" (es decir, al turno siguiente estará viva).
				if gameState[x][y] == 0 && s == 3 {
					newState[x][y] = 1
				}
				//Una célula viva con 2 o 3 células vecinas vivas sigue viva, en otro caso muere (por "soledad" o "superpoblación").
				if gameState[x][y] == 1 && (s < 0 || s > 3) {
					newState[x][y] = 0
				}

				/*
					if newState[x][y] == 1 {
						fmt.Printf("%v", " ")

					} else {
						fmt.Printf("%v", "█")
					}
				*/

			}
			//fmt.Printf("\n")
		}
		gameState = newState
		//fmt.Printf("%v\n", gameState)
		for x := 0; x < gameSize; x++ {

			for y := 0; y < gameSize; y++ {

				//gameState[x][y] = x * y
				//fmt.Printf("Game: %v, x: %d, y: %d\n", gameState[x][y], x, y)
				if gameState[x][y] == 1 {
					fmt.Printf("%v", " ")

				} else {
					fmt.Printf("%v", "█")
				}

			}
			fmt.Printf("\n")
		}

	}

}
