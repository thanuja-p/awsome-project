package main

//var ApiKey = os.Getenv("AMPLITUDE_API_KEY")
var ApiKey = "d5f5ee124684abaecc85d3ec26c6a0ef"

func mapCheckoutToAmplitudeEvent(orderCheckout OrderCheckout) AmplitudeEvent {
	//customerId := "123000"
	//userId := strconv.FormatInt(orderCheckout.CustomerId, 10)
	checkoutEventAmplitudeProperties := AmplitudeCheckoutEventProperties{
		BusinessDate:           orderCheckout.BusinessDate.Date.String(),
		StoreCode:              orderCheckout.StoreCode,
		TillCode:               orderCheckout.TillCode,
		InvoiceNo:              orderCheckout.InvoiceNo,
		LoyaltyCardNo:          orderCheckout.LoyaltyCardNo,
		SaleTotalQuantity:      orderCheckout.SaleTotalQuantity,
		SaleNetValue:           orderCheckout.SaleNetValue,
		SaleTotalTaxValue:      orderCheckout.SaleTotalTaxValue,
		SaleTotalDiscountValue: orderCheckout.SaleTotalDiscountValue,
		UserId:                 orderCheckout.CustomerId,
		Vertical:               Offline,
		CheckoutMethod:         getCheckoutMethod(orderCheckout.IsSelfCheckout),
	}

	return AmplitudeEvent{
		UserId:          orderCheckout.CustomerId,
		EventType:       OrderCheckoutCompleted,
		EventProperties: checkoutEventAmplitudeProperties,
	}
}

func getCheckoutMethod(isSelfCheckout string) CheckoutMethod {
	if isSelfCheckout == Yes {
		return SelfCheckout
	}
	return CashierCheckout
}

func MapEventsToAmplitudeRequests(checkouts []OrderCheckout) *AmplitudeRequest {
	var checkoutAmplitudeEvents []AmplitudeEvent
	for i := 0; i < len(checkouts); i++ {
		amplitudeEvent := mapCheckoutToAmplitudeEvent(checkouts[i])
		checkoutAmplitudeEvents = append(checkoutAmplitudeEvents, amplitudeEvent)
	}
	return &AmplitudeRequest{
		APIKey: ApiKey,
		Events: checkoutAmplitudeEvents,
	}
}
