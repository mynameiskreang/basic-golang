package main

import (
	"basic-golang/opn/cipher"
	"bytes"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const (
	// Read these from environment variables or configuration files!
	OmisePublicKey = "pkey_test_5o7hm6cr225xp9o1tdq"
	OmiseSecretKey = "skey_test_5o7hm6cr2f7ihcf3fg8"
)

func main() {
	client, e := omise.NewClient(OmisePublicKey, OmiseSecretKey)
	if e != nil {
		log.Fatal(e)
	}

	// Name,AmountSubunits,CCNumber,CVV,ExpMonth,ExpYear
	// Mr. Grossman R Oldbuck,2879410,5375543637862918,488,11,2021
	token, createToken := &omise.Token{}, &operations.CreateToken{
		Name:            "Mr. Grossman R Oldbuck",
		Number:          "5375543637862918",
		ExpirationMonth: 11,
		ExpirationYear:  2021,
	}

	if e := client.Do(token, createToken); e != nil {
		log.Fatal(e)
		client.CloseIdleConnections()
	}

	spew.Dump(token)
	log.Printf("token_id: %s  card_id: %s\n", token.ID, token.Card.ID)
	client.CloseIdleConnections()
}

func main2() {
	//client, e := omise.NewClient(OmisePublicKey, OmiseSecretKey)
	//if e != nil {
	//	log.Fatal(e)
	//}
	//
	///** Retrieve a token from a request
	// * A token should be created from a client side by using our client-side libraries
	// * https://www.omise.co/libraries#client-side-libraries
	// * More information:
	// * - https://www.omise.co/collecting-card-information
	// * - https://www.omise.co/security-best-practices
	// **/
	//token := "tokn_xxxxxxxxxxxxx"
	//
	//// Creates a charge from the token
	//charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
	//	Amount:   100000, // ฿ 1,000.00
	//	Currency: "thb",
	//	Card:     token.ID,
	//}
	//if e := client.Do(charge, createCharge); e != nil {
	//	log.Fatal(e)
	//}
	//
	//log.Printf("charge: %s  amount: %s %d\n", charge.ID, charge.Currency, charge.Amount)
}

func readCSV() {
	cf, _ := os.Create("./fng.1000.csv")
	f, _ := os.Open("./fng.1000.csv.rot128")
	b, _ := ioutil.ReadAll(f)

	byRead := make([]byte, len(b))
	f2, _ := os.Open("./fng.1000.csv.rot128")
	cipReader, _ := cipher.NewRot128Reader(f2)
	cipReader.Read(byRead)
	cf.Write(byRead)
}

func mainOld() {
	cf, _ := os.Create("./fng.1000.txt")
	f, _ := os.Open("./fng.1000.csv.rot128")
	cipReader, _ := cipher.NewRot128Reader(f)
	byRead := make([]byte, bytes.MinRead)
	for {
		n, err := cipReader.Read(byRead)
		fmt.Println(string(byRead[:n]))
		cf.Write(byRead)
		if err == io.EOF {
			break
			cf.Close()
		}
	}
}
