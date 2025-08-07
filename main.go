package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.New()
	w1 := myApp.NewWindow("Container")
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	text1.Move(fyne.NewPos(10, 10))
	text2.Move(fyne.NewPos(10, 30))
	content := container.NewWithoutLayout(text1, text2)

	w1.SetContent(content)
	w1.ShowAndRun()
}
