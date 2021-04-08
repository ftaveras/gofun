package screen

import "fmt"

func ClearScr() {
	fmt.Printf("\033[J")
}

func GotoXY(x, y int) {
	fmt.Printf("\033[%d;%dH", y, x)
}

func PrintXY(a string, x, y int) {
	fmt.Printf("\033[%d;%dH%s", y, x, a)
}

func DrawBox(w, h int) {
	/*
		Ascii Symbols: https://en.wikipedia.org/wiki/Box-drawing_character
	*/
	//Print Corners
	top := 1
	lft := 1
	ClearScr()
	//Print horizotal bars
	for x := lft; x < w+lft; x++ {
		PrintXY("═", x, top)
		PrintXY("═", x, h+top)
	}

	for x := top; x < h+top; x++ {
		PrintXY("║", lft, x)
		PrintXY("║", w, x)
	}
	PrintXY("╔", lft, top) //tl
	PrintXY("╗", w, top)   //tr
	PrintXY("╚", 0, h+top) //bl
	PrintXY("╝", w, h+top) //br
	GotoXY(lft+1, top+1)   //position
}
