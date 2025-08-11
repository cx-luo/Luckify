// Package containers coding=utf-8
// @Project : luckify
// @Time    : 2025/1/7 17:01
// @Author  : chengxiang.luo
// @Email   : chengxiang.luo@foxmail.com
// @File    : main_menu.go
// @Software: GoLand
package containers

import (
	"encoding/csv"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)

type UserInfo struct {
	UserId   string
	UserName string
}

var UserList [][]string

// SelectCSVFile 定义选择文件的函数
func SelectCSVFile(w fyne.Window) *fyne.MainMenu {
	// 创建文件过滤器，只允许选择CSV文件
	csvFilter := storage.NewExtensionFileFilter([]string{".csv"})

	// 创建文件选择对话框的回调函数
	fileSelected := func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		if reader == nil {
			return
		}
		defer reader.Close()

		// 读取文件内容
		csvReader := csv.NewReader(reader)
		records, err := csvReader.ReadAll()
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		for _, record := range records {
			UserList = append(UserList, record)
		}
	}

	// 创建一个菜单项，当点击时会打开文件选择器
	openCSVItem := fyne.NewMenuItem("Open CSV", func() {
		fileopen := dialog.NewFileOpen(fileSelected, w)
		fileopen.SetFilter(csvFilter)
		fileopen.Show()
	})

	// 创建文件菜单并添加菜单项
	fileMenu := fyne.NewMenu("File", openCSVItem)

	// 创建主菜单并添加文件菜单
	mainMenu := fyne.NewMainMenu(
		fileMenu,
	)

	return mainMenu
}
