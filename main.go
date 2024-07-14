package main

import (
	"calculator/master"
	"calculator/model"
	"calculator/request"
	"errors"
	"fmt"
	"math"
)

func main() {
	customer1 := request.Cart{
		CartItems: []request.CartItem{{ItemName: "red_set", Amount: 1}, {ItemName: "green_set", Amount: 1}},
		IsMember:  false,
	}
	totalPriceCustomer1, err := GetTotalPrice(customer1)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Println("Total price of Customer1 is ", totalPriceCustomer1)

}

func InitItemInStore() map[string]model.ItemStore {
	
	itemsInStoreData := master.GetItemInStore()
	itemsInStore := map[string]model.ItemStore{}
	
	for _, item := range itemsInStoreData {

		itemsInStore[item.Name] = item

	}
	return itemsInStore
}

func GetTotalPrice(Cart request.Cart) (float64, error) {
	
	GroupByCartByName(&Cart)
	itemsInStore := InitItemInStore()
	totalPice := 0.0

	for _, cartItem := range Cart.CartItems {

		if item, ok := itemsInStore[cartItem.ItemName]; ok {

			if item.IsHavePromotion {
				totalPice += CalculatePromotion(cartItem.Amount, item)
			} else {
				totalPice += float64(cartItem.Amount) * item.Price
			}
		} else {

			errMsg := fmt.Sprintf("error from GetTotalPrice :no item named %s in this store", cartItem.ItemName)
			err := errors.New(errMsg)
			return 0, err

		}
	}

	if Cart.IsMember {
		
		ratio := (100 - master.PercentDiscountMember) / 100
		if ratio <= 0 {
			ratio = 0
		}
		totalPice = totalPice * ratio
	}
	return RoundTo2Decimal(totalPice), nil
}

func CalculatePromotion(amount int, itemInStore model.ItemStore) float64 {

	ratio := (100 - itemInStore.DiscountPercent) / 100
	if ratio <= 0 {
		ratio = 0
	}
	numberOfPair := amount / itemInStore.ConditionNumber
	numberOfMod := amount % itemInStore.ConditionNumber

	if numberOfPair > 0 {

		priceAfterDiscountBundle := ((float64(itemInStore.ConditionNumber) * float64(numberOfPair) * itemInStore.Price) * ratio)
		priceNotDiscount := (float64(numberOfMod) * itemInStore.Price)
		
		return priceAfterDiscountBundle + priceNotDiscount
	}

	return itemInStore.Price * float64(amount)
}

func GroupByCartByName(cart *request.Cart) {

	cartItemInMap := make(map[string]int)

	for _, item := range cart.CartItems {

		cartItemInMap[item.ItemName] += item.Amount

	}

	updatedCart := request.Cart{IsMember: cart.IsMember}
	for name, amount := range cartItemInMap {

		cartItem := request.CartItem{ItemName: name, Amount: amount}
		updatedCart.CartItems = append(updatedCart.CartItems, cartItem)

	}

	*cart = updatedCart
}


func RoundTo2Decimal(val float64) float64 {
	if val == 0 {
		return 0
	}
	return math.Round(val*100) / 100
}
