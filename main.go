package main

import (
	"error_handling/utils"
	"errors"
	"fmt"
)

var (
	ErrorFoodType   = fmt.Errorf("invalid type")
	ErrorFoodAmount = fmt.Errorf("invalid amount")
)

func eat(food_type, amount int) (bool, error) {
	if food_type != 0 {
		return false, ErrorFoodType
	} else if amount > 10 {
		return false, ErrorFoodAmount
	}
	fmt.Println("good to eat !!")
	return true, nil
}

func smell(food_type, amount int) (bool, error) {
	if food_type != 0 {
		return false, ErrorFoodType
	} else if amount > 5 {
		return false, ErrorFoodAmount
	}
	fmt.Println("good to smell !!")
	return true, nil
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			if errors.Is(err.(error), ErrorFoodType) {
				fmt.Println("error handling for FoodTypeError")
			} else if errors.Is(err.(error), ErrorFoodAmount) {
				fmt.Println("error handling for FoodAmountError")
			}
		}
	}()
	for i := range 5 {
		ok, err := smell(i, 0)
		if !ok {
			fmt.Printf("smell not good, but it's ok: %s\n", err)
		}
	}

	utils.Must(eat(0, 1))
	utils.Must(eat(1, 1))
	utils.Must(eat(0, 6))
}
