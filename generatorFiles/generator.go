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

func generateFile(nameFile string, wg *sync.WaitGroup, lock *sync.Mutex) {
	var amount int
	defer wg.Done()
	lock.Lock()
	fmt.Println("Ingrese la cantidad a generar para" + nameFile)
	fmt.Scanln(&amount)
	lock.Unlock()
	start := time.Now()
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
	fmt.Println(time.Since(start))
}

func main() {

	var lock sync.Mutex
	var wg sync.WaitGroup
	wg.Add(1)
	go generateFile(nameFileBuy, &wg, &lock)
	wg.Add(1)
	go generateFile(nameFileSale, &wg, &lock)
	wg.Wait()

}
