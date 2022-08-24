package main

import (
	"cloud.google.com/go/civil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMapEventsToAmplitudeRequests(t *testing.T) {

	orderCheckout := OrderCheckout{
		CustomerId:             12345,
		LoyaltyCardNo:          "LCN-12345",
		BusinessDate:           civil.DateTimeOf(time.Now().Round(0)),
		IsSelfCheckout:         "Y",
		StoreCode:              310,
		TillCode:               123123123,
		InvoiceNo:              12345123123123,
		SaleTotalQuantity:      10,
		SaleNetValue:           100.5,
		SaleTotalTaxValue:      10.05,
		SaleTotalDiscountValue: 0,
	}

	t.Run("Test should successfully map Order Checkouts to Amplitude events and return Amplitude "+
		"request for self checkouts",
		func(t *testing.T) {
			//Given
			orderCheckout.IsSelfCheckout = "Y"
			orderCheckoutsList := []OrderCheckout{orderCheckout}

			//When
			amplitudeRequest := MapEventsToAmplitudeRequests(orderCheckoutsList)

			//Then
			assert.NotNil(t, amplitudeRequest)
			assert.NotNil(t, amplitudeRequest.Events)
			assert.Equal(t, 1, len(amplitudeRequest.Events))
			orderCompletedEvent := amplitudeRequest.Events[0]
			orderCheckoutEventProperties := orderCompletedEvent.EventProperties.(AmplitudeCheckoutEventProperties)
			assert.Equal(t, OrderCheckoutCompleted, orderCompletedEvent.EventType)
			assert.Equal(t, orderCheckout.CustomerId, orderCompletedEvent.UserId)
			assert.NotNil(t, orderCheckoutEventProperties)
			assert.Equal(t, SelfCheckout, orderCheckoutEventProperties.CheckoutMethod)
			assert.Equal(t, orderCheckout.InvoiceNo, orderCheckoutEventProperties.InvoiceNo)
			assert.Equal(t, orderCheckout.CustomerId, orderCheckoutEventProperties.UserId)
			assert.Equal(t, orderCheckout.StoreCode, orderCheckoutEventProperties.StoreCode)
			assert.Equal(t, orderCheckout.TillCode, orderCheckoutEventProperties.TillCode)
			assert.Equal(t, orderCheckout.SaleTotalQuantity, orderCheckoutEventProperties.SaleTotalQuantity)
			assert.Equal(t, orderCheckout.SaleNetValue, orderCheckoutEventProperties.SaleNetValue)
			assert.Equal(t, orderCheckout.SaleTotalTaxValue, orderCheckoutEventProperties.SaleTotalTaxValue)
			assert.Equal(t, orderCheckout.SaleTotalDiscountValue, orderCheckoutEventProperties.SaleTotalDiscountValue)
			assert.Equal(t, Offline, orderCheckoutEventProperties.Vertical)
			assert.Equal(t, orderCheckout.BusinessDate.Date.String(), orderCheckoutEventProperties.BusinessDate)
		})

	t.Run("Test should successfully map Order Checkouts to Amplitude events and return Amplitude "+
		"request for cashier checkouts",
		func(t *testing.T) {
			//Given
			orderCheckout.IsSelfCheckout = "N"
			orderCheckoutsList := []OrderCheckout{orderCheckout}

			//When
			amplitudeRequest := MapEventsToAmplitudeRequests(orderCheckoutsList)

			//Then
			assert.NotNil(t, amplitudeRequest)
			assert.NotNil(t, amplitudeRequest.Events)
			assert.Equal(t, 1, len(amplitudeRequest.Events))
			orderCompletedEvent := amplitudeRequest.Events[0]
			orderCheckoutEventProperties := orderCompletedEvent.EventProperties.(AmplitudeCheckoutEventProperties)
			assert.Equal(t, OrderCheckoutCompleted, orderCompletedEvent.EventType)
			assert.Equal(t, orderCheckout.CustomerId, orderCompletedEvent.UserId)
			assert.NotNil(t, orderCheckoutEventProperties)
			assert.Equal(t, CashierCheckout, orderCheckoutEventProperties.CheckoutMethod)
			assert.Equal(t, orderCheckout.InvoiceNo, orderCheckoutEventProperties.InvoiceNo)
			assert.Equal(t, orderCheckout.CustomerId, orderCheckoutEventProperties.UserId)
			assert.Equal(t, orderCheckout.StoreCode, orderCheckoutEventProperties.StoreCode)
			assert.Equal(t, orderCheckout.TillCode, orderCheckoutEventProperties.TillCode)
			assert.Equal(t, orderCheckout.SaleTotalQuantity, orderCheckoutEventProperties.SaleTotalQuantity)
			assert.Equal(t, orderCheckout.SaleNetValue, orderCheckoutEventProperties.SaleNetValue)
			assert.Equal(t, orderCheckout.SaleTotalTaxValue, orderCheckoutEventProperties.SaleTotalTaxValue)
			assert.Equal(t, orderCheckout.SaleTotalDiscountValue, orderCheckoutEventProperties.SaleTotalDiscountValue)
			assert.Equal(t, Offline, orderCheckoutEventProperties.Vertical)
			assert.Equal(t, orderCheckout.BusinessDate.Date.String(), orderCheckoutEventProperties.BusinessDate)
		})
}
