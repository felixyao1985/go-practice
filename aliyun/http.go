package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const AliAPI = "https://barcode100.market.alicloudapi.com"
const AliAppCode = "911eb358504943959e8d6c1646747720"

type aliCloud struct {
	barcodeURL string
	code       string
}

type MedicineBarcode struct {
	Barcode           string `json:"Barcode"`
	ItemName          string `json:"ItemName"`
	ItemClassCode     string `json:"ItemClassCode"`
	BrandName         string `json:"BrandName"`
	ItemSpecification string `json:"ItemSpecification"`
	FirmName          string `json:"FirmName"`
	FirmAddress       string `json:"FirmAddress"`
}

func newAliCloudCli() *aliCloud {
	return &aliCloud{
		barcodeURL: AliAPI,
		code:       AliAppCode,
	}
}
func (a *aliCloud) GetMedicine(c string) (*MedicineBarcode, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", a.barcodeURL+"/getBarcode?Code="+c, nil)

	if err != nil {
		return &MedicineBarcode{}, err
	}

	req.Header.Set("Authorization", "APPCODE "+a.code)
	fmt.Println(req)
	resp, err := client.Do(req)
	fmt.Println(resp)
	if err != nil {
		return &MedicineBarcode{Barcode: c}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body), err)
	p := &MedicineBarcode{}

	err = json.Unmarshal(body, p)

	if err != nil {
		return &MedicineBarcode{Barcode: c}, err
	}

	return p, err
}

//curl -i -k --get --include 'https://barcode100.market.alicloudapi.com/getBarcode?Code=6923450605288'  -H 'Authorization:APPCODE 911eb358504943959e8d6c1646747720'

func main() {
	var AliCloudCli = func() *aliCloud {
		return newAliCloudCli()
	}()

	fmt.Println(AliCloudCli.GetMedicine("6921850500769"))
	fmt.Println(AliCloudCli.GetMedicine("6909760990160"))
	fmt.Println(AliCloudCli.GetMedicine("6916783880064"))

}
