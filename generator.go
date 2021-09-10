package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	archivo, err := os.Create("./solicitudes_compra.cvs")
	if err != nil {
		fmt.Println(err)
		return
	}

	min := 5000
	max := 10000

	for i := 1; i <= 10; i++ {
		rand.Seed(time.Now().UnixNano())
		cantidad := rand.Intn(max-min) + min
		fmt.Fprintf(archivo, "o%d,energiaBasica,%d\n", i, cantidad)
	}
	archivo.Close()
}
