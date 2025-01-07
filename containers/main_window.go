// Package containers coding=utf-8
// @Project : lucky-draw
// @Time    : 2025/1/7 16:08
// @Author  : chengxiang.luo
// @Email   : chengxiang.luo@foxmail.com
// @File    : main_window.go
// @Software: GoLand
package containers

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func GenMainWindow() fyne.Window {
	a := app.New()
	w := a.NewWindow("Lucky Draw")
	//fileEntry := utils.SelectCSVFile(w)
	menu := SelectCSVFile(w)
	w.SetMainMenu(menu)
	w.Resize(fyne.NewSize(1024, 768)) // 宽度为400像素，高度为300像素
	return w
}
