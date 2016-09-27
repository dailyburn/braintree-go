package braintree

import "github.com/dailyburn/braintree-go/nullable"

type DiscountList struct {
	XMLName   string     `xml:"discounts"`
	Discounts []Discount `xml:"discount"`
}

type Discount struct {
	XMLName string `xml:"discount"`
	Modification
}

type DiscountItem struct {
	Id                    string              `xml:"id,omitempty"`
	Name                  string              `xml:"name,omitempty"`
	NeverExpires          bool                `xml:"never-expires,omitempty"`
	Amount                *Decimal            `xml:"amount,omitempty"`
	CurrentBillingCycle   int                 `xml:"current-billing-cycle,omitempty"`
	NumberOfBillingCycles *nullable.NullInt64 `xml:"number-of-billing-cycles,omitempty"`
	Quantity              int                 `xml:"quantity,omitempty"`
}

type DiscountResults struct {
	Discount []DiscountItem `xml:"discount,omitempty"`
}
