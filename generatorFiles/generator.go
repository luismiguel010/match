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
	file, err := os.Create(nameFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	max := 10000
	min := 5000

	for i := 1; i <= 5000; i++ {
		rand.Seed(time.Now().UnixNano())
		cantidad := rand.Intn(max-min) + min
		fmt.Fprintf(file, "o%d,energiaBasica,%d,%s\n", i, cantidad, time.Now().UTC())
	}
}

func main() {
	generateFile(nameFileBuy)
	generateFile(nameFileSale)
}
