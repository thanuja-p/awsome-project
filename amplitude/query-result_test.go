package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChuckArray(t *testing.T) {

	orderCheckout := OrderCheckout{
		CustomerId:             12345,
		LoyaltyCardNo:          "LCN-12345",
		IsSelfCheckout:         "Y",
		StoreCode:              310,
		TillCode:               123123123,
		InvoiceNo:              12345123123123,
		SaleTotalQuantity:      10,
		SaleNetValue:           100.5,
		SaleTotalTaxValue:      10.05,
		SaleTotalDiscountValue: 0,
	}

	t.Run("Test should chunk Order Checkouts list",
		func(t *testing.T) {
			//Given
			orderCheckoutsList := []OrderCheckout{orderCheckout, orderCheckout, orderCheckout}

			//When
			chunkedArray := ChuckArray(orderCheckoutsList, 2)

			//Then
			assert.NotNil(t, chunkedArray)
			assert.Equal(t, 2, len(chunkedArray))
			assert.Equal(t, 2, len(chunkedArray[0]))
			assert.Equal(t, 1, len(chunkedArray[1]))
		})
}

//func TestIterateQueryResults(t *testing.T) {
//
//	type RowIteratorMock struct {
//		// NextFunc mocks the Next method.
//		NextFunc func(in1 interface{}) error
//
//		// PageInfoFunc mocks the PageInfo method.
//		PageInfoFunc func() *iterator.PageInfo
//
//		// SchemaFunc mocks the Schema method.
//		SchemaFunc func() bigquery.Schema
//
//		// SetStartIndexFunc mocks the SetStartIndex method.
//		SetStartIndexFunc func(in1 uint64)
//
//		// TotalRowsFunc mocks the TotalRows method.
//		TotalRowsFunc func() uint64
//		// contains filtered or unexported fields
//	}
//
//t.Run("Test should iterate query results and return Order checkout array",
//	func(t *testing.T) {
//		//Given
//		ctx := context.Background()
//		it := *bigquery.RowIterator{
//			TotalRows: 10,
//		}
//
//		type RowIterator interface {
//			SetStartIndex(uint64)
//			Schema() bigquery.Schema
//			TotalRows() uint64
//			Next(interface{}) error
//			PageInfo() *iterator.PageInfo
//			// contains filtered or unexported methods
//		}
//
//		//When
//		chunkedArray := iterateQueryResults(ctx, it)
//
//		//Then
//		assert.NotNil(t, chunkedArray)
//	})
//}
