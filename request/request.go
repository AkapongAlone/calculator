package request

type CartItem struct {
	ItemName string 
	Amount   int
}

type Cart struct {
	CartItems []CartItem
	IsMember bool
}
