package main

import (
	"basic-golang/opn/controller"
	"basic-golang/opn/helper"
	"basic-golang/opn/schema"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/omise/omise-go"
	"strconv"
	"time"
)

var fileInput string
var fileOutput string

func init() {
	flag.StringVar(&fileInput, "f", "", "-f path file .rot128")
	flag.StringVar(&fileOutput, "o", "", "-o path file .csv")
}
func main() {
	flag.Parse()
	fmt.Println(fileInput, fileOutput)
	cipReader := helper.Rot128ToCSV(fileInput)

	// Step 1 Read File .rot128 to CSV
	records, err := csv.NewReader(cipReader).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	// Step 2 Prepare Schema TamBoon
	tamBoons := []schema.TamBoon{}
	for i := 1; i < len(records); i++ {
		record := records[i]
		name, amountSubunits, ccNumber, cvv, expMonth, expYear := record[0], record[1], record[2], record[3], record[4], record[5]
		amount, _ := strconv.Atoi(amountSubunits)
		cvvI, _ := strconv.Atoi(cvv)
		expMonthI, _ := strconv.Atoi(expMonth)
		expYearI, _ := strconv.Atoi(expYear)
		tamBoon := schema.TamBoon{
			Name: name, AmountSubunits: int64(amount), CCNumber: ccNumber, CVV: cvvI, ExpMonth: time.Month(expMonthI), ExpYear: expYearI,
		}
		tamBoons = append(tamBoons, tamBoon)
	}

	txnTamBoon := schema.TxnTamBoon{}
	amountTamBoon := schema.AmountTamBoon{}
	mapNameAmount := make(map[string]int64)
	for i := 0; i < len(tamBoons) && i < 10; i++ {
		tamBoon := tamBoons[i]
		txnTamBoon.Total = txnTamBoon.Total + 1                            // Person
		amountTamBoon.Total = amountTamBoon.Total + tamBoon.AmountSubunits // Amount

		// Step 3 Create Token Per Record
		token, errCreateToken := controller.CreateToken(tamBoon)

		// Step 4 Create Charge Per Record
		errCreateCharge := new(error)
		charge := &omise.Charge{}
		if errCreateToken == nil {
			if charge, errCreateCharge = controller.CreateCharge(tamBoon, token); errCreateCharge == nil {
				txnTamBoon.Success = txnTamBoon.Success + 1
				amountTamBoon.Success = amountTamBoon.Success + charge.Amount
				if val, isExist := mapNameAmount[tamBoon.Name]; isExist {
					mapNameAmount[tamBoon.Name] = val + tamBoon.AmountSubunits
				} else {
					mapNameAmount[tamBoon.Name] = tamBoon.AmountSubunits
				}
			}
		}

		if errCreateToken != nil || errCreateCharge != nil {
			txnTamBoon.Fail = txnTamBoon.Fail + 1
			amountTamBoon.Fail = amountTamBoon.Fail + tamBoon.AmountSubunits
		}
	}

	// Step 5 Summary
	fmt.Println(txnTamBoon)
	fmt.Println(amountTamBoon)
	fmt.Println(mapNameAmount)

	//// Mr. Grossman R Oldbuck,2879410,5375543637862918,488,11,2021
	//tamBoon := schema.TamBoon{
	//	Name:           "Mr. Grossman R Oldbuck",
	//	AmountSubunits: 2879410,
	//	CCNumber:       "5375543637862918",
	//	CVV:            488,
	//	ExpMonth:       11,
	//	ExpYear:        2021,
	//}
	//
	//token := controller.CreateToken(tamBoon)
	//_ = controller.CreateCharge(tamBoon, token)
}
