package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

var nameFileBuy string = "./solicitudes_compra.cvs"
var nameFileSale string = "./solicitudes_venta.cvs"
var i int = 0
var j int = 0

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

func match(valuesBuy [][]string, valuesSale [][]string) ([][]string, [][]string) {

	if i == len(valuesBuy) || j == len(valuesBuy) {
		return valuesBuy, valuesSale
	}

	a, _ := strconv.Atoi(valuesBuy[i][2])
	b, _ := strconv.Atoi(valuesSale[j][2])

	if b > a {
		b = b - a
		a = 0
		valuesBuy[i][2] = strconv.Itoa(a)
		valuesSale[j][2] = strconv.Itoa(b)
		i++
	} else if b < a {
		a = a - b
		b = 0
		valuesBuy[i][2] = strconv.Itoa(a)
		valuesSale[j][2] = strconv.Itoa(b)
		j++
	} else {
		valuesBuy[i][2] = strconv.Itoa(0)
		valuesSale[j][2] = strconv.Itoa(0)
		i++
		j++
	}
	return match(valuesBuy, valuesSale)
}

func generatorResult(nameFile string, values [][]string) {
	file, err := os.Create(nameFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for _, value := range values {
		fmt.Fprintf(file, "%s\n", value)
	}
}

func main() {
	start := time.Now()
	recordsBuy := readCsvFile(nameFileBuy)
	recordsSale := readCsvFile(nameFileSale)
	valuesBuy, valuesSale := match(recordsBuy, recordsSale)
	generatorResult("solicitudes_compra_result.cvs", valuesBuy)
	generatorResult("solicitudes_venta_result.cvs", valuesSale)
	fmt.Println(time.Since(start))
}
