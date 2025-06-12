package history

import (
	"fmt"
	"os"
)

type DataConvert struct {
	Fahrenheit int
	Kelvin     int
	Reamur     int
}

func HistoryTempConvert(data []DataConvert) {
	fmt.Println("======== WELCOME ========")
	for x := range data {
		fmt.Sprintf("%d. f: %d, k: %d, r: %d",x + 1 , data[x].Fahrenheit, data[x].Kelvin, data[x].Reamur)
	}

	fmt.Print("ketik 0 untuk konversi kembali")
	var choice string
	fmt.Scanln(&choice)
	switch choice {
	case "0":
		return
	default :
		os.Exit(0)
	}
}