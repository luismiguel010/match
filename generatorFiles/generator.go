package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var nameFileBuy string = "./solicitudes_compra.cvs"
var nameFileSale string = "./solicitudes_venta.cvs"

func generateFile(nameFile string, amount int, wg *sync.WaitGroup, lock *sync.Mutex) {
	var nameIdenficator string
	var maxValues int = 10000
	var minValues int = 5000
	var maxCost int = 20
	var minCost int = 10
	var maxTol int = 3
	var minTol int = 0
	var value int
	var cost int
	var tol int

	defer wg.Done()
	file, err := os.Create(nameFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for i := 1; i <= amount; i++ {
		rand.Seed(time.Now().UnixNano())
		value = rand.Intn(maxValues-minValues) + minValues
		cost = rand.Intn(maxCost-minCost) + minCost
		tol = rand.Intn(maxTol-minTol) + minTol
		if nameFile == nameFileBuy {
			nameIdenficator = "C"
			fmt.Fprintf(file, "%s%d,energiaBasica,%d,%d,%d,%s\n", nameIdenficator, i, value, cost, tol, time.Now().UTC())
		} else {
			nameIdenficator = "V"
			fmt.Fprintf(file, "%s%d,energiaBasica,%d,%d,%s\n", nameIdenficator, i, value, cost, time.Now().UTC())
		}
	}
}

func main() {

	var lock sync.Mutex
	var wg sync.WaitGroup

	var amount int
	fmt.Println("Ingrese la cantidad a generar de los dos archivos")
	fmt.Scanln(&amount)

	if amount != 0 {
		wg.Add(1)
		go generateFile(nameFileBuy, amount, &wg, &lock)
		wg.Add(1)
		go generateFile(nameFileSale, amount, &wg, &lock)
	}

	wg.Wait()

}
