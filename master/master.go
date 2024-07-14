package master

import "calculator/model"

const PercentDiscountMember = 10.0
func GetItemInStore() []model.ItemStore {
	return []model.ItemStore{
		{Name: "red_set", Price: 50, IsHavePromotion: false},
		{Name: "green_set", Price: 40, IsHavePromotion: true, ConditionNumber: 2, DiscountPercent: 5},
		{Name: "blue_set", Price: 30, IsHavePromotion: false},
		{Name: "yellow_set", Price: 50, IsHavePromotion: false},
		{Name: "pink_set", Price: 80, IsHavePromotion: true, ConditionNumber: 2, DiscountPercent: 5},
		{Name: "purple_set", Price: 90, IsHavePromotion: false},
		{Name: "orange_set", Price: 120, IsHavePromotion: true, ConditionNumber: 2, DiscountPercent: 5},
		{Name: "test", Price: 100, IsHavePromotion: true, ConditionNumber: 3, DiscountPercent: 10},
	}
}

