package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var nameFileBuy string = "./solicitudes_compra.cvs"
var nameFileSale string = "./solicitudes_venta.cvs"

func generateFile(nameFile string) {
	var amount int
	fmt.Println("Ingrese la cantidad a generar para" + nameFile)
	fmt.Scanln(&amount)
	file, err := os.Create(nameFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	max := 10000
	min := 5000
	var nameIdenficator string
	for i := 1; i <= amount; i++ {
		rand.Seed(time.Now().UnixNano())
		cantidad := rand.Intn(max-min) + min
		if nameFile == nameFileBuy {
			nameIdenficator = "C"
		} else {
			nameIdenficator = "V"
		}
		fmt.Fprintf(file, "%s%d,energiaBasica,%d,%s\n", nameIdenficator, i, cantidad, time.Now().UTC())
	}
}

func main() {
	start := time.Now()
	generateFile(nameFileBuy)
	generateFile(nameFileSale)
	fmt.Println(time.Since(start))
}
