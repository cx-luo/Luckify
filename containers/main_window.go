// Package containers coding=utf-8
// @Project : luckify
// @Time    : 2025/1/7 16:08
// @Author  : chengxiang.luo
// @Email   : chengxiang.luo@foxmail.com
// @File    : main_window.go
// @Software: GoLand
package containers

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GenMainWindow() fyne.Window {
	a := app.New()
	w := a.NewWindow("Lucky Draw")

	// 加载背景图片
	backgroundImage := canvas.NewImageFromFile("../static/bg2.png")
	backgroundImage.FillMode = canvas.ImageFillOriginal // 保持图片原始大小

	drawCountEntry := SetUserCnt()
	//resultsText := ShowLucky()
	t := ShowLucky()
	progress := SetProcessBar()
	drawButton := widget.NewButton("Start", nil)
	drawButton.OnTapped = func() {

	}
	drawButton = widget.NewButton("Draw", func() {
		LuckyDraw(w, drawCountEntry, t, progress, UserList)
	})

	w.SetContent(container.NewStack(
		backgroundImage,
		container.NewVBox(drawCountEntry, drawButton, progress, t),
	))

	t.Resize(fyne.NewSize(backgroundImage.Size().Width, 200))

	menu := SelectCSVFile(w)
	w.SetMainMenu(menu)
	w.Resize(fyne.NewSize(backgroundImage.Size().Width, backgroundImage.Size().Height)) // 宽度为400像素，高度为300像素
	return w
}
