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
var counterBuyRow int = 0
var counterSaleRow int = 0
var COLUMNUNITSVALUE int = 2

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

func evaluator(valueUnitBuy string, valueUnitSale string, valuesBuy *[][]string, valuesSale *[][]string, file *os.File) {
	unitSale, _ := strconv.Atoi(valueUnitSale)
	unitBuy, _ := strconv.Atoi(valueUnitBuy)
	result := unitSale - unitBuy
	switch {
	case result > 0:
		fmt.Fprintf(file, "La C%d solicita comprar %s unidades, hace match con la V%d que tiene %s unidades disponibles; quedando C%d con %d unidades solicitadas y V%d con %d unidades disponibles.\n", counterBuyRow, valueUnitBuy, counterSaleRow, valueUnitSale, counterBuyRow, 0, counterSaleRow, unitSale-unitBuy)
		(*valuesBuy)[counterBuyRow][COLUMNUNITSVALUE] = strconv.Itoa(0)
		(*valuesSale)[counterSaleRow][COLUMNUNITSVALUE] = strconv.Itoa(unitSale - unitBuy)
		counterBuyRow++
	case result < 0:
		fmt.Fprintf(file, "La C%d solicita comprar %s unidades, hace match con la V%d que tiene %s unidades disponibles; quedando C%d con %d unidades solicitadas y V%d con %d unidades disponibles.\n", counterBuyRow, valueUnitBuy, counterSaleRow, valueUnitSale, counterBuyRow, unitBuy-unitSale, counterSaleRow, 0)
		(*valuesBuy)[counterBuyRow][COLUMNUNITSVALUE] = strconv.Itoa(unitBuy - unitSale)
		(*valuesSale)[counterSaleRow][COLUMNUNITSVALUE] = strconv.Itoa(0)
		counterSaleRow++
	default:
		fmt.Fprintf(file, "La C%d solicita comprar %s unidades, hace match con la V%d que tiene %s unidades disponibles; quedando C%d con %d unidades solicitadas y V%d con %d unidades disponibles.\n", counterBuyRow, valueUnitBuy, counterSaleRow, valueUnitSale, counterBuyRow, 0, counterSaleRow, 0)
		(*valuesBuy)[counterBuyRow][COLUMNUNITSVALUE] = strconv.Itoa(0)
		(*valuesSale)[counterSaleRow][COLUMNUNITSVALUE] = strconv.Itoa(0)
		counterBuyRow++
		counterSaleRow++
	}
}

func match(valueUnitBuy string, valueUnitSale string, valuesBuy *[][]string, valuesSale *[][]string, file *os.File) ([][]string, [][]string) {

RUTINA:
	valueUnitBuy = (*valuesBuy)[counterBuyRow][COLUMNUNITSVALUE]
	valueUnitSale = (*valuesSale)[counterSaleRow][COLUMNUNITSVALUE]

	evaluator(valueUnitBuy, valueUnitSale, valuesBuy, valuesSale, file)

	goto RUTINA
	if counterBuyRow == len(*valuesBuy) || counterSaleRow == len(*valuesBuy) {
		return *valuesBuy, *valuesSale
	}
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
	valueUnitBuy := recordsBuy[counterBuyRow][COLUMNUNITSVALUE]
	valueUnitSale := recordsSale[counterSaleRow][COLUMNUNITSVALUE]

	valuesBuy, valuesSale := match(valueUnitBuy, valueUnitSale, &recordsBuy, &recordsSale, file)
	generatorResult("solicitudes_compra_result.cvs", valuesBuy)
	generatorResult("solicitudes_venta_result.cvs", valuesSale)
	fmt.Println(time.Since(start))
}
