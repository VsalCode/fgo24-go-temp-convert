package history

import (
	"fmt"
	"os"
)

type DataConvert struct {
	Base string
	Fahrenheit int
	Kelvin     int
	Reamur     int
}

func HistoryTempConvert(data []DataConvert) {
	fmt.Println("======== LIST HISTORY ========")
	for x, item := range data {
		fmt.Printf("%d. Base Celcius: %s",x + 1, item.Base)
		fmt.Printf("f: %d°C, k: %d°C, r: %d°C\n", item.Fahrenheit, item.Kelvin, item.Reamur)
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