package main

const (
	keyProfit = "1"
)

const (
	tblStockState = "stock_state"
)

// StockState is StockState object for DynamoDB
type StockState struct {
	Key string `dynamo:"key"`
	Value int `dynamo:"value"`
}

// ReflectInProfit updates profit
func ReflectInProfit(profit int) (int, error) {
	tbl := NewTable(tblStockState)
	var stockState StockState
	err := tbl.Get("key", keyProfit).One(&stockState)

	if err != nil {
		return 0, err
	}

	reflectedInProfit := stockState.Value + profit
	err = tbl.Update("key", keyProfit).Set("value", reflectedInProfit).Run()
	
	if err != nil {
		return 0, err
	}
	
	return reflectedInProfit, nil
}