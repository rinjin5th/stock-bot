package main

import (
	"strings"
	"errors"
)

const (
	prefix = "stk"
	commandPricelist = "pricelist"
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
	default: 
		return "", errors.New("invalid parameter")
	}
}