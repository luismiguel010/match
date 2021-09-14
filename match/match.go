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

	if math.Abs(float64(costBuy-costSale)) <= float64(tolBuy) {
		switch {
		case result > 0:
			//registerResults(file, counterBuyRow, *valueUnitBuy, counterSaleRow, *valueUnitSale, counterBuyRow, 0, counterSaleRow, unitSale-unitBuy)
			(*valuesBuy)[counterBuyRow][COLUMNUNITSVALUE] = strconv.Itoa(0)
			(*valuesSale)[counterSaleRow][COLUMNUNITSVALUE] = strconv.Itoa(unitSale - unitBuy)
			counterBuyRow++
			counterSaleRow = 0
		case result < 0:
			//registerResults(file, counterBuyRow, *valueUnitBuy, counterSaleRow, *valueUnitSale, counterBuyRow, unitBuy-unitSale, counterSaleRow, 0)
			(*valuesBuy)[counterBuyRow][COLUMNUNITSVALUE] = strconv.Itoa(unitBuy - unitSale)
			(*valuesSale)[counterSaleRow][COLUMNUNITSVALUE] = strconv.Itoa(0)
			counterSaleRow++
		default:
			//registerResults(file, counterBuyRow, *valueUnitBuy, counterSaleRow, *valueUnitSale, counterBuyRow, 0, counterSaleRow, 0)
			(*valuesBuy)[counterBuyRow][COLUMNUNITSVALUE] = strconv.Itoa(0)
			(*valuesSale)[counterSaleRow][COLUMNUNITSVALUE] = strconv.Itoa(0)
			counterBuyRow++
			counterSaleRow++
		}
	}
	counterSaleRow++
	if counterSaleRow == len(*valuesSale) {
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
	*valueCostSale = (*valuesBuy)[counterBuyRow][COLUMNCOSTVALUE]

	evaluator(valueUnitBuy, valueCostBuy, valueTolBuy, valueUnitSale, valueCostSale, valuesBuy, valuesSale, file)

	if counterBuyRow == len(*valuesBuy) {
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

func registerResults(file *os.File, a int, b string, c int, d string, e int, f int, g int, h int) {
	fmt.Fprintf(file, "La C%d solicita comprar %s unidades, hace match con la V%d que tiene %s unidades disponibles; quedando C%d con %d unidades solicitadas y V%d con %d unidades disponibles.\n", a, b, c, d, e, f, g, h)
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
