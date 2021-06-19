package controller

import (
	"basic-golang/opn/connector"
	"basic-golang/opn/schema"
	"fmt"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"log"
)

// func CreateToken(tamBoon interface{}) *omise.Token,error {
func CreateToken(tamBoon schema.TamBoon) (*omise.Token, error) {
	oClient := connector.OmiseCreateClient()
	token, createToken := &omise.Token{}, &operations.CreateToken{
		Name:            tamBoon.Name,
		Number:          tamBoon.CCNumber,
		ExpirationMonth: tamBoon.ExpMonth,
		ExpirationYear:  tamBoon.ExpYear,
	}

	if e := oClient.Do(token, createToken); e != nil {
		fmt.Errorf("%s", e)
		return token, e
	}

	log.Printf("token_id: %s  card_id: %s\n", token.ID, token.Card.ID)
	connector.OmiseCloseClient(oClient)
	return token, nil
}
