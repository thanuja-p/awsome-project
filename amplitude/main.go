package main

import (
	"cloud.google.com/go/bigquery"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
)

func mainOld() {
	queryBasic()
}

//
//type OrderCheckout struct {
//	customerId    int64  `bigquery:"customerId"`
//	loyaltyCardNo string `bigquery:"loyaltyCardNo"`
//}

func main1() {
	fmt.Println("Hello world")
	//ctx := context.Background()
	//request, order := MapOGOrderPlacementToAmplitudeRequests(getOrder())
	//logging.WithCtx(ctx).Printf("Going to trigger OG Amplitude request for customerId: %v", order.CustomerId)
	//fmt.Println(order.CustomerId)
	////SendEvent(ctx, request)
	//SendBatchEvent(ctx, request)
	//dataset()
	queryBasic()
	//projectID := "ne-fprt-data-cloud-production"
	//datasetID := "fp_sale_head"
	//entity := "thanuja.perera@ntucenterprise.sg"
	//projectID := "fairprice-bigquery"
	//datasetID := "dim_fpon_customer_mapping"
	//revokeDatasetAccess(projectID, datasetID, entity)
	fmt.Println("DONE!")
}

func revokeDatasetAccess(projectID, datasetID, entity string) error {
	// projectID := "my-project-id"
	// datasetID := "mydataset"
	// entity := "user@mydomain.com"
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	ds := client.Dataset(datasetID)

	meta, err := ds.Metadata(ctx)
	if err != nil {
		return err
	}

	fmt.Println("meta.Access")
	fmt.Println(meta.Access)

	var newAccessList []*bigquery.AccessEntry
	for _, entry := range meta.Access {
		if entry.Entity != entity {
			newAccessList = append(newAccessList, entry)
		}
	}

	fmt.Println("newAccessList")
	fmt.Println(newAccessList)

	// Only proceed with update if something in the access list was removed.
	// Additionally, we use the ETag from the initial metadata to ensure no
	// other changes were made to the access list in the interim.
	//if len(newAccessList) < len(meta.Access) {
	//
	//	update := bigquery.DatasetMetadataToUpdate{
	//		Access: newAccessList,
	//	}
	//	if _, err := ds.Update(ctx, update, meta.ETag); err != nil {
	//		return err
	//	}
	//}
	return nil
}

// queryBasic demonstrates issuing a getResultIterator and reading results.
//func queryBasic(w io.Writer, projectID string) error {
func queryBasic() error {

	projectID := "fairprice-bigquery"
	//projectID := "ne-fprt-data-cloud-production"
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %v", err)
	}

	defer client.Close()
	q := client.Query(
		"select * from ne-central-data-cloud-prod.data_asset.scv_cust_master_sale_bridge where txn_type is not null LIMIT 5")
	//q := client.Query(
	//	"SELECT * FROM `fairprice-bigquery.cdm_grocery.ads_amplitude_customer_metrics` LIMIT 2")

	job, err := q.Run(ctx)

	if err != nil {
		return err
	}
	status, err := job.Wait(ctx)
	if err != nil {
		return err
	}
	if err := status.Err(); err != nil {
		return err
	}
	it, err := job.Read(ctx)
	//fmt.Println("total rows: " + string(it.TotalRows))
	for {
		var row []bigquery.Value
		//var row OrderCheckout
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println("row")
		fmt.Println(row)

	}
	return nil
}

func dataset() {

	db, err := sql.Open("bigquery", "bigquery://ne-fprt-data-cloud-production/fp_sale_head")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	SQL := `WITH races AS (
		  SELECT "800M" AS race,
		    [STRUCT("Ben" as name, [23.4, 26.3] as splits), 
		 	 STRUCT("Frank" as name, [23.4, 26.3] as splits)
			]
		       AS participants)
		SELECT
		  race,
		  participant
		FROM races r
		CROSS JOIN UNNEST(r.participants) as participant`

	rows, err := db.QueryContext(context.Background(), SQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		//var race string
		//var participant Participant
		//err = rows.Scan(&race, &participant)
		//fmt.Printf("fetched: %v %+v\n", race, participant)
		//if err != nil {
		//	log.Fatal(err)
		//}
	}
}

