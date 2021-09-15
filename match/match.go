package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

var nameFileBuy string = "./solicitudes_compra.cvs"
var nameFileSale string = "./solicitudes_venta.cvs"
var counterBuyRow int = 0
var counterSaleRow int = 0
var COLUMNUNITSVALUE int = 2
var COLUMNCOSTVALUE int = 3
var COLUMNTOLVALUE int = 4

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

func evaluator(valueUnitBuy *string, valueCostBuy *string, valueTolBuy *string, valueUnitSale *string, valueCostSale *string, valuesBuy *[][]string, valuesSale *[][]string, file *os.File) {

	unitBuy, _ := strconv.Atoi(*valueUnitBuy)
	costBuy, _ := strconv.Atoi(*valueCostBuy)
	tolBuy, _ := strconv.Atoi(*valueTolBuy)

	unitSale, _ := strconv.Atoi(*valueUnitSale)
	costSale, _ := strconv.Atoi(*valueCostSale)
	result := unitSale - unitBuy
	if math.Abs(float64(costBuy-costSale)) <= float64(tolBuy) && unitSale != 0 {
		switch {
		case result > 0:
			registerResults(file, counterBuyRow, *valueUnitBuy, costBuy, tolBuy, counterSaleRow, *valueUnitSale, costSale, counterBuyRow, 0, counterSaleRow, unitSale-unitBuy)
			(*valuesBuy)[counterBuyRow][COLUMNUNITSVALUE] = strconv.Itoa(0)
			(*valuesSale)[counterSaleRow][COLUMNUNITSVALUE] = strconv.Itoa(unitSale - unitBuy)
			counterBuyRow++
			counterSaleRow = 0
		case result < 0:
			registerResults(file, counterBuyRow, *valueUnitBuy, costBuy, tolBuy, counterSaleRow, *valueUnitSale, costSale, counterBuyRow, unitBuy-unitSale, counterSaleRow, 0)
			(*valuesBuy)[counterBuyRow][COLUMNUNITSVALUE] = strconv.Itoa(unitBuy - unitSale)
			(*valuesSale)[counterSaleRow][COLUMNUNITSVALUE] = strconv.Itoa(0)
			counterSaleRow++
		default:
			registerResults(file, counterBuyRow, *valueUnitBuy, costBuy, tolBuy, counterSaleRow, *valueUnitSale, costSale, counterBuyRow, 0, counterSaleRow, 0)
			(*valuesBuy)[counterBuyRow][COLUMNUNITSVALUE] = strconv.Itoa(0)
			(*valuesSale)[counterSaleRow][COLUMNUNITSVALUE] = strconv.Itoa(0)
			counterBuyRow++
			counterSaleRow++
		}
	} else {
		counterSaleRow++
	}
	if counterSaleRow == len(*valuesSale)-1 {
		counterBuyRow++
		counterSaleRow = 0
	}
}

func match(valueUnitBuy *string, valueCostBuy *string, valueTolBuy *string, valueUnitSale *string, valueCostSale *string, valuesBuy *[][]string, valuesSale *[][]string, file *os.File) ([][]string, [][]string) {

RUTINA:
	*valueUnitBuy = (*valuesBuy)[counterBuyRow][COLUMNUNITSVALUE]
	*valueCostBuy = (*valuesBuy)[counterBuyRow][COLUMNCOSTVALUE]
	*valueTolBuy = (*valuesBuy)[counterBuyRow][COLUMNTOLVALUE]

	*valueUnitSale = (*valuesSale)[counterSaleRow][COLUMNUNITSVALUE]
	*valueCostSale = (*valuesSale)[counterSaleRow][COLUMNCOSTVALUE]

	evaluator(valueUnitBuy, valueCostBuy, valueTolBuy, valueUnitSale, valueCostSale, valuesBuy, valuesSale, file)

	if counterBuyRow == len(*valuesBuy)-1 {
		return *valuesBuy, *valuesSale
	}
	goto RUTINA
}

func generatorResult(nameFile string, values [][]string) {
	file, err := os.Create(nameFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	for _, value := range values {
		fmt.Fprintf(file, "%s\n", value)
	}
}

func registerResults(file *os.File, a int, b string, c int, d int, e int, f string, g int, h int, i int, j int, k int) {
	fmt.Fprintf(file, "La C%d solicita comprar %s unidades a un precio de %d con tolerancia %d e hizo match con V%d que tenía %s unidades a un precio de %d; quedando C%d con %d unidades y V%d con %d unidades.\n", a, b, c, d, e, f, g, h, i, j, k)
}

func main() {
	start := time.Now()
	file, _ := os.Create("sales.cvs")
	recordsBuy := readCsvFile(nameFileBuy)
	recordsSale := readCsvFile(nameFileSale)
	valueUnitBuy := recordsBuy[counterBuyRow][COLUMNUNITSVALUE]
	fmt.Println("Hola aqui esta el tamaño", len(recordsBuy))
	valueUnitSale := recordsSale[counterSaleRow][COLUMNUNITSVALUE]
	valueCostBuy := recordsBuy[counterBuyRow][COLUMNCOSTVALUE]
	valueTolBuy := recordsBuy[counterBuyRow][COLUMNTOLVALUE]
	valueCostSale := recordsBuy[counterBuyRow][COLUMNCOSTVALUE]
	valuesBuy, valuesSale := match(&valueUnitBuy, &valueCostBuy, &valueTolBuy, &valueUnitSale, &valueCostSale, &recordsBuy, &recordsSale, file)
	generatorResult("solicitudes_compra_result.cvs", valuesBuy)
	generatorResult("solicitudes_venta_result.cvs", valuesSale)
	fmt.Println(time.Since(start))
}
