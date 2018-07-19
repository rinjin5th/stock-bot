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
}

// Add puts stock in DynamoDB
func (stock Stock) Add() (error) {
	tbl := NewTable(tableName)
	return tbl.Put(stock).Run() 
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