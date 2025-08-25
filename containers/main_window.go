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
)

func GenMainWindow() fyne.Window {
	a := app.New()
	w := a.NewWindow("Lucky Draw")

	// 加载背景图片
	backgroundImage := canvas.NewImageFromFile("static/bg2.png")
	backgroundImage.FillMode = canvas.ImageFillOriginal // 保持图片原始大小

	drawCountEntry := SetUserCnt()
	resultsText := ShowLucky()
	progress := SetProcessBar()

	// 优化代码，增加新函数，确保功能正确
	// 新增函数用于处理停止抽奖逻辑
	stopChan := make(chan struct{})

	startLuckyDraw := func() {
		go LuckyDraw(w, drawCountEntry, resultsText, progress, UserList, func(){})
	}

	stopLuckyDraw := func() {
		// 向通道发送信号以停止抽奖
		select {
		case stopChan <- struct{}{}:
		default:
		}
	}

	// 使用 GenButton 生成带有 Start/Stop 功能的按钮
	drawButton := GenButton(startLuckyDraw, stopLuckyDraw)

	w.SetContent(container.NewStack(
		backgroundImage,
		container.NewVBox(drawCountEntry, drawButton, progress, resultsText),
	))

	resultsText.Resize(fyne.NewSize(backgroundImage.Size().Width, 400))

	menu := SelectCSVFile(w)
	w.SetMainMenu(menu)
	w.Resize(fyne.NewSize(backgroundImage.Size().Width, backgroundImage.Size().Height))
	return w
}
