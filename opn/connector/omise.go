package connector

import (
	"github.com/omise/omise-go"
	"log"
)

var POmiseClient *omise.Client

func OmiseCreateClient() *omise.Client {
	omiseClient, err := omise.NewClient(OmisePublicKey, OmiseSecretKey)
	if err != nil {
		log.Fatal(err)
	}
	return omiseClient
}

func OmiseCloseClient(omiseClient *omise.Client) {
	omiseClient.CloseIdleConnections()
}

func PConnection() {
	POmiseClient = OmiseCreateClient()
}
