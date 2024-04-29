package main

import (
	"game/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	a := game.Board{}

	screenHeight := 1080
	screenLength := 1920
	a.NewProportion(108, screenHeight, screenLength)
	a.GenerateRandomState()
	rl.InitWindow(int32(screenLength), int32(screenHeight), "Game of life")
	rl.SetTargetFPS(60)
	counter := 0
	kcounter := 0
	pos := 0
	sortColor := []rl.Color{
		rl.Yellow,
		rl.Lime,
		rl.Green,
		rl.SkyBlue,
		rl.Blue,
		rl.DarkBlue,
		rl.DarkGreen,
		rl.Red,
		rl.Orange,
		rl.Pink,
		rl.Purple,
		rl.Violet,
	}
	a.SetColor(sortColor[pos],rl.Black)
	for !rl.WindowShouldClose() {
		if counter == 800000 {
			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)
			a.NextState()
			a.DrawState()
			rl.EndMode2D()
			rl.EndDrawing()
			counter =0
		}
		counter++
		kcounter++
		if kcounter == 300000000 {
			a.GenerateRandomState()
			pos++
			pos = pos % len(sortColor)
			kcounter = 0
			a.SetColor(sortColor[pos],rl.Black)

		}
	}

	rl.CloseWindow()
	
}
