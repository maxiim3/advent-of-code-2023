package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/maxiim3/aoc2023/datatypes"
)

// Use fmt.Sprintf to replace %s by the currency
const apiUrl = "https://cex.io/api/ticker/%s/USD"

// Makes the api call to cex crypto exchange
// We return the pointer to datatypes.Rate in order to return a nil value (If we are handling an error)
func GetRate(currency string) (*datatypes.Rate, error) {
	u := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, u))

	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK {

		// read the body of the response
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		// parse the json response
		json, err := UnmarshalApi(bodyBytes)

		if err != nil {
			return nil, err
		}

		// converts the price to a float
		price, err := json.convertToFloat()

		if err != nil {
			return nil, err
		}

		// creates the object structure
		rate := datatypes.Rate{Currency: currency, Price: *price}

		return &rate, nil
	} else {
		return nil, fmt.Errorf("status code received : %v", res.StatusCode)
	}
}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

func UnmarshalApi(data []byte) (ApiResponse, error) {
	var r ApiResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ApiResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ApiResponse struct {
	Timestamp             string  `json:"timestamp"`
	Low                   string  `json:"low"`
	High                  string  `json:"high"`
	Last                  string  `json:"last"`
	Volume                string  `json:"volume"`
	Volume30D             string  `json:"volume30d"`
	Bid                   float64 `json:"bid"`
	Ask                   float64 `json:"ask"`
	PriceChange           string  `json:"priceChange"`
	PriceChangePercentage string  `json:"priceChangePercentage"`
	Pair                  string  `json:"pair"`
}

func (r *ApiResponse) convertToFloat() (*float64, error) {
	convert, err := strconv.ParseFloat(r.Last, 64)
	if err != nil {
		return nil, err
	}
	return &convert, nil
}

func AskCex(currency string) {
	i := 0
	previousPrice := float64(0.00)
	for i < 10 {

		rate, err := GetRate(currency)

		if err != nil {
			fmt.Println(err)
		}

		if previousPrice < float64(1000) {
			fmt.Printf("%s is currently $%f\n", rate.Currency, rate.Price)
		} else {
			switch {
			case previousPrice > rate.Price:
				fmt.Printf("%v went up ! from $%f to $%f\n", rate.Currency, previousPrice, rate.Price)
			case previousPrice == rate.Price:
				fmt.Printf("%v stayed the same ! $%f\n", rate.Currency, previousPrice)
			default:
				fmt.Printf("%v went down... ! from $%f to $%f\n", rate.Currency, previousPrice, rate.Price)
			}
		}

		i++
		previousPrice = rate.Price
		time.Sleep(250 * time.Millisecond)
	}
}
