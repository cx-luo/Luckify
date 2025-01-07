// Package utils coding=utf-8
// @Project : lucky-draw
// @Time    : 2025/1/7 16:39
// @Author  : chengxiang.luo
// @Email   : chengxiang.luo@foxmail.com
// @File    : get_user.go
// @Software: GoLand
package utils

import (
	"encoding/csv"
	"os"
)

func ReadRecordsFromCsv(csvFilePath string) [][]string {
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		panic(err)
	}
	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {
			panic(err)
		}
	}(csvFile)

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	return records
}
