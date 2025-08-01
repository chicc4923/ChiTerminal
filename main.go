package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	mainWindow := myApp.NewWindow("自适应窗口示例")

	// 创建子窗口内容
	subContent := container.NewVBox(
		widget.NewLabel("子窗口内容"),
		widget.NewButton("点击", func() {
			println("按钮被点击")
		}),
	)

	// 创建子窗口背景和边框
	subBg := canvas.NewRectangle(color.RGBA{R: 230, G: 230, B: 230, A: 255})
	subBg.SetMinSize(fyne.NewSize(300, 200))

	// 使用 container.New 替代 NewMax/NewStacked
	subWindow := container.New(
		layout.NewMaxLayout(),
		subBg,
		container.NewPadded(subContent),
	)

	// 主窗口布局
	mainContent := container.New(
		layout.NewBorderLayout(
			widget.NewLabel("主窗口标题"),
			nil, nil, nil,
		),
		widget.NewLabel("主窗口标题"),
		container.NewCenter(subWindow),
	)

	mainWindow.SetContent(mainContent)
	mainWindow.Resize(fyne.NewSize(600, 400))
	mainWindow.ShowAndRun()
}
