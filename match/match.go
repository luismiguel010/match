package main

import (
	"bufio"
	"encoding/csv"
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
	records := readCsvFile("./solicitudes_compra.cvs")
	fmt.Println(records[0][2])
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Unable to parse file as CSV for "+filePath, err)
	}
	return records
}
