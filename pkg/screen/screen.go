package screen

import "fmt"

func Clear() {
	fmt.Printf("\033[2J")
}

func GotoXY(x, y int) {
	fmt.Printf("\033[%d;%dH", x, y)
}

func PrintXY(a interface{}, x, y int) {
	fmt.Printf("\033[%d;%dH%v", x, y, a)
}

func DrawBox(top, lft, w, h int) {
	/*
		Ascii Symbols: https://en.wikipedia.org/wiki/Box-drawing_character
	*/
	//Print Corners
	//top := 1
	//lft := 1
	//Print horizotal bars
	for y := lft; y < w+lft; y++ {
		PrintXY("═", top, y)
		PrintXY("═", h+top, y)
	}

	for x := top; x < h+top; x++ {
		PrintXY("║", x, lft)
		PrintXY("║", x, w+lft)
	}
	PrintXY("╔", top, lft)     //tl
	PrintXY("╗", top, w+lft)   //tr
	PrintXY("╚", h+top, lft)   //bl
	PrintXY("╝", h+top, w+lft) //br
}
