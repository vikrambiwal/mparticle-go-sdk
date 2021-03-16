package events

// ProductAction struct for ProductAction
type ProductAction struct {
	Action            ProductActionType `json:"action"`
	CheckoutStep      *int64            `json:"checkout_step,omitempty"`
	CheckoutOptions   string            `json:"checkout_options,omitempty"`
	ProductActionList string            `json:"product_action_list,omitempty"`
	ProductListSource string            `json:"product_list_source,omitempty"`
	TransactionID     string            `json:"transaction_id,omitempty"`
	Affiliation       string            `json:"affiliation,omitempty"`
	TotalAmount       *float64          `json:"total_amount,omitempty"`
	TaxAmount         *float64          `json:"tax_amount,omitempty"`
	ShippingAmount    *float64          `json:"shipping_amount,omitempty"`
	CouponCode        string            `json:"coupon_code,omitempty"`
	Products          []Product         `json:"products,omitempty"`
}

// ProductActionType determines which action the user took
type ProductActionType string

// List of ProductActionType
const (
	AddToCartAction          ProductActionType = "add_to_cart"
	RemoveFromCartAction     ProductActionType = "remove_from_cart"
	CheckoutAction           ProductActionType = "checkout"
	CheckoutOptionAction     ProductActionType = "checkout_option"
	ClickAction              ProductActionType = "click"
	ViewDetailAction         ProductActionType = "view_detail"
	PurchaseAction           ProductActionType = "purchase"
	RefundAction             ProductActionType = "refund"
	AddToWishlistAction      ProductActionType = "add_to_wishlist"
	RemoveFromWishlistAction ProductActionType = "remove_from_wish_list"
)
