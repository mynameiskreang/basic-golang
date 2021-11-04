package schema

import "time"

// TamBoon
// Name,AmountSubunits,CCNumber,CVV,ExpMonth,ExpYear
// Mr. Grossman R Oldbuck,2879410,5375543637862918,488,11,2021
type TamBoon struct {
	Name           string     `json:"name"`
	AmountSubunits int64      `json:"amountSubunits"`
	CCNumber       string     `json:"ccNumber"`
	CVV            int        `json:"cvv"`
	ExpMonth       time.Month `json:"expMonth"`
	ExpYear        int        `json:"expYear"`
}

type TxnTamBoon struct {
	Total   int `json:"total"`
	Success int `json:"success"`
	Fail    int `json:"fail"`
}

type AmountTamBoon struct {
	Total   int64 `json:"total"`
	Success int64 `json:"success"`
	Fail    int64 `json:"fail"`
}
