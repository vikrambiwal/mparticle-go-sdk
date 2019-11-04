 
package events
// ProductImpression struct for ProductImpression
type ProductImpression struct {
	ProductImpressionList string `json:"product_impression_list"`
	Products []Product `json:"products"`
}
