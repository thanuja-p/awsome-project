package main

type AmplitudeEventType string
type CheckoutMethod string

const (
	OrderCheckoutCompleted AmplitudeEventType = "order checkout completed"
	Offline                string             = "Offline"
	Yes                    string             = "Y"
)

const (
	CashierCheckout CheckoutMethod = "Cashier checkout"
	SelfCheckout    CheckoutMethod = "Self checkout"
)

type AmplitudeCheckoutEventProperties struct {
	UserId                 int64          `json:"user id,omitempty"`
	Vertical               string         `json:"vertical,omitempty"`
	CheckoutMethod         CheckoutMethod `json:"checkout method,omitempty"`
	LoyaltyCardNo          string         `json:"loyalty card no,omitempty"`
	BusinessDate           string         `json:"business date,omitempty"`
	StoreCode              int64          `json:"store code,omitempty"`
	TillCode               int64          `json:"till code,omitempty"`
	InvoiceNo              int64          `json:"invoice no,omitempty"`
	SaleTotalQuantity      int64          `json:"sale total quantity,omitempty"`
	SaleNetValue           float64        `json:"sale net value,omitempty"`
	SaleTotalTaxValue      float64        `json:"sale total tax value,omitempty"`
	SaleTotalDiscountValue float64        `json:"sale total discount Value,omitempty"`
}

type AmplitudeEvent struct {
	UserId          int64              `json:"user_id,omitempty"`
	EventType       AmplitudeEventType `json:"event_type,omitempty"`
	EventProperties interface{}        `json:"event_properties"`
}

type AmplitudeRequest struct {
	APIKey string           `json:"api_key"`
	Events []AmplitudeEvent `json:"events"`
}
