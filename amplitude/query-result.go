package main

import (
	"bitbucket.org/ntuclink/devex-sdk-go/logging/v2"
	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/civil"
	"context"
	"google.golang.org/api/iterator"
)

const (
	projectID string = "fairprice-bigquery"
)

type OrderCheckout struct {
	CustomerId             int64          `bigquery:"customer_id"`
	LoyaltyCardNo          string         `bigquery:"card_no"`
	BusinessDate           civil.DateTime `bigquery:"business_date"`
	IsSelfCheckout         string         `bigquery:"auto_self_check_out_flag"`
	StoreCode              int64          `bigquery:"store_code"`
	TillCode               int64          `bigquery:"till_code"`
	InvoiceNo              int64          `bigquery:"invoice_no"`
	SaleTotalQuantity      int64          `bigquery:"sale_tot_qty"`
	SaleNetValue           float64        `bigquery:"sale_net_val"`
	SaleTotalTaxValue      float64        `bigquery:"sale_tot_tax_val"`
	SaleTotalDiscountValue float64        `bigquery:"sale_tot_disc_val"`
}

func GetOrderCheckouts(ctx context.Context) []OrderCheckout {
	client := getClient(ctx, projectID)
	queryString := getQuery()
	it := getResultIterator(ctx, client, queryString)
	err := client.Close()
	if err != nil {
		logging.WithCtx(ctx).Errorf("Failed to close the big query client: %v", err)
	}
	return iterateQueryResults(ctx, it)
}

func getClient(ctx context.Context, projectID string) *bigquery.Client {
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		logging.WithCtx(ctx).Errorf("Failed to crate big query new client: %v", err)
	}
	return client
}

func getResultIterator(ctx context.Context, client *bigquery.Client, queryString string) *bigquery.RowIterator {
	query := client.Query(queryString)
	it, err := query.Read(ctx)
	if err != nil {
		logging.WithCtx(ctx).Errorf("Failed to read query: %v", err)
	}
	return it
}

func iterateQueryResults(ctx context.Context, iter *bigquery.RowIterator) []OrderCheckout {
	var loyaltyCardsList []OrderCheckout
	for {
		var row OrderCheckout
		err := iter.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			logging.WithCtx(ctx).Errorf("Failed to iterate query results: %v", err)
		}
		loyaltyCardsList = append(loyaltyCardsList, row)
	}
	return loyaltyCardsList
}

func ChuckArray(fullArray []OrderCheckout, chunkSize int) [][]OrderCheckout {
	var chunks [][]OrderCheckout
	for i := 0; i < len(fullArray); i += chunkSize {
		end := i + chunkSize
		if end > len(fullArray) {
			end = len(fullArray)
		}
		chunks = append(chunks, fullArray[i:end])
	}
	return chunks
}
