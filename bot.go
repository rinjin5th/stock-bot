package main

import (
	"strings"
	"errors"
	"fmt"
)

const (
	prefix = "stk"
	commandPricelist = "pricelist"
	commandAdd = "add"
	commandDel = "del"
)

// ProcessCommand is execute received command.
func ProcessCommand(text string) (string, error) {
	parsedCommand := strings.Split(text, " ")
	if parsedCommand[0] != prefix || len(parsedCommand) < 2 {
		return "", errors.New("invalid parameter")
	}

	switch parsedCommand[1] {
	case commandPricelist:
		if len(parsedCommand) < 2 {
			return "", errors.New("invalid parameter")
		}
		stocks, err := AllStocks()
		if err != nil {
			return "", err
		}
		return joinStocks(stocks), nil
	case commandAdd:
		if len(parsedCommand) < 3 {
			return "", errors.New("invalid parameter")
		}
		code := parsedCommand[2]
		stock := Stock{Code: code}
		err := stock.Add()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Stock(%s) Added.", code), nil
	case commandDel:
		if len(parsedCommand) < 3 {
			return "", errors.New("invalid parameter")
		}
		code := parsedCommand[2]
		stock := Stock{Code: code}
		err := stock.Delete()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Stock(%s) Deleted.", code), nil
	default: 
		return "", errors.New("invalid parameter")
	}
}