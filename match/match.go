package match

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
var COLUMNCOSTVALUE int = 3
var COLUMNTOLVALUE int = 4

/*readCsvFile lee una cadena con la ruta de un archivo cvs y lo convierte en una matriz*/
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

/*match permite hacer match entre solicitudes de compra y solicitudes de venta, basados en precios y tolerancias,
para que una compra haga match con una venta se debe respetar el rango que rige la tolerancia*/
func match(valueUnitBuy *string, valueCostBuy *string, valueTolBuy *string, valueUnitSale *string, valueCostSale *string, valuesBuy *[][]string, valuesSale *[][]string, file *os.File) ([][]string, [][]string) {

RUTINA:
	assingValues(valueUnitBuy, valueCostBuy, valueTolBuy, valueUnitSale, valueCostSale, valuesBuy, valuesSale)
	evaluator(castValues(valueUnitBuy), castValues(valueCostBuy), castValues(valueTolBuy), castValues(valueUnitSale), castValues(valueCostSale), valuesBuy, valuesSale, file)
	if counterBuyRow == len(*valuesBuy)-1 {
		return *valuesBuy, *valuesSale
	}
	goto RUTINA
}

/*castValues permite castear valores de tipo string a int*/
func castValues(value *string) *int {
	result, _ := strconv.Atoi(*value)
	return &result
}

/*assingValues asigna valores teniendo en cuenta variables globales que cambian en el transcurso del codigo*/
func assingValues(valueUnitBuy *string, valueCostBuy *string, valueTolBuy *string, valueUnitSale *string, valueCostSale *string, valuesBuy *[][]string, valuesSale *[][]string) {
	*valueUnitBuy = (*valuesBuy)[counterBuyRow][COLUMNUNITSVALUE]
	*valueCostBuy = (*valuesBuy)[counterBuyRow][COLUMNCOSTVALUE]
	*valueTolBuy = (*valuesBuy)[counterBuyRow][COLUMNTOLVALUE]
	*valueUnitSale = (*valuesSale)[counterSaleRow][COLUMNUNITSVALUE]
	*valueCostSale = (*valuesSale)[counterSaleRow][COLUMNCOSTVALUE]
}

/*generatorResult permite generar un archivo con los resultados*/
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

/*registerResults lleva el registro de las transacciones(match) que se van teniendo de comprar y ventas*/
func registerResults(file *os.File, a int, b int, c int, d int, e int, f int, g int, h int, i int, j int, k int) {
	fmt.Fprintf(file, "La C%d solicita comprar %d unidades a un precio de %d con tolerancia %d e hizo match con V%d que tenía %d unidades a un precio de %d; quedando C%d con %d unidades y V%d con %d unidades.\n", a, b, c, d, e, f, g, h, i, j, k)
}

func evaluator(valueUnitBuy *int, valueCostBuy *int, valueTolBuy *int, valueUnitSale *int, valueCostSale *int, valuesBuy *[][]string, valuesSale *[][]string, file *os.File) {
	result := *valueUnitSale - *valueUnitBuy
	if ((*valueCostBuy - *valueCostSale) <= *valueTolBuy) && (*valueUnitSale != 0) {
		cases(&result, valueUnitBuy, valueCostBuy, valueTolBuy, valueUnitSale, valueCostSale, valuesBuy, valuesSale, file)
	} else {
		counterSaleRow++
	}
	if counterSaleRow == len(*valuesSale)-1 {
		counterBuyRow++
		counterSaleRow = 0
	}
}

/*cases posee los casos basicos para lograr hacer el match entre unidades*/
func cases(result *int, valueUnitBuy *int, valueCostBuy *int, valueTolBuy *int, valueUnitSale *int, valueCostSale *int, valuesBuy *[][]string, valuesSale *[][]string, file *os.File) {
	switch {
	case *result > 0:
		registerResults(file, counterBuyRow, *valueUnitBuy, *valueCostBuy, *valueTolBuy, counterSaleRow, *valueUnitSale, *valueCostSale, counterBuyRow, 0, counterSaleRow, *valueUnitSale-*valueUnitBuy)
		(*valuesBuy)[counterBuyRow][COLUMNUNITSVALUE] = strconv.Itoa(0)
		(*valuesSale)[counterSaleRow][COLUMNUNITSVALUE] = strconv.Itoa(*valueUnitSale - *valueUnitBuy)
		counterBuyRow++
		counterSaleRow = 0
	case *result < 0:
		registerResults(file, counterBuyRow, *valueUnitBuy, *valueCostBuy, *valueTolBuy, counterSaleRow, *valueUnitSale, *valueCostSale, counterBuyRow, *valueUnitBuy-*valueUnitSale, counterSaleRow, 0)
		(*valuesBuy)[counterBuyRow][COLUMNUNITSVALUE] = strconv.Itoa(*valueUnitBuy - *valueUnitSale)
		(*valuesSale)[counterSaleRow][COLUMNUNITSVALUE] = strconv.Itoa(0)
		counterSaleRow++
	default:
		registerResults(file, counterBuyRow, *valueUnitBuy, *valueCostBuy, *valueTolBuy, counterSaleRow, *valueUnitSale, *valueCostSale, counterBuyRow, 0, counterSaleRow, 0)
		(*valuesBuy)[counterBuyRow][COLUMNUNITSVALUE] = strconv.Itoa(0)
		(*valuesSale)[counterSaleRow][COLUMNUNITSVALUE] = strconv.Itoa(0)
		counterBuyRow++
		counterSaleRow++
	}
}

/*main funcion principal que ejecuta el programa*/
func TotalMatch() {
	start := time.Now()
	file, _ := os.Create("sales.cvs")
	recordsBuy := readCsvFile(nameFileBuy)
	recordsSale := readCsvFile(nameFileSale)
	var valueUnitBuy, valueUnitSale, valueCostBuy, valueTolBuy, valueCostSale string
	valuesBuy, valuesSale := match(&valueUnitBuy, &valueCostBuy, &valueTolBuy, &valueUnitSale, &valueCostSale, &recordsBuy, &recordsSale, file)
	generatorResult("solicitudes_compra_result.cvs", valuesBuy)
	generatorResult("solicitudes_venta_result.cvs", valuesSale)
	fmt.Println(time.Since(start))
	counterBuyRow = 0
	counterSaleRow = 0
}
