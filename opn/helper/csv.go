package helper

import (
	"basic-golang/opn/cipher"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

/*func Rot128ToCSV(inputFile, outputFile string) *cipher.Rot128Reader {
	// inputFile = "./fng.1000.csv.rot128"
	// outputFIle = "./fng.1000.csv"
	f, _ := os.Open(inputFile)
	cf, _ := os.Create(outputFile)
	b, _ := ioutil.ReadAll(f)

	byRead := make([]byte, len(b))
	f2, _ := os.Open(inputFile)
	cipReader, _ := cipher.NewRot128Reader(f2)
	cipReader.Read(byRead)
	cf.Write(byRead)
	return cipReader
}*/

func Rot128ToCSV(inputFile string) *cipher.Rot128Reader {
	f, errF := os.Open(inputFile)
	if errF != nil {
		log.Fatalln(errF)
	}
	cipReader, errCip := cipher.NewRot128Reader(f)
	if errCip != nil {
		log.Fatalln(errCip)
	}
	return cipReader
}

func writeCSVWithBuffer() {
	cf, _ := os.Create("./fng.1000.csv")
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
