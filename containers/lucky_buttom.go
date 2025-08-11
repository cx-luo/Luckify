// Package containers coding=utf-8
// @Project : luckify
// @Time    : 2025/1/8 9:18
// @Author  : chengxiang.luo
// @Email   : chengxiang.luo@foxmail.com
// @File    : lucky_button.go
// @Software: GoLand
package containers

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"strconv"
	"time"
)

func SetUserCnt() *widget.Entry {
	// 创建一个输入框让用户输入要抽取的数量
	drawCountEntry := widget.NewEntry()
	drawCountEntry.SetPlaceHolder("Enter number of draws")
	return drawCountEntry
}

func SetProcessBar() *widget.ProgressBar {
	// 创建一个进度条
	progress := widget.NewProgressBar()
	progress.Min = 0
	progress.Max = 100
	progress.Value = 0
	return progress
}

func ShowLucky() *widget.Entry {
	// 创建一个文本框用于显示结果
	resultsText := widget.NewMultiLineEntry()
	resultsText.Wrapping = fyne.TextWrapWord
	resultsText.MultiLine = true
	resultsText.Disable()
	return resultsText
}

func generateUniqueRandomNumbers(max, count int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, count)
	seen := make(map[int]bool, count)
	for len(seen) < count {
		number := rand.Intn(max)
		if !seen[number] {
			seen[number] = true
			numbers = append(numbers, number)
		}
	}
	return numbers
}

func LuckyDraw(w fyne.Window, drawCountEntry, resultsText *widget.Entry, progress *widget.ProgressBar, prizes [][]string) {
	// 获取用户输入的抽取数量
	drawCount, err := strconv.Atoi(drawCountEntry.Text)
	if err != nil || drawCount <= 0 {
		dialog.ShowError(fmt.Errorf("please enter a valid positive integer"), w)
		return
	}

	// 清空结果文本框
	resultsText.SetText("")

	// 更新进度条并抽取奖项
	for i := 0; i < drawCount; i++ {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(100 * time.Millisecond) // 模拟抽奖过程中的延迟
		progress.Value = float64(i+1) / float64(drawCount) * 100
		selectedPrize := prizes[rand.Intn(len(prizes))]
		resultsText.SetText(resultsText.Text + fmt.Sprintf("%s %s", selectedPrize[0], selectedPrize[1]) + "\n")
	}
}

func GenButton() *widget.Button {
	drawButton := widget.NewButton("Start", nil)
	// 定义一个标志变量，用于控制随机显示的状态
	isRunning := false

	drawButton.OnTapped = func() {
		if isRunning {
			// 如果正在运行，则停止
			isRunning = false
			drawButton.SetText("Start")
		} else {
			// 如果未运行，则开始
			isRunning = true
			drawButton.SetText("Stop")
			// 这里可以添加开始抽奖的逻辑
			// 例如调用LuckyDraw等函数
		}
	}
	return drawButton
}
