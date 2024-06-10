package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type Button struct {
	Text      string
	X         int32
	Y         int32
	Width     int32
	Height    int32
	Color     rl.Color
	TextColor rl.Color
}

func DrawButton(b *Button) {
	rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, b.Color)
	rl.DrawText(b.Text, b.X+b.Width/2-36, b.Y+b.Height/2-4, 14, b.TextColor)
}
