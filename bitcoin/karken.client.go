package bitcoin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type KrakenClient struct {
	client *http.Client
}

func New() *KrakenClient {
	return &KrakenClient{&http.Client{}}
}

func (cli *KrakenClient) Query(url string, param url.Values, returnType interface{}) (interface{}, error) {
	reqUrl := url + param.Encode()
	fmt.Printf("reqUrl = %+v\n", reqUrl)
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return returnType, fmt.Errorf("Error in creating request! %s", err.Error())
	}
	resp, err := cli.client.Do(req)
	if err != nil {
		fmt.Printf("err = resp %+v\n", err)
		return returnType, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(returnType); err != nil {
		fmt.Printf("err = %+v\n", err)
	}
	return returnType, nil
}

func (cli *KrakenClient) CalculateTotalBook(resp KarkenBookResponse) (float64, float64) {
	askTotal := calcuateTotal(resp.Result.XXBTZUSD.Asks)
	bidTotal := calcuateTotal(resp.Result.XXBTZUSD.Bids)
	return askTotal, bidTotal
}

func (cli *KrakenClient) CalculateTotalBuySell(resp KarkenHistoryResponse) (float64, float64) {
	var buysum, sellsum float64
	for _, v := range resp.Result.XXBTZUSD {
		var prev float64 = 1.0
		for index, inter := range v {
			if index < 2 {
				if unbox, ok := inter.(string); ok {
					// fmt.Printf("unbox = %+v\n", unbox)
					if f, err := strconv.ParseFloat(unbox, 64); err == nil {
						prev *= f
					}
				}
			}
		}
		if unbox, ok := v[3].(string); ok {
			if unbox == "s" {
				sellsum += prev
			}
			if unbox == "b" {
				buysum += prev
			}
		}
	}
	return buysum, sellsum

}
func calcuateTotal(arr []Item) float64 {
	var asksum float64
	for _, v := range arr {
		var prev float64 = 1.0
		for _, inter := range v {
			if unbox, ok := inter.(string); ok {
				fmt.Printf("unbox = %+v\n", unbox)
				if f, err := strconv.ParseFloat(unbox, 64); err == nil {
					prev *= f
				}
			}
		}
		asksum += prev
	}
	return asksum
}
