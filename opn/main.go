package main

import (
	"basic-golang/opn/connector"
	"basic-golang/opn/controller"
	"basic-golang/opn/helper"
	"basic-golang/opn/schema"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/omise/omise-go"
	"sort"
	"strconv"
	"sync"
	"time"
)

var fileInput string
var fileOutput string

func init() {
	flag.StringVar(&fileInput, "f", "", "-f path file .rot128")
	flag.StringVar(&fileOutput, "o", "", "-o path file .csv")
	connector.PConnection()
}
func main() {
	flag.Parse()
	fmt.Println(fileInput, fileOutput)

	// Step 1 Read File .rot128 to CSV
	cipReader := helper.Rot128ToCSV(fileInput)
	records, err := csv.NewReader(cipReader).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	// Step 2 Prepare Schema TamBoon
	tamBoons := []schema.TamBoon{}
	for _, record := range records {
		name, amountSubunits, ccNumber, cvv, expMonth, expYear := record[0], record[1], record[2], record[3], record[4], record[5]
		amount, e1 := strconv.Atoi(amountSubunits)
		cvvI, e2 := strconv.Atoi(cvv)
		expMonthI, e3 := strconv.Atoi(expMonth)
		expYearI, e4 := strconv.Atoi(expYear)
		if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
			fmt.Errorf("strconv.Atoi Error: %s, %s, %s, %s", e1, e2, e3, e4)
		}
		tamBoon := schema.TamBoon{
			Name: name, AmountSubunits: int64(amount), CCNumber: ccNumber, CVV: cvvI, ExpMonth: time.Month(expMonthI), ExpYear: expYearI,
		}
		tamBoons = append(tamBoons, tamBoon)
	}

	txnTamBoon := schema.TxnTamBoon{}
	amountTamBoon := schema.AmountTamBoon{}
	mapNameAmount := make(map[string]int64)
	waitCreate := sync.WaitGroup{}
	createWorkers := make(chan bool, 3)

	for _, tamBoon := range tamBoons {
		waitCreate.Add(1)
		createWorkers <- true
		go func(tb schema.TamBoon) {
			txnTamBoon.Total = txnTamBoon.Total + 1                       // Person
			amountTamBoon.Total = amountTamBoon.Total + tb.AmountSubunits // Amount
			// Step 3 Create Token Per Record
			token, errCreateToken := controller.CreateToken(tb)
			// Step 4 Create Charge Per Record
			errCreateCharge := new(error)
			charge := &omise.Charge{}
			if errCreateToken == nil {
				if charge, errCreateCharge = controller.CreateCharge(tb, token); errCreateCharge == nil {
					txnTamBoon.Success = txnTamBoon.Success + 1
					amountTamBoon.Success = amountTamBoon.Success + charge.Amount
					if val, isExist := mapNameAmount[tb.Name]; isExist {
						mapNameAmount[tb.Name] = val + tb.AmountSubunits
					} else {
						mapNameAmount[tb.Name] = tb.AmountSubunits
					}
				}
			}

			if errCreateToken != nil || errCreateCharge != nil {
				txnTamBoon.Fail = txnTamBoon.Fail + 1
				amountTamBoon.Fail = amountTamBoon.Fail + tb.AmountSubunits
			}
			<-createWorkers
			waitCreate.Done()
		}(tamBoon)
		/*// Step 3 Create Token Per Record
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
		}*/
	}
	waitCreate.Wait()

	// Step 5 Done... Summary
	fmt.Println("performing donations...")
	fmt.Println("done.")

	//fmt.Println(mapNameAmount)
	//fmt.Println(amountTamBoon)
	//fmt.Println(txnTamBoon)

	summary(mapNameAmount, amountTamBoon, txnTamBoon)

	fmt.Println()
	fmt.Println("mapNameAmount", len(mapNameAmount), mapNameAmount)
	fmt.Println("amountTamBoon", amountTamBoon)
	fmt.Println("txnTamBoon", txnTamBoon)

}

func summary(mapNameAmount map[string]int64, amountTamBoon schema.AmountTamBoon, txnTamBoon schema.TxnTamBoon) {
	fmt.Println()
	fmt.Printf("total received: THB\t\t %s \n", humanize.Comma(amountTamBoon.Total))
	fmt.Printf("successfully donated: THB\t %s \n", humanize.Comma(amountTamBoon.Success))
	fmt.Printf("faulty donation: THB\t\t %s \n", humanize.Comma(amountTamBoon.Fail))
	fmt.Printf("average per person: THB\t\t %s \n", humanize.Commaf(float64(amountTamBoon.Success)/float64(txnTamBoon.Success)))
	fmt.Println()
	var es helper.Entries
	for k, v := range mapNameAmount {
		es = append(es, helper.Entry{Key: k, Value: v})
	}
	sort.Sort(es)
	if len(es) > 0 {
		fmt.Printf("Top 3 Donors:\t %s", es[0].Key)
		fmt.Printf("\n\t\t %s", es[1].Key)
		fmt.Printf("\n\t\t %s", es[2].Key)
	}
}
