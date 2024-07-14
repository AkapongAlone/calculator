package main

import (
	"calculator/request"
	"errors"
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

func main() {
	Customer1 := request.Cart{
		CartItems: []request.CartItem{request.CartItem{ItemName: "red_set", Amount: 1}, request.CartItem{ItemName: "green_set", Amount: 1}},
		IsMember:  false,
	}
	TotalPriceCustomer1, err := GetTotalPrice(Customer1)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Println("Total price of Customer1 is ", TotalPriceCustomer1)

	Customer2 := request.Cart{
		CartItems: []request.CartItem{request.CartItem{ItemName: "orange_set", Amount: 5}},
		IsMember:  false,
	}
	TotalPriceCustomer2, err := GetTotalPrice(Customer2)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Println("Total price of Customer2 is ", TotalPriceCustomer2)

	Customer3 := request.Cart{
		CartItems: []request.CartItem{request.CartItem{ItemName: "orange_set", Amount: 5}},
		IsMember:  true,
	}
	TotalPriceCustomer3, err := GetTotalPrice(Customer3)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Println("Total price of Customer3 is ", TotalPriceCustomer3)

	Customer4 := request.Cart{
		CartItems: []request.CartItem{request.CartItem{ItemName: "green_set", Amount: 5}, request.CartItem{ItemName: "purple_set", Amount: 5}},
		IsMember:  false,
	}
	TotalPriceCustomer4, err := GetTotalPrice(Customer4)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Println("Total price of Customer4 is ", TotalPriceCustomer4)

	Customer5 := request.Cart{
		CartItems: []request.CartItem{request.CartItem{ItemName: "red_set", Amount: 1}, request.CartItem{ItemName: "green_set", Amount: 1}, request.CartItem{ItemName: "green_set", Amount: 1}},
		IsMember:  false,
	}
	TotalPriceCustomer5, err := GetTotalPrice(Customer5)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Println("Total price of Customer5 is ", TotalPriceCustomer5)

}

func InitItemInStore() map[string]float64 {
	ItemsInStore := map[string]float64{
		"red_set":    50,
		"green_set":  40,
		"blue_set":   30,
		"yellow_set": 50,
		"pink_set":   80,
		"purple_set": 90,
		"orange_set": 120,
	}
	return ItemsInStore
}

func InitItemInPromotion() ([]string, error) {
	ItemInPromotion := []string{"green_set", "pink_set", "orange_set"}
	ItemsInStore := InitItemInStore()
	for _, ItemName := range ItemInPromotion {
		if _, ok := ItemsInStore[ItemName]; !ok {
			errMsg := fmt.Sprintf("error from InitItemInPromotion : no item named %s in this store", ItemName)
			err := errors.New(errMsg)
			return []string{}, err
		}
	}
	return ItemInPromotion, nil
}

func GetTotalPrice(Cart request.Cart) (float64, error) {
	var DiscountForMember float64
	CombineCart(&Cart)
	ItemsInStore := InitItemInStore()
	ItemInPromotion, err := InitItemInPromotion()
	if err != nil {
		return 0, err
	}
	TotalPice := 0.0
	for _, cartItem := range Cart.CartItems {
		if price, ok := ItemsInStore[cartItem.ItemName]; ok {
			if IsContainsSlice(cartItem.ItemName, ItemInPromotion) {
				TotalPice += CalculatePromotion(cartItem.Amount, price)
			} else {
				TotalPice += float64(cartItem.Amount) * price
			}
		} else {
			errMsg := fmt.Sprintf("error from GetTotalPrice :no item named %s in this store", cartItem.ItemName)
			err := errors.New(errMsg)
			return 0, err
		}
	}
	if Cart.IsMember {
		DiscountForMember = 10
		ratio := (100 - DiscountForMember) / 100
		if ratio <= 0 {
			ratio = 0
		}
		TotalPice = TotalPice * ratio
	}
	return RoundTo2Decimal(TotalPice), nil
}

func CalculatePromotion(amount int, price float64) float64 {
	var ConditionNumber int       //เงื่อนไขจำนวนสินค้าที่จะเข้าpromotion
	var DiscountPromotion float64 //percent ส่วนลด
	ConditionNumber = 2
	DiscountPromotion = 5
	ratio := (100 - DiscountPromotion) / 100
	if ratio <= 0 {
		ratio = 0
	}
	numberOfPair := amount / ConditionNumber
	numberOfMod := amount % ConditionNumber
	if numberOfPair > 0 {
		return ((float64(ConditionNumber) * float64(numberOfPair) * price) * ratio) + (float64(numberOfMod) * price)
	}
	return price * float64(amount)
}

func CombineCart(cart *request.Cart) {
	CartItemInMap := make(map[string]request.CartItem)
	for _, item := range cart.CartItems {
		if cartItem, ok := CartItemInMap[item.ItemName]; ok {
			cartItem.Amount += item.Amount
			CartItemInMap[item.ItemName] = cartItem
		} else {
			CartItemInMap[item.ItemName] = item
		}
	}
	updatedCart := request.Cart{IsMember: cart.IsMember}
	for _, item := range CartItemInMap {
		updatedCart.CartItems = append(updatedCart.CartItems, item)
	}
	*cart = updatedCart
}

func IsContainsSlice[T constraints.Ordered](item T, items []T) bool {
	for _, value := range items {
		if item == value {
			return true
		}
	}
	return false
}

func RoundTo2Decimal(val float64) float64 {
	if val == 0 {
		return 0
	}
	return math.Round(val*100) / 100
}
