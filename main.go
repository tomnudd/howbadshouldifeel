package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"os"
	//"sync"
)

type ProductInfo struct {
	Code           string
	Status_verbose string
}

// PrettyPrint from https://stackoverflow.com/questions/19038598/how-can-i-pretty-print-json-using-go
func PrettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
}

func getProductInfo(barcodeNumber string, ptr *ProductInfo) error {
	var scrapeUrl string = fmt.Sprintf("https://world.openfoodfacts.org/api/v0/product/%v.json", barcodeNumber)

	req, err := http.NewRequest("GET", scrapeUrl, nil)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	query := req.URL.Query()
	query.Add("url", scrapeUrl)

	req.URL.RawQuery = query.Encode()

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Response code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	err = json.Unmarshal(body, &ptr)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil

}

func main() {
	product := ProductInfo{}

	getProductInfo("5449000000286", &product)

	PrettyPrint(product)

}
