package main

import (
	"strings"
	"errors"
	"fmt"
	"strconv"

	"github.com/guregu/dynamo"
)

const (
	prefix = "stk"
	commandPricelist = "pricelist"
	commandAdd = "add"
	commandDel = "del"
	commandBuy = "buy"
	commandSell = "sell"
	commandProfit = "profit"
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
		err := Delete(code)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Stock(%s) Deleted.", code), nil
	case commandBuy:
		if len(parsedCommand) < 5 {
			return "", errors.New("invalid parameter")
		}
		code := parsedCommand[2]
		price, err := strconv.Atoi(parsedCommand[3])
		if err != nil {
			return "", err
		}
		unit, err := strconv.Atoi(parsedCommand[4])
		if err != nil {
			return "", err
		}

		var stock Stock
		stock, err = GetStock(code)

		if err == nil {
			err = UpdatePurchasePrice(code, price, unit)
			if err != nil {
				return "", err
			}
		} else if err == dynamo.ErrNotFound {
			stock = Stock{Code: code, PurchasePrice: price, Unit: unit}
			err = stock.Add()
			if err != nil {
				return "", err
			}
		} else {
			return "", err
		}

		return fmt.Sprintf("Stock(%s) bought.", code), nil
	case commandSell:
		if len(parsedCommand) < 4 {
			return "", errors.New("invalid parameter")
		}
		code := parsedCommand[2]
		price, err := strconv.Atoi(parsedCommand[3])
		if err != nil {
			return "", err
		}

		oldStock,err := GetStock(code)
		if err != nil {
			return "", err
		}

		profit := (price - oldStock.PurchasePrice) * 100 * oldStock.Unit - (100 * oldStock.Unit)
		reflectedInProfit, err := ReflectInProfit(profit)
		if err != nil {
			return "", err
		}

		err = Delete(code)
		if err != nil {
			return "", err
		}
		
		return fmt.Sprintf("Stock(%s) sold. Now, profit is (%d) yen.", code, reflectedInProfit), nil
	case commandProfit:
		profit, err := GetProfit()
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("Profit is %d", profit), nil
	default: 
		return "", errors.New("invalid parameter")
	}
}