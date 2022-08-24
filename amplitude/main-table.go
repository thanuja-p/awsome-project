package main

//import (
//	"cloud.google.com/go/bigquery"
//	"context"
//	"fmt"
//	"google.golang.org/api/iterator"
//)
//
//type OrderCheckout struct {
//	CustomerId    int64  `bigquery:"customerId"`
//	LoyaltyCardNo string `bigquery:"loyaltyCardNo"`
//}
//
//func main() {
//	fmt.Println("Hello world")
//
//	ctx := context.Background()
//
//	list := GetOrderCheckouts(ctx)
//	request := MapEventsToAmplitudeRequests(list)
//	//logging.WithCtx(ctx).Printf("Going to trigger OG Amplitude request for customerId: %v", order.CustomerId)
//	//fmt.Println(order.CustomerId)
//	SendBatchEvent(ctx, request)
//	logResults(list)
//	fmt.Println("DONE!")
//}
//
//func GetOrderCheckouts(ctx context.Context) []OrderCheckout {
//	client := getClient(ctx)
//	queryString := "SELECT customer_id as customerId, " +
//		"loyalty_card_no as loyaltyCardNo " +
//		"FROM `fairprice-bigquery.cdm_grocery.dim_fpon_customer_mapping` LIMIT 2"
//	it := getResultIterator(ctx, client, queryString)
//	return iterateQueryResults(it)
//}
//
//// queryBasic demonstrates issuing a getResultIterator and reading results.
////func queryBasic(w io.Writer, projectID string) error {
//func getClient(ctx context.Context) *bigquery.Client {
//	//projectID := "ne-fprt-data-cloud-production"
//	projectID := "fairprice-bigquery"
//
//	client, err := bigquery.NewClient(ctx, projectID)
//
//	if err != nil {
//		fmt.Errorf("bigquery.NewClient: %v", err)
//	}
//	defer client.Close()
//	return client
//}
//
//func getResultIterator(ctx context.Context, client *bigquery.Client, queryString string) *bigquery.RowIterator {
//
//	// [START bigquery_simple_app_query]
//	getResultIterator := client.Query(queryString)
//	iterator, err := getResultIterator.Read(ctx)
//	if err != nil {
//		fmt.Errorf("bigquery.ReadQuery: %v", err)
//	}
//	return iterator
//}
//
//// printResults prints results from a getResultIterator to the Stack Overflow public dataset.
//func iterateQueryResults(iter *bigquery.RowIterator) []OrderCheckout {
//	var loyaltyCardsList []OrderCheckout
//	for {
//		//var row []bigquery.Value
//		//var row []string
//		var row OrderCheckout
//		err := iter.Next(&row)
//		if err == iterator.Done {
//			fmt.Println("done-itr")
//			break
//		}
//		if err != nil {
//			fmt.Errorf("error iterating through results: %v", err)
//		}
//		loyaltyCardsList = append(loyaltyCardsList, row)
//	}
//	return loyaltyCardsList
//}
//
//func logResults(cards []OrderCheckout) {
//	// using for loop
//	for i := 0; i < len(cards); i++ {
//		fmt.Printf("Entry: %d \t customerId: %d \t loyalty card no: %s\n", i+1, cards[i].CustomerId, cards[i].LoyaltyCardNo)
//	}
//}
