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

// ProductsReturn is to decode the product return
type ProductsReturn struct {
	Id                int           `json:"id"`
	Name              string        `json:"name"`
	Slug              string        `json:"slug"`
	Permalink         string        `json:"permalink"`
	DateCreated       string        `json:"date_created"`
	DateCreatedGmt    string        `json:"date_created_gmt"`
	DateModified      string        `json:"date_modified"`
	DateModifiedGmt   string        `json:"date_modified_gmt"`
	Type              string        `json:"type"`
	Status            string        `json:"status"`
	Featured          bool          `json:"featured"`
	CatalogVisibility string        `json:"catalog_visibility"`
	Description       string        `json:"description"`
	ShortDescription  string        `json:"short_description"`
	Sku               string        `json:"sku"`
	Price             string        `json:"price"`
	RegularPrice      string        `json:"regular_price"`
	SalePrice         string        `json:"sale_price"`
	DateOnSaleFrom    interface{}   `json:"date_on_sale_from"`
	DateOnSaleFromGmt interface{}   `json:"date_on_sale_from_gmt"`
	DateOnSaleTo      interface{}   `json:"date_on_sale_to"`
	DateOnSaleToGmt   interface{}   `json:"date_on_sale_to_gmt"`
	OnSale            bool          `json:"on_sale"`
	Purchasable       bool          `json:"purchasable"`
	TotalSales        int           `json:"total_sales"`
	Virtual           bool          `json:"virtual"`
	Downloadable      bool          `json:"downloadable"`
	Downloads         []interface{} `json:"downloads"`
	DownloadLimit     int           `json:"download_limit"`
	DownloadExpiry    int           `json:"download_expiry"`
	ExternalUrl       string        `json:"external_url"`
	ButtonText        string        `json:"button_text"`
	TaxStatus         string        `json:"tax_status"`
	TaxClass          string        `json:"tax_class"`
	ManageStock       bool          `json:"manage_stock"`
	StockQuantity     interface{}   `json:"stock_quantity"`
	Backorders        string        `json:"backorders"`
	BackordersAllowed bool          `json:"backorders_allowed"`
	Backordered       bool          `json:"backordered"`
	LowStockAmount    interface{}   `json:"low_stock_amount"`
	SoldIndividually  bool          `json:"sold_individually"`
	Weight            string        `json:"weight"`
	Dimensions        struct {
		Length string `json:"length"`
		Width  string `json:"width"`
		Height string `json:"height"`
	} `json:"dimensions"`
	ShippingRequired bool          `json:"shipping_required"`
	ShippingTaxable  bool          `json:"shipping_taxable"`
	ShippingClass    string        `json:"shipping_class"`
	ShippingClassId  int           `json:"shipping_class_id"`
	ReviewsAllowed   bool          `json:"reviews_allowed"`
	AverageRating    string        `json:"average_rating"`
	RatingCount      int           `json:"rating_count"`
	UpsellIds        []interface{} `json:"upsell_ids"`
	CrossSellIds     []interface{} `json:"cross_sell_ids"`
	ParentId         int           `json:"parent_id"`
	PurchaseNote     string        `json:"purchase_note"`
	Categories       []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"categories"`
	Tags              []interface{} `json:"tags"`
	Images            []interface{} `json:"images"`
	Attributes        []interface{} `json:"attributes"`
	DefaultAttributes []interface{} `json:"default_attributes"`
	Variations        []interface{} `json:"variations"`
	GroupedProducts   []interface{} `json:"grouped_products"`
	MenuOrder         int           `json:"menu_order"`
	PriceHtml         string        `json:"price_html"`
	RelatedIds        []interface{} `json:"related_ids"`
	MetaData          []interface{} `json:"meta_data"`
	StockStatus       string        `json:"stock_status"`
	Links             struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
	} `json:"_links"`
}

// Products are to get a list of all products per page
func Products(page int, r *Request) ([]ProductsReturn, error) {

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/products?page=%d&per_page=100", page), "GET", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return nil, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode []ProductsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return nil, err
	}

	// Return data
	return decode, err

}

// DeleteProduct is to remove a product from woocommerce
func DeleteProduct(id int, force bool, r *Request) (ProductsReturn, error) {

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/products/%d?force=%t", id, force), "DELETE", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductsReturn{}, err
	}

	// Return data
	return decode, err

}