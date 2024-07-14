package main

import (
	"calculator/request"
	"testing"
)

func TestGetTotalPrice(t *testing.T) {
	tests := []struct {
		name         string
		cart         request.Cart
		expected     float64
		expectingErr bool
	}{
		{
			name: "Customer1",
			cart: request.Cart{
				CartItems: []request.CartItem{
					{ItemName: "red_set", Amount: 1},
					{ItemName: "green_set", Amount: 2},
				},
				IsMember: false,
			},
			expected:     126,
			expectingErr: false,
		},
		{
			name: "Customer2",
			cart: request.Cart{
				CartItems: []request.CartItem{
					{ItemName: "red_set", Amount: 1},
					{ItemName: "green_set", Amount: 1},
					{ItemName: "green_set", Amount: 1},
				},
				IsMember: false,
			},
			expected:     126,
			expectingErr: false,
		},
		{
			name: "Customer3",
			cart: request.Cart{
				CartItems: []request.CartItem{
					{ItemName: "red_set", Amount: 1},
					{ItemName: "green_set", Amount: 1},
					{ItemName: "green_set", Amount: 1},
				},
				IsMember: true,
			},
			expected:     113.4,
			expectingErr: false,
		},
		{
			name: "Customer4",
			cart: request.Cart{
				CartItems: []request.CartItem{
					{ItemName: "red_set", Amount: 1},
					{ItemName: "green_set", Amount: 3},
					{ItemName: "green_set", Amount: 1},
					{ItemName: "orange_set", Amount: 5},
					{ItemName: "pink_set", Amount: 4},
				},
				IsMember: true,
			},
			expected:     973.8,
			expectingErr: false,
		},
		{
			name: "Customer5",
			cart: request.Cart{
				CartItems: []request.CartItem{},
				IsMember:  true,
			},
			expected:     0,
			expectingErr: false,
		},
		{
			name: "Customer6",
			cart: request.Cart{
				CartItems: []request.CartItem{
					{ItemName: "red_set", Amount: 1},
					{ItemName: "green_set", Amount: 3},
					{ItemName: "green_set", Amount: 1},
					{ItemName: "orange_set", Amount: 5},
					{ItemName: "pink_set", Amount: 4},
				},
				IsMember: false,
			},
			expected:     1082,
			expectingErr: false,
		},
		{
			name: "Customer7",
			cart: request.Cart{
				CartItems: []request.CartItem{
					{ItemName: "test", Amount: 6},	
				},
				IsMember: false,
			},
			expected:     540,
			expectingErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			totalPrice, err := GetTotalPrice(tt.cart)
			if totalPrice != tt.expected {
				t.Errorf("expected %f but got %f", tt.expected, totalPrice)
			}
			if (err != nil) != tt.expectingErr {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

