package generatorfiles

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var nameFileBuy string = "./solicitudes_compra.cvs"
var nameFileSale string = "./solicitudes_venta.cvs"

func createFiles(nameFile string, amount int, wg *sync.WaitGroup, lock *sync.Mutex) {
	var maxValues, minValues, maxCost, minCost, maxTol, minTol int = 10000, 5000, 20, 10, 3, 0
	var nameIdenficator string
	var value, cost, tol int
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

func GeneratorFiles(amount int) {
	var lock sync.Mutex
	var wg sync.WaitGroup
	if amount != 0 {
		wg.Add(1)
		go createFiles(nameFileBuy, amount, &wg, &lock)
		wg.Add(1)
		go createFiles(nameFileSale, amount, &wg, &lock)
	}
	wg.Wait()
}
