package braintree

import "github.com/dailyburn/braintree-go/nullable"

type AddOnList struct {
	XMLName string  `xml:"add-ons"`
	AddOns  []AddOn `xml:"add-on"`
}

type AddOn struct {
	XMLName string `xml:"add-on"`
	Modification
}

type AddOnItem struct {
	Amount                *Decimal            `xml:"amount,omitempty"`
	CurrentBillingCycle   int                 `xml:"current-billing-cycle,omitempty"`
	Id                    string              `xml:"id,omitempty"`
	Name                  string              `xml:"name,omitempty"`
	NeverExpires          bool                `xml:"never-expires,omitempty"`
	NumberOfBillingCycles *nullable.NullInt64 `xml:"number-of-billing-cycles,omitempty"`
	Quantity              int                 `xml:"quantity,omitempty"`
}

type AddOnResults struct {
	AddOn []AddOnItem `xml:"add-on,omitempty"`
}