func getOrder() map[string]interface{} {
	orderJson := `{
       "LoyaltyCardNo": "lcn-848940030",
       "cartId" : "848940030",
	   "amount":"63.00",
	   "cancellation":{
		  "cancellationReason":null,
		  "chargeAmount":0,
		  "createdAt":null,
		  "isCancellable":true
	   },
	   "couponCode":null,
	   "couponDiscount":"0.00",
	   "createdAt":"2022-06-16 15:12:01",
	   "currentServiceFee":"0.00",
	   "customer":{
		  "emails":[
			 {
				"email":"indunil@ntucenterprise.sg"
			 }
		  ],
		  "id":6835103,
		  "image":"",
		  "joinedOn":"2022-05-23",
		  "joinedTime":"2022-05-23 16:17:20",
		  "phones":[
			 {
				"phone":"+6500000000"
			 }
		  ],
		  "preference":null,
		  "status":"ENABLED",
		  "uid":"140277438935832551",
		  "updatedAt":"2022-05-25 13:18:16"
	   },
	   "discount":"63.00000",
	   "edited":false,
	   "id":52286,
	   "invoiceAmount":"5.00",
	   "invoiceAmountBalance":"5.00",
	   "isSubstitutable":true,
	   "isSubstituted":false,
	   "is_cancellable":true,
	   "items":[
		  {
			 "bulkOrderThreshold":99999,
			 "clientItemId":"10983610",
			 "createdAt":"2019-04-27T04:16:03+08:00",
			 "description":"Value Pack",
			 "handlingDays":0,
			 "hasVariants":0,
			 "id":1174191,
			 "imagesExtra":null,
			 "languages":"",
			 "metaData":{
				"Country of Origin":"South Korea",
				"Dietary Attributes":[
				   "Trans-Fat Free"
				],
				"DisplayUnit":"200g (3 per pack)",
				"Fp Seller Name":"ORIENT FOODS PTE LTD",
				"Ingredients":"Wheat Flour, Water, Thickener (E1442, E405), Refined Salt, Lipase (E1104), Acidity Regulators (E260, E575, E350, E270, E325), Contains Gluten (Wheat Flour)",
				"LinkPoint Eligible":true,
				"SAP Product Name":"FAIRPRICE INSTANT JAPANESE FRESH UDON TRIPLE VALUE PACK 200G",
				"SAP SubClass":220550603,
				"Storage Information":"Store at room temperature",
				"Storage Type":"C2",
				"Unit Of Measurement":"EA",
				"Unit Of Weight":"200G",
				"Weight":"200 gm"
			 },
			 "name":"FairPrice Instant Fresh Japanese Udon ",
			 "offer":[
				{
				   "addToWallet":false,
				   "appliedCount":2,
				   "appliedPromoCode":null,
				   "clientId":"ONLINE-004-10983610",
				   "customerRedemptionLimit":null,
				   "description":"Buy 1 FairPrice Instant Fresh Japanese Udon  @ $0.15 Off",
				   "discount":0.3,
				   "endDeliveryDate":null,
				   "hasUniquePromocode":false,
				   "id":10493529,
				   "ignoreSegment":false,
				   "imageUrl":null,
				   "metaData":{
					  "promotionNumber":"2245958",
					  "sapReference":"ZSP"
				   },
				   "offerType":"BXATP",
				   "offerValidFrom":"2022-04-14 04:01:00",
				   "offerValidTill":"2023-01-01 04:00:00",
				   "orderType":[
					  "BULK",
					  "DELIVERY",
					  "PICKUP",
					  "B2B"
				   ],
				   "paymentType":null,
				   "promoCode":null,
				   "rule":{
					  "buy":{
						 "1174191":{
							"q":1
						 }
					  },
					  "total":{
						 "t":"ABSOLUTE_OFF",
						 "v":0.15
					  }
				   },
				   "ruleDetail":{
					  "buy":{
						 "discountType":"ABSOLUTE_OFF",
						 "discountValue":0.15,
						 "isCombo":false,
						 "minCartPrice":null,
						 "products":[
							{
							   "barcodes":[
								  "8888030026915"
							   ],
							   "brand":{
								  "clientId":"FAIRPRICE",
								  "description":null,
								  "id":28962,
								  "image":null,
								  "languages":"",
								  "logo":null,
								  "name":"Fairprice",
								  "productsCount":1531,
								  "slug":"fairprice-2",
								  "status":"ENABLED"
							   },
							   "bulkOrderThreshold":99999,
							   "clientItemId":"10983610",
							   "createdAt":"2019-04-27T04:16:03+08:00",
							   "description":"Value Pack",
							   "handlingDays":0,
							   "hasVariants":0,
							   "id":1174191,
							   "imagesExtra":null,
							   "languages":"",
							   "metaData":{
								  "Country of Origin":"South Korea",
								  "Dietary Attributes":[
									 "Trans-Fat Free"
								  ],
								  "DisplayUnit":"200g (3 per pack)",
								  "Fp Seller Name":"ORIENT FOODS PTE LTD",
								  "Ingredients":"Wheat Flour, Water, Thickener (E1442, E405), Refined Salt, Lipase (E1104), Acidity Regulators (E260, E575, E350, E270, E325), Contains Gluten (Wheat Flour)",
								  "LinkPoint Eligible":true,
								  "SAP Product Name":"FAIRPRICE INSTANT JAPANESE FRESH UDON TRIPLE VALUE PACK 200G",
								  "SAP SubClass":220550603,
								  "Storage Information":"Store at room temperature",
								  "Storage Type":"C2",
								  "Unit Of Measurement":"EA",
								  "Unit Of Weight":"200G",
								  "Weight":"200 gm"
							   },
							   "name":"FairPrice Instant Fresh Japanese Udon ",
							   "organizationId":"2",
							   "primaryCategory":{
								  "clientId":"13308",
								  "createdAt":null,
								  "deletedAt":null,
								  "description":null,
								  "id":289,
								  "image":"banners/web/category/FreshChilled_L3.jpg",
								  "languages":"",
								  "name":"Fresh \u0026 Chilled",
								  "parentCategory":{
									 "clientId":"58704",
									 "createdAt":null,
									 "deletedAt":null,
									 "description":null,
									 "id":2927,
									 "image":null,
									 "languages":"",
									 "name":"Noodles.",
									 "parentCategory":{
										"clientId":"12703",
										"createdAt":null,
										"deletedAt":null,
										"description":null,
										"id":537,
										"parentCategory":null,
										"productsCount":7,
										"slug":"baking-cooking",
										"status":"HIDDEN",
										"updatedAt":null,
										"updatedBy":"1"
									 },
									 "productsCount":0,
									 "slug":"noodles--1",
									 "status":"ENABLED",
									 "updatedAt":null,
									 "updatedBy":""
								  },
								  "productsCount":23,
								  "slug":"fresh-chilled",
								  "status":"ENABLED",
								  "updatedAt":null,
								  "updatedBy":""
							   },
							   "quantity":1,
							   "secondaryCategoryIds":[
								  435,
								  609,
								  683,
								  2115
							   ],
							   "slug":"fairprice-instant-fresh-japanese-udon-10983610",
							   "soldByWeight":0,
							   "status":"ENABLED",
							   "stockOverride":{
								  "maxPurchasableStock":null,
								  "stockBuffer":5,
								  "storeBulkOrderThreshold":5
							   },
							   "tagIds":[
								  4,
								  122,
								  123,
								  124
							   ],
							   "updatedBy":"",
							   "variants":null,
							   "vendor_code":0
							}
						 ]
					  },
					  "elementGroup":null
				   },
				   "stackable":false,
				   "startDeliveryDate":null,
				   "status":"ENABLED",
				   "storeId":"4,8,188,204,213,235,237,261,262,267,270,324,358,389,395,396,397,398,399,400,401,402,403,404,405",
				   "totalRedemption":null,
				   "updatedAt":"2022-04-14T18:12:33+08:00"
				}
			 ],
			 "orderDetails":{
				"deliveredQuantity":null,
				"discount":"0.15000",
				"disputeQuantity":null,
				"isPickupPending":0,
				"metaData":{
				   "Country of Origin":"South Korea",
				   "Dietary Attributes":[
					  "Trans-Fat Free"
				   ],
				   "DisplayUnit":"200g (3 per pack)",
				   "Fp Seller Name":"ORIENT FOODS PTE LTD",
				   "Ingredients":"Wheat Flour, Water, Thickener (E1442, E405), Refined Salt, Lipase (E1104), Acidity Regulators (E260, E575, E350, E270, E325), Contains Gluten (Wheat Flour)",
				   "LinkPoint Eligible":true,
				   "SAP Product Name":"FAIRPRICE INSTANT JAPANESE FRESH UDON TRIPLE VALUE PACK 200G",
				   "SAP SubClass":220550603,
				   "Storage Information":"Store at room temperature",
				   "Storage Type":"C2",
				   "Unit Of Measurement":"EA",
				   "Unit Of Weight":"200G",
				   "Weight":"200 gm"
				},
				"mrp":"2.6000000000",
				"orderItemId":127286,
				"orderedQuantity":"2.00000",
				"status":"PENDING",
				"substitutedItemId":null,
				"tax":null,
				"tripId":null
			 },
			 "organizationId":"2",
			 "primaryCategory":{
				"clientId":"13308",
				"createdAt":null,
				"deletedAt":null,
				"description":null,
				"id":289,
				"image":"banners/web/category/FreshChilled_L3.jpg",
				"languages":"",
				"name":"Fresh \u0026 Chilled",
				"parentCategory":{
				   "clientId":"58704",
				   "createdAt":null,
				   "deletedAt":null,
				   "description":null,
				   "id":2927,
				   "image":null,
				   "languages":"",
				   "name":"Noodles.",
				   "parentCategory":{
					  "clientId":"12703",
					  "createdAt":null,
					  "deletedAt":null,
					  "description":null,
					  "id":537,
					  "image":"banners/web/category/BakingCooking_L1.jpg",
					  "languages":"",
					  "name":"Rice, Noodles \u0026 Cooking Ingredients",
					  "parentCategory":null,
					  "productsCount":7,
					  "slug":"baking-cooking",
					  "status":"HIDDEN",
					  "updatedAt":null,
					  "updatedBy":"1"
				   },
				   "productsCount":0,
				   "slug":"noodles--1",
				   "status":"ENABLED",
				   "updatedAt":null,
				   "updatedBy":""
				},
				"productsCount":23,
				"slug":"fresh-chilled",
				"status":"ENABLED",
				"updatedAt":null,
				"updatedBy":""
			 },
			 "soldByWeight":0,
			 "status":"ENABLED",
			 "tagIds":[
				4,
				122,
				123,
				124
			 ],
			 "updatedBy":"",
			 "variants":null,
			 "vendor_code":0
		  }
	   ],
	   "linkpoints":null,
	   "metaData":{
		  "dbMemberLinkStatus":0,
		  "isDBMember":true,
          "redemptionLinkPoints": "399.00"
	   },
	   "offers":[
		  {
			 "addToWallet":false,
			 "appliedPromoCode":"CART80",
			 "clientId":null,
			 "customerRedemptionLimit":null,
			 "description":"$80 Online Voucher",
			 "discount":62.7,
			 "discountType":"ORDER",
			 "endDeliveryDate":null,
			 "hasUniquePromocode":false,
			 "id":10498630,
			 "ignoreSegment":false,
			 "imageUrl":"",
			 "metaData":{
				"sapReference":"ZKP3 80 Dollars Off"
			 },
			 "offerType":"SFXGCD",
			 "offerValidFrom":"2021-12-15 02:59:03",
			 "offerValidTill":"2023-07-31 11:59:03",
			 "orderType":[
				"BULK",
				"DELIVERY",
				"PICKUP",
				"B2B"
			 ],
			 "paymentType":null,
			 "promoCode":[
				"CART80"
			 ],
			 "rule":[
				{
				   "cartDiscount":{
					  "t":"ABSOLUTE_OFF",
					  "v":80
				   },
				   "cartPrice":0.01
				}
			 ],
			 "ruleDetail":{
				"elementGroup":[
				   {
					  "cartDiscount":{
						 "type":"ABSOLUTE_OFF",
						 "value":80
					  },
					  "maxDiscount":null,
					  "minCartPrice":0.01
				   }
				]
			 },
			 "stackable":true,
			 "startDeliveryDate":null,
			 "status":"ENABLED",
			 "storeId":"4,8,204,356,358,396",
			 "totalRedemption":null,
			 "updatedAt":"2022-04-21T18:44:57+08:00",
			 "userSet":{
				"data":[
				   {
					  "id":"AllCustomers"
				   }
				],
				"type":"SEGMENTS"
			 }
		  }
	   ],
	   "packingDetails":[
		  
	   ],
	   "paidAmount":"0",
	   "pay":{
		  "amount":"5.00",
		  "bankTransactionId":"016153570198200",
		  "completedAt":null,
		  "createdAt":"2022-06-16 15:12:03",
		  "gatewayTransactionId":"6553635228876194403954",
		  "id":48070,
		  "metaData":{
			 "acquirerId":"DP2022061600000000191182",
			 "orderType":"DELIVERY",
			 "reconciliationId":"",
			 "systemTraceAuditNumber":"135010"
		  },
		  "mode":"ONLINE",
		  "paymentService":"CYBERSOURCE",
		  "paymentServiceId":"191182",
		  "status":"PENDING",
		  "transactionId":"70161338-ONLINE-1655363523"
	   },
	   "payment":[
		  {
			 "amount":"5.00",
			 "bankTransactionId":"016153570198200",
			 "completedAt":null,
			 "createdAt":"2022-06-16 15:12:03",
			 "gatewayTransactionId":"6553635228876194403954",
			 "id":48070,
			 "metaData":{
				"acquirerId":"DP2022061600000000191182",
				"cardNumber":"498843xxxxxx4305",
				"cardType":"VISA",
				"orderType":"DELIVERY",
				"reconciliationId":"",
				"systemTraceAuditNumber":"135010",
				"type":"CREDIT"
			 },
			 "mode":"ONLINE",
			 "paymentServiceId":"191182",
			 "status":"PENDING",
			 "transactionId":"70161338-ONLINE-1655363523"
		  }
	   ],
	   "paymentStatus":"PENDING",
	   "pendingAmount":"5.00",
	   "pickupLocationId":null,
	   "placedFrom":"DESKTOP_NA_WEB",
	   "placedOn":"2022-06-16",
	   "preferredDate":"2022-06-18",
	   "preorder":false,
	   "quickbuyStatus":null,
	   "referenceNumber":"70161338",
	   "refund":[
		  
	   ],
	   "refundAmount":"0.00",
	   "seller":null,
	   "shipping":"5.00",
	   "slotEndTime":"22:00:00",
	   "slotStartTime":"20:00:00",
	   "slotSurcharge":"0",
	   "slotType":"STANDARD",
	   "standardServiceFee":"3.99",
	   "standardShippingAmount":"5.00",
	   "status":"PENDING",
	   "store":{
		  "address":"50 Jurong Gateway Road #B1-21/22 \u0026 #B3-01, Singapore 608549",
		  "b2bTierSetupId":null,
		  "businessHours":null,
		  "clientId":"469",
		  "hasClickCollect":false,
		  "hasDeliveryHub":true,
		  "hasPicking":true,
		  "hasSelfCheckout":false,
		  "id":358,
		  "languages":"",
		  "latitude":1.3328361,
		  "longitude":103.7430889,
		  "metaData":{
			 "B2B Delivery Days":0,
			 "B2B Enabled":false,
			 "Batch Order Enabled":true,
			 "Batch Order Percentage":100,
			 "Delivery Days":3,
			 "Delivery type":"",
			 "Email":"",
			 "FFS Exclusive":false,
			 "Fax":"",
			 "Fresh Report Picking":false,
			 "Is SnG Store":true,
			 "Link ID":"",
			 "Past Stock Reservation Days":2,
			 "Phone":"",
			 "SAP Code":"469",
			 "SAP Inventory Download":false,
			 "SAP Posting Disabled":false,
			 "SAP Zone Code":"HJEM",
			 "Search And Browse Enabled":false,
			 "SnG Check-in Blocked":false,
			 "SnG Geo Checkin Radius":"",
			 "SnG Operational Hours":"",
			 "Store Code":"HJEM",
			 "Store Format":"FairPrice",
			 "WCS Id":"",
			 "Zone Picking Enabled":false
		  },
		  "name":"Hyper JEM (FFS)",
		  "pickingStoreId":null,
		  "status":"ENABLED",
		  "tierSetupId":64
	   },
	   "surcharge":"0",
	   "type":{
		  "id":2,
		  "name":"DELIVERY"
	   },
	   "verificationStatus":null,
	   "vouchers":null
	}`
	var order map[string]interface{}
	json.Unmarshal([]byte(orderJson), &order)
	return order
}

////q.DefaultProjectID = projectID
////q.DefaultDatasetID = "fp_sale_head"
//fmt.Println("dataset: " + q.DefaultDatasetID)
//fmt.Println("projectId: " + q.DefaultProjectID)
////// Location must match that of the dataset(s) referenced in the getResultIterator.
////q.Location = "US"
////// Run the getResultIterator and logResults results when the getResultIterator job is completed.
