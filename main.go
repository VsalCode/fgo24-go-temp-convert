package main

import (
	"fmt"
	"os"
	"temp-convert/convert"
	"temp-convert/history"
)

var mainPage = `
======== WELCOME ========
1. Konversi Suhu
2. History Konversi
0. Exit
`

var DataConvertTemp []history.DataConvert

func main() {

	fmt.Print(mainPage)
	var choice string
	fmt.Scanln(&choice)
	switch choice {
	case "1":
		convert.TempConvert(&DataConvertTemp)
	case "2":
		history.HistoryTempConvert(DataConvertTemp)
	case "0":
		os.Exit(0)
	}
}
