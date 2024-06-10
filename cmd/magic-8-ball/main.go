package main

import (
	"github.com/dlewis89/magic-8-ball/pkg/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	type Window struct {
		WIDTH  int32
		HEIGHT int32
		FPS    int32
	}

	window := &Window{
		WIDTH:  800,
		HEIGHT: 450,
		FPS:    60,
	}

	rl.InitWindow(window.WIDTH, window.HEIGHT, "Magic 8 Ball")
	defer rl.CloseWindow()

	rl.SetTargetFPS(window.FPS)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		getAnswerBtn := ui.Button{
			Text:      "Get Answer",
			X:         window.WIDTH/2 - 64,
			Y:         window.HEIGHT - 64,
			Width:     128,
			Height:    32,
			Color:     rl.Blue,
			TextColor: rl.White,
		}

		ui.DrawButton(&getAnswerBtn)

		rl.EndDrawing()
	}
}
