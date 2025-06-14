package convert

import (
	"fmt"
	"os"
	"strconv"
	"temp-convert/history"
)

var tempData history.DataConvert

func TempConvert(data *[]history.DataConvert) {

	fmt.Println("========= CONVERT TEMP =========")
	fmt.Print("Masukkan suhu dalam bentuk celcius: ")
	var input string
	fmt.Scanln(&input)
	// convert.TempConvert(input)

	c, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Something Wrong!")
	}

	f := (9/5)*c + 32
	k := c + 273
	r := c * 4 / 5

	tempData = history.DataConvert{
		Base: input,
		Fahrenheit: f,
		Kelvin:     k,
		Reamur:     r,
	}

	*data = append(*data, tempData)

	fmt.Printf("Base Celcius: %s°C\n", input)
	fmt.Printf("--------------------\n")
	fmt.Printf("Fahrenheit: %d°C\n", f)
	fmt.Printf("Kelvin: %d°C\n", k)
	fmt.Printf("Reamur: %d°C\n", r)

	fmt.Print("ketik 0 untuk kembali ke menu: ")
	var choice string
	fmt.Scanln(&choice)
	switch choice {
	case "0":
		return
	default:
		os.Exit(0)
	}
}
