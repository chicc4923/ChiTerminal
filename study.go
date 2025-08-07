package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func studyMain() {
	// 先声明一个框架实例
	myApp := app.New()
	// 创建一个新窗口，title为hello
	myWindow := myApp.NewWindow("Hello")
	// 给窗口创建一个标签组件，内容是hello
	// myWindow.SetContent(widget.NewLabel("hello"))
	clock := widget.NewLabel("")
	updateTime(clock)
	myWindow.SetContent(clock)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	myWindow.Resize(fyne.NewSize(100, 100))

	// 展示这个窗口
	myWindow.Show()

	myWindow2 := myApp.NewWindow("Larger")
	myWindow2.SetContent(widget.NewLabel("New Content"))
	myWindow2.SetContent(widget.NewButton("Open New Button", func() {
		w3 := myApp.NewWindow("Third")
		w3.SetContent(widget.NewLabel("Third Content"))
		w3.Show()
	}))
	myWindow2.Show()

	w4 := myApp.NewWindow("Fourth")
	w4.SetContent(container.NewVBox(makeUI()))
	w4.Show()

	// 启动事件循环
	myApp.Run()
	tidyUp()
}

// 退出时控制台输出
func tidyUp() {
	fmt.Println("Exited")
}

func updateTime(clock *widget.Label) {
	format := time.Now().Format("Time: 03:04:05")
	clock.SetText(format)
}

func makeUI() (*widget.Label, *widget.Entry) {
	out := widget.NewLabel("Hello world!")
	in := widget.NewEntry()

	in.OnChanged = func(content string) {
		out.SetText("Hello " + content + "!")
	}
	return out, in
}
