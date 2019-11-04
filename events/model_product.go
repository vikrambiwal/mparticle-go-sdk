package events

// Product struct for Product
type Product struct {
	ID                 string            `json:"id,omitempty"`
	Name               string            `json:"name,omitempty"`
	Brand              string            `json:"brand,omitempty"`
	Category           string            `json:"category,omitempty"`
	Variant            string            `json:"variant,omitempty"`
	Position           *int64            `json:"position,omitempty"`
	Price              *float64          `json:"price,omitempty"`
	Quantity           *int32            `json:"quantity,omitempty"`
	CouponCode         string            `json:"coupon_code,omitempty"`
	AddedToCartTimeMS  int64             `json:"added_to_cart_time_ms,omitempty"`
	TotalProductAmount *float64          `json:"total_product_amount,omitempty"`
	CustomAttributes   map[string]string `json:"custom_attributes,omitempty"`
}
