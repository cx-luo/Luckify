// Package lucky_draw coding=utf-8
// @Project : lucky-draw
// @Time    : 2025/1/7 14:55
// @Author  : chengxiang.luo
// @Email   : chengxiang.luo@foxmail.com
// @File    : main.go.go
// @Software: GoLand
package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"lucky-draw/containers"
)

func main() {
	w := containers.GenMainWindow()

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
