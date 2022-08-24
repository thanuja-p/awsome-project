package main

func getQuery() string {

	queryString := `WITH
  sales AS (
  SELECT
    coalesce(proof_of_identity,
      loyalty_card_no) AS card_no,
    auto_self_check_out_flag,
    business_date,
    store_code,
    till_code,
    invoice_no,
    sale_tot_qty,
    sale_net_val,
    sale_tot_tax_val,
    sale_tot_disc_val
  FROM
    ne-fprt-data-cloud-production.fp.fp_sale_head
  WHERE
    coalesce(proof_of_identity,
      loyalty_card_no) IS NOT NULL
    AND TIMESTAMP(invoice_date) > TIMESTAMP_ADD(CURRENT_TIMESTAMP(), INTERVAL -1 MINUTE)),
  customer AS (
  SELECT
    card_no,
    fpon_customer_id AS customer_id
  FROM
    fairprice-bigquery.cdm_grocery.ads_amplitude_customer_metrics)
SELECT
  c.card_no,
  c.customer_id,
  s.auto_self_check_out_flag,
  s.business_date,
  s.store_code,
  s.till_code,
  s.invoice_no,
  CAST(s.sale_tot_qty as INT64) as sale_tot_qty,
  CAST(s.sale_net_val as FLOAT64) as sale_net_val,
  CAST(s.sale_tot_tax_val as FLOAT64) as sale_tot_tax_val,
  CAST(s.sale_tot_disc_val as FLOAT64) as sale_tot_disc_val
FROM
  sales AS s
JOIN
  customer AS c
ON
  s.card_no= c.card_no`
	return queryString
}
