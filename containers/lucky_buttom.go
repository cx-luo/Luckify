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
	"fyne.io/fyne/v2/theme"
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
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
	// 设置字体颜色为黑色
	resultsText.TextStyle = fyne.TextStyle{Monospace: false, Bold: false, Italic: false}
	resultsText.SetMinRowsVisible(12)
	return resultsText
}

func generateUniqueRandomNumbers(max, count int) []int {
	if count > max || max <= 0 || count <= 0 {
		return nil
	}
	indices := rand.Perm(max)
	return indices[:count]
}

func LuckyDraw(w fyne.Window, drawCountEntry, resultsText *widget.Entry, progress *widget.ProgressBar, prizes [][]string, onFinish func()) {
	// 获取用户输入的抽取数量
	drawCount, err := strconv.Atoi(drawCountEntry.Text)
	if err != nil || drawCount <= 0 {
		dialog.ShowError(fmt.Errorf("please enter a valid positive integer"), w)
		if onFinish != nil {
			onFinish()
		}
		return
	}

	if len(prizes) == 0 {
		dialog.ShowError(fmt.Errorf("no prize data loaded"), w)
		if onFinish != nil {
			onFinish()
		}
		return
	}

	if drawCount > len(prizes) {
		dialog.ShowError(fmt.Errorf("draw count exceeds available prizes"), w)
		if onFinish != nil {
			onFinish()
		}
		return
	}

	// 清空结果文本框
	resultsText.SetText("")

	// 生成不重复的随机索引
	indices := generateUniqueRandomNumbers(len(prizes), drawCount)
	if indices == nil {
		dialog.ShowError(fmt.Errorf("failed to generate unique random numbers"), w)
		if onFinish != nil {
			onFinish()
		}
		return
	}

	for i, idx := range indices {
		time.Sleep(100 * time.Millisecond) // 模拟抽奖过程中的延迟
		progress.SetValue(float64(i+1) / float64(drawCount) * 100)
		selectedPrize := prizes[idx]
		if len(selectedPrize) >= 2 {
			resultsText.SetText(resultsText.Text + fmt.Sprintf("%s %s\n", selectedPrize[0], selectedPrize[1]))
		} else if len(selectedPrize) == 1 {
			resultsText.SetText(resultsText.Text + fmt.Sprintf("%s\n", selectedPrize[0]))
		} else {
			resultsText.SetText(resultsText.Text + "Invalid prize data\n")
		}
	}

	// 抽奖结束后自动变为stop
	if onFinish != nil {
		onFinish()
	}
}

// GenButton creates a button that toggles between "Start" and "Stop" states,
// and triggers the provided onStart and onStop callbacks.
func GenButton(onStart func(), onStop func()) *widget.Button {
	drawButton := widget.NewButton("Start", nil)
	isRunning := false

	drawButton.OnTapped = func() {
		if isRunning {
			isRunning = false
			drawButton.SetText("Start")
			if onStop != nil {
				onStop()
			}
		} else {
			isRunning = true
			drawButton.SetText("Stop")
			if onStart != nil {
				onStart()
			}
		}
	}
	return drawButton
}

// ExampleButtonUsage demonstrates how to use GenButton with custom logic.
// 实现抽奖结束后自动变为stop
func ExampleButtonUsage(w fyne.Window, drawCountEntry, resultsText *widget.Entry, progress *widget.ProgressBar, prizes [][]string) *widget.Button {
	var drawButton *widget.Button
	var stopFunc func()

	// onFinish 用于抽奖结束后自动切换按钮状态
	onFinish := func() {
		if drawButton != nil {
			drawButton.SetText("Start")
		}
		if stopFunc != nil {
			stopFunc()
		}
	}

	drawButton = GenButton(
		func() { // onStart
			go LuckyDraw(w, drawCountEntry, resultsText, progress, prizes, onFinish)
		},
		func() { // onStop
			dialog.ShowInformation("Stopped", "The draw has been stopped.", w)
		},
	)
	// 保存 onStop 以便 onFinish 调用
	stopFunc = func() {
		dialog.ShowInformation("Stopped", "The draw has been stopped.", w)
	}
	return drawButton
}
