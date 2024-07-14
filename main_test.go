package main

import (
	"calculator/request"
	"testing"
)

func TestGetTotalPrice(t *testing.T) {
	Customer1 := request.Cart{
		CartItems: []request.CartItem{request.CartItem{ItemName: "red_set", Amount: 1}, request.CartItem{ItemName: "green_set", Amount: 2}},
		IsMember:  false,
	}
	TotalPriceCustomer1, err := GetTotalPrice(Customer1)
	if TotalPriceCustomer1 != 126 {
		t.Errorf("result should equal 126 but got %f", TotalPriceCustomer1)
	} else if err != nil {
		t.Error(err)
	}

	Customer2 := request.Cart{
		CartItems: []request.CartItem{request.CartItem{ItemName: "red_set", Amount: 1}, request.CartItem{ItemName: "green_set", Amount: 1}, request.CartItem{ItemName: "green_set", Amount: 1}},
		IsMember:  false,
	}
	TotalPriceCustomer2, err := GetTotalPrice(Customer2)
	if TotalPriceCustomer2 != 126 {
		t.Errorf("result should equal 126 but got %f", TotalPriceCustomer2)
	} else if err != nil {
		t.Error(err)
	}

	Customer3 := request.Cart{
		CartItems: []request.CartItem{request.CartItem{ItemName: "red_set", Amount: 1}, request.CartItem{ItemName: "green_set", Amount: 1}, request.CartItem{ItemName: "green_set", Amount: 1}},
		IsMember:  true,
	}
	TotalPriceCustomer3, err := GetTotalPrice(Customer3)
	if TotalPriceCustomer3 != 113.4 {
		t.Errorf("result should equal 113.4 but got %f", TotalPriceCustomer3)
	} else if err != nil {
		t.Error(err)
	}


	Customer4 := request.Cart{
		CartItems: []request.CartItem{request.CartItem{ItemName: "red_set", Amount: 1}, request.CartItem{ItemName: "green_set", Amount: 3}, request.CartItem{ItemName: "green_set", Amount: 1}, request.CartItem{ItemName: "orange_set", Amount: 5}, request.CartItem{ItemName: "pink_set", Amount: 4}},
		IsMember:  true,
	}
	TotalPriceCustomer4, err := GetTotalPrice(Customer4)
	if TotalPriceCustomer4 != 973.8 {
		t.Errorf("result should equal 973.8 but got %f", TotalPriceCustomer4)
	} else if err != nil {
		t.Error(err)
	}

	Customer5 := request.Cart{
		CartItems: []request.CartItem{},
		IsMember:  true,
	}
	TotalPriceCustomer5, err := GetTotalPrice(Customer5)
	if TotalPriceCustomer5 != 0 {
		t.Errorf("result should equal 0 but got %f", TotalPriceCustomer5)
	} else if err != nil {
		t.Error(err)
	}

	Customer6 := request.Cart{
		CartItems: []request.CartItem{request.CartItem{ItemName: "red_set", Amount: 1}, request.CartItem{ItemName: "green_set", Amount: 3}, request.CartItem{ItemName: "green_set", Amount: 1}, request.CartItem{ItemName: "orange_set", Amount: 5}, request.CartItem{ItemName: "pink_set", Amount: 4}},
		IsMember:  false,
	}
	TotalPriceCustomer6, err := GetTotalPrice(Customer6)
	if TotalPriceCustomer6 != 1082 {
		t.Errorf("result should equal 1082 but got %f", TotalPriceCustomer6)
	} else if err != nil {
		t.Error(err)
	}
}
