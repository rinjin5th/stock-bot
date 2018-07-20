package main

import (
	"fmt"
	"errors"
)

const (
	tableName = "stock"
)

// Stock is stock object for DynamoDB
type Stock struct {
	Code  string `dynamo:"code"`
	CompanyName string `dynamo:"company_name"`
	Price int `dynamo:"price"`
	PurchasePrice int `dynamo:"purchase_price"`
}

// Add puts stock in DynamoDB
func (stock Stock) Add() (error) {
	tbl := NewTable(tableName)
	return tbl.Put(stock).Run() 
}

// Delete deletes stock in DynamoDB
func (stock Stock) Delete() (error) {
	tbl := NewTable(tableName)
	return tbl.Delete("code", stock.Code).Run() 
}

// AllStocks obtains all stock data in DynamoDB
func AllStocks() ([]Stock, error) {
	tbl := NewTable(tableName)
	var stocks []Stock
	err := tbl.Scan().All(&stocks)

	if err != nil {
		return nil, err
	}

	if len(stocks) == 0 {
		return nil, errors.New("Not found stocks")
	}

	return stocks, nil
}

// GetStock obtains stock data by the code in DynamoDB
func GetStock(code string) (Stock, error) {
	tbl := NewTable(tableName)
	var stock Stock
	err := tbl.Get("code", code).One(&stock)
	return stock, err
}

// UpdatePurchasePrice updates stock's purchase_price
func UpdatePurchasePrice(code string, price int) (error) {
	tbl := NewTable(tableName)
	return tbl.Update("code", code).Set("purchase_price", price).Run()
}

func joinStocks(stocks []Stock) (string){
	switch len(stocks) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("%+v", stocks[0])
	case 2:
		return fmt.Sprintf("%+v\n%+v", stocks[0], stocks[1])
	case 3:
		return fmt.Sprintf("%+v\n%+v\n%+v", stocks[0], stocks[1], stocks[2])
	}
	n := len("\n") * (len(stocks) - 1)
	for i := 0; i < len(stocks); i++ {
		n += len(fmt.Sprintf("%+v", stocks[i]))
	}

	b := make([]byte, n)
	bp := copy(b, fmt.Sprintf("%+v", stocks[0]))
	for _, stock := range stocks[1:] {
		bp += copy(b[bp:], "\n")
		bp += copy(b[bp:], fmt.Sprintf("%+v", stock))
	}
	return string(b)
}