package main

const (
	keyProfit = 1 + iota
)

const (
	tblMyState = "my_state"
)

// MyState is MyState object for DynamoDB
type MyState struct {
	Key int `dynamo:"key"`
	Value int `dynamo:"value"`
}

// ReflectInProfit updates profit
func ReflectInProfit(profit int) (int, error) {
	tbl := NewTable(tblMyState)
	var myState MyState
	err := tbl.Get("key", keyProfit).One(&myState)

	if err != nil {
		return 0, err
	}

	reflectedInProfit := myState.Value + profit
	err = tbl.Update("key", keyProfit).Set("value", reflectedInProfit).Run()
	
	if err != nil {
		return 0, err
	}
	
	return reflectedInProfit, nil
}

// GetProfit obtains profit data from the DynamoDB
func GetProfit() (int, error) {
	tbl := NewTable(tblMyState)
	var myState MyState
	err := tbl.Get("key", keyProfit).One(&myState)
	if err != nil {
		return 0, err
	}
	return myState.Value, nil
}