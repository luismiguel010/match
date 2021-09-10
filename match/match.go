package main

import (
	"bufio"
	"fmt"
	"os"
)

func leerArchivo() {
	archivo, err := os.Open("./solicitudes_compra.cvs")
	if err != nil {
		fmt.Println("Error" + err.Error())
	} else {
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {
			registro := scanner.Text()
			fmt.Printf("linea > " + registro + "\n")
		}
	}
	archivo.Close()
}

func main() {
	leerArchivo()
}
