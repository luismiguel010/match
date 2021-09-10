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

func evaluator(a int, b int, valuesBuy [][]string, valuesSale [][]string, file *os.File) {
	if b > a {
		aInicial := a
		bInicial := b
		b = b - a
		a = 0
		valuesBuy[i][2] = strconv.Itoa(a)
		valuesSale[j][2] = strconv.Itoa(b)
		fmt.Fprintf(file, "La C%d solicita comprar %d unidades, hace match con la V%d que tiene %d unidades disponibles; quedando C%d con %d unidades solicitadas y V%d con %d unidades disponibles.\n", i, aInicial, j, bInicial, i, a, j, b)
		i++
	} else if b < a {
		aInicial := a
		bInicial := b
		a = a - b
		b = 0
		valuesBuy[i][2] = strconv.Itoa(a)
		valuesSale[j][2] = strconv.Itoa(b)
		fmt.Fprintf(file, "La C%d solicita comprar %d unidades, hace match con la V%d que tiene %d unidades disponibles; quedando C%d con %d unidades solicitadas y V%d con %d unidades disponibles.\n", i, aInicial, j, bInicial, i, a, j, b)
		j++
	} else {
		aInicial := a
		bInicial := b
		a = 0
		b = 0
		fmt.Fprintf(file, "La C%d solicita comprar %d unidades, hace match con la V%d que tiene %d unidades disponibles; quedando C%d con %d unidades solicitadas y V%d con %d unidades disponibles.\n", i, aInicial, j, bInicial, i, a, j, b)
		valuesBuy[i][2] = strconv.Itoa(0)
		valuesSale[j][2] = strconv.Itoa(0)
		i++
		j++
	}
}

func match(valuesBuy [][]string, valuesSale [][]string, file *os.File) ([][]string, [][]string) {

	if i == len(valuesBuy) || j == len(valuesBuy) {
		return valuesBuy, valuesSale
	}

	a, _ := strconv.Atoi(valuesBuy[i][2])
	b, _ := strconv.Atoi(valuesSale[j][2])

	evaluator(a, b, valuesBuy, valuesSale, file)

	return match(valuesBuy, valuesSale, file)
}

func generatorRegisterSales(sales string) {
	file, err := os.Create("sales.cvs")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(file, sales)

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
	file, _ := os.Create("sales.cvs")
	recordsBuy := readCsvFile(nameFileBuy)
	recordsSale := readCsvFile(nameFileSale)
	valuesBuy, valuesSale := match(recordsBuy, recordsSale, file)
	generatorResult("solicitudes_compra_result.cvs", valuesBuy)
	generatorResult("solicitudes_venta_result.cvs", valuesSale)
	fmt.Println(time.Since(start))
}
