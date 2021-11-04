package controller

import (
	"basic-golang/opn/connector"
	"basic-golang/opn/schema"
	"fmt"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"log"
)

func CreateCharge(tamBoon schema.TamBoon, token *omise.Token) (*omise.Charge, *error) {
	//oClient := connector.OmiseCreateClient()
	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   tamBoon.AmountSubunits,
		Currency: "thb",
		Card:     token.ID,
	}
	if e := connector.POmiseClient.Do(charge, createCharge); e != nil {
		fmt.Println("Fail to create charge", e)
		return charge, &e
	}
	log.Printf("charge: %s  amount: %s %d\n", charge.ID, charge.Currency, charge.Amount)
	//connector.OmiseCloseClient(oClient)
	return charge, nil
}
