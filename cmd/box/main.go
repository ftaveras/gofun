//Draw a box on the console
package main

import "github.com/ftaveras/gofun/pkg/screen"

func main() {
	screen.Clear()
	t, l := 2, 5
	screen.DrawBox(t, l, 80, 22)
	screen.GotoXY(t+1, l+1)
}
