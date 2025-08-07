package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func studyCanvas() {
	myApp := app.New()
	w1 := myApp.NewWindow("Canvas")
	myCanvas := w1.Canvas()

	// blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	// rect := canvas.NewRectangle(blue)
	// myCanvas.SetContent(rect)

	// go func() {
	// 	time.Sleep(3 * time.Second)
	// 	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	// 	rect.FillColor = green
	// 	rect.Refresh()
	// }()
	setContentToText(myCanvas)
	setContentToCircle(myCanvas)

	w1.Resize(fyne.NewSize(100, 100))
	w1.ShowAndRun()
}

func setContentToText(c fyne.Canvas) {
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text := canvas.NewText("Text", green)
	text.TextStyle.Bold = true
	c.SetContent(text)
}

func setContentToCircle(c fyne.Canvas) {
	red := color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = red
	c.SetContent(circle)
}
