package bitcoin

type KarkenBookResponse struct {
	Result TradeBookResult `json:"result"`
}

type KarkenHistoryResponse struct {
	Result TradeHistoryResult `json:"result"`
}
type Item []interface{}

type XXBTZUSDBook struct {
	Asks []Item `json:"asks"`
	Bids []Item `json:"bids"`
}

type TradeBookResult struct {
	XXBTZUSD XXBTZUSDBook `json:"XXBTZUSD"`
}

type TradeHistoryResult struct {
	XXBTZUSD []Item `json:"XXBTZUSD"`
}
