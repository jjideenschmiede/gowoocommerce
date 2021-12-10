//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of gowoocommerce.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package gowoocommerce

import (
	"encoding/json"
	"fmt"
)

// OrdersReturn is to decode the json data
type OrdersReturn struct {
	Id               int    `json:"id"`
	ParentId         int    `json:"parent_id"`
	Status           string `json:"status"`
	Currency         string `json:"currency"`
	Version          string `json:"version"`
	PricesIncludeTax bool   `json:"prices_include_tax"`
	DateCreated      string `json:"date_created"`
	DateModified     string `json:"date_modified"`
	DiscountTotal    string `json:"discount_total"`
	DiscountTax      string `json:"discount_tax"`
	ShippingTotal    string `json:"shipping_total"`
	ShippingTax      string `json:"shipping_tax"`
	CartTax          string `json:"cart_tax"`
	Total            string `json:"total"`
	TotalTax         string `json:"total_tax"`
	CustomerId       int    `json:"customer_id"`
	OrderKey         string `json:"order_key"`
	Billing          struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Company   string `json:"company"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Postcode  string `json:"postcode"`
		Country   string `json:"country"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	} `json:"billing"`
	Shipping struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Company   string `json:"company"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Postcode  string `json:"postcode"`
		Country   string `json:"country"`
		Phone     string `json:"phone"`
	} `json:"shipping"`
	PaymentMethod      string        `json:"payment_method"`
	PaymentMethodTitle string        `json:"payment_method_title"`
	TransactionId      string        `json:"transaction_id"`
	CustomerIpAddress  string        `json:"customer_ip_address"`
	CustomerUserAgent  string        `json:"customer_user_agent"`
	CreatedVia         string        `json:"created_via"`
	CustomerNote       string        `json:"customer_note"`
	DateCompleted      interface{}   `json:"date_completed"`
	DatePaid           interface{}   `json:"date_paid"`
	CartHash           string        `json:"cart_hash"`
	Number             string        `json:"number"`
	MetaData           []interface{} `json:"meta_data"`
	LineItems          []struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		ProductId   int    `json:"product_id"`
		VariationId int    `json:"variation_id"`
		Quantity    int    `json:"quantity"`
		TaxClass    string `json:"tax_class"`
		Subtotal    string `json:"subtotal"`
		SubtotalTax string `json:"subtotal_tax"`
		Total       string `json:"total"`
		TotalTax    string `json:"total_tax"`
		Taxes       []struct {
			Id       int    `json:"id"`
			Total    string `json:"total"`
			Subtotal string `json:"subtotal"`
		} `json:"taxes"`
		MetaData   []interface{} `json:"meta_data"`
		Sku        string        `json:"sku"`
		Price      float64       `json:"price"`
		ParentName interface{}   `json:"parent_name"`
	} `json:"line_items"`
	TaxLines []struct {
		Id               int           `json:"id"`
		RateCode         string        `json:"rate_code"`
		RateId           int           `json:"rate_id"`
		Label            string        `json:"label"`
		Compound         bool          `json:"compound"`
		TaxTotal         string        `json:"tax_total"`
		ShippingTaxTotal string        `json:"shipping_tax_total"`
		RatePercent      int           `json:"rate_percent"`
		MetaData         []interface{} `json:"meta_data"`
	} `json:"tax_lines"`
	ShippingLines []struct {
		Id          int           `json:"id"`
		MethodTitle string        `json:"method_title"`
		MethodId    string        `json:"method_id"`
		InstanceId  string        `json:"instance_id"`
		Total       string        `json:"total"`
		TotalTax    string        `json:"total_tax"`
		Taxes       []interface{} `json:"taxes"`
		MetaData    []struct {
			Id           int         `json:"id"`
			Key          string      `json:"key"`
			Value        interface{} `json:"value"`
			DisplayKey   string      `json:"display_key"`
			DisplayValue interface{} `json:"display_value"`
		} `json:"meta_data"`
	} `json:"shipping_lines"`
	FeeLines         []interface{} `json:"fee_lines"`
	CouponLines      []interface{} `json:"coupon_lines"`
	Refunds          []interface{} `json:"refunds"`
	DateCreatedGmt   string        `json:"date_created_gmt"`
	DateModifiedGmt  string        `json:"date_modified_gmt"`
	DateCompletedGmt interface{}   `json:"date_completed_gmt"`
	DatePaidGmt      interface{}   `json:"date_paid_gmt"`
	CurrencySymbol   string        `json:"currency_symbol"`
	Links            struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
	} `json:"_links"`
}

// UpdateOrderBody is to structure the data
type UpdateOrderBody struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

// Orders is to get all orders since id
func Orders(page int, afterDate string, r Request) ([]OrdersReturn, error) {

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/orders?page=%d&per_page=100&after=%s", page, afterDate), "GET", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return nil, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode []OrdersReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return nil, err
	}

	// Return data
	return decode, err

}

// UpdateOrder is to update the status of an order
func UpdateOrder(body UpdateOrderBody, r Request) (OrdersReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return OrdersReturn{}, err
	}

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/orders/%d", body.Id), "PUT", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return OrdersReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode OrdersReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return OrdersReturn{}, err
	}

	// Return data
	return decode, err

}
