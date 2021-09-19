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

// ProductsBody is to structure the request data
type ProductsBody struct {
	Name              string                   `json:"name,omitempty"`
	Slug              string                   `json:"slug,omitempty"`
	Type              string                   `json:"type,omitempty"`
	Status            string                   `json:"status,omitempty"`
	Featured          bool                     `json:"featured,omitempty"`
	CatalogVisibility string                   `json:"catalog_visibility,omitempty"`
	Description       string                   `json:"description,omitempty"`
	ShortDescription  string                   `json:"short_description,omitempty"`
	Sku               string                   `json:"sku,omitempty"`
	Price             string                   `json:"price,omitempty"`
	RegularPrice      string                   `json:"regular_price,omitempty"`
	SalePrice         string                   `json:"sale_price,omitempty"`
	DateOnSaleFrom    interface{}              `json:"date_on_sale_from,omitempty"`
	DateOnSaleFromGmt interface{}              `json:"date_on_sale_from_gmt,omitempty"`
	DateOnSaleTo      interface{}              `json:"date_on_sale_to,omitempty"`
	DateOnSaleToGmt   interface{}              `json:"date_on_sale_to_gmt,omitempty"`
	OnSale            bool                     `json:"on_sale,omitempty"`
	Purchasable       bool                     `json:"purchasable,omitempty"`
	TotalSales        int                      `json:"total_sales,omitempty"`
	Virtual           bool                     `json:"virtual,omitempty"`
	Downloadable      bool                     `json:"downloadable,omitempty"`
	Downloads         []interface{}            `json:"downloads,omitempty"`
	DownloadLimit     int                      `json:"download_limit,omitempty"`
	DownloadExpiry    int                      `json:"download_expiry,omitempty"`
	ExternalUrl       string                   `json:"external_url,omitempty"`
	ButtonText        string                   `json:"button_text,omitempty"`
	TaxStatus         string                   `json:"tax_status,omitempty"`
	TaxClass          string                   `json:"tax_class,omitempty"`
	ManageStock       bool                     `json:"manage_stock,omitempty"`
	StockQuantity     interface{}              `json:"stock_quantity,omitempty"`
	Backorders        string                   `json:"backorders,omitempty"`
	BackordersAllowed bool                     `json:"backorders_allowed,omitempty"`
	Backordered       bool                     `json:"backordered,omitempty"`
	LowStockAmount    interface{}              `json:"low_stock_amount,omitempty"`
	SoldIndividually  bool                     `json:"sold_individually,omitempty"`
	Weight            string                   `json:"weight,omitempty"`
	Dimensions        ProductsBodyDimensions   `json:"dimensions,omitempty"`
	ShippingRequired  bool                     `json:"shipping_required,omitempty"`
	ShippingTaxable   bool                     `json:"shipping_taxable,omitempty"`
	ShippingClass     string                   `json:"shipping_class,omitempty"`
	ShippingClassId   int                      `json:"shipping_class_id,omitempty"`
	ReviewsAllowed    bool                     `json:"reviews_allowed,omitempty"`
	AverageRating     string                   `json:"average_rating,omitempty"`
	RatingCount       int                      `json:"rating_count,omitempty"`
	UpsellIds         []interface{}            `json:"upsell_ids,omitempty"`
	CrossSellIds      []interface{}            `json:"cross_sell_ids,omitempty"`
	ParentId          int                      `json:"parent_id,omitempty"`
	PurchaseNote      string                   `json:"purchase_note,omitempty"`
	Categories        []ProductsBodyCategories `json:"categories,omitempty"`
	Tags              []interface{}            `json:"tags,omitempty"`
	Images            []ProductsBodyImages     `json:"images,omitempty"`
	Attributes        []ProductsBodyAttributes{}            `json:"attributes,omitempty"`
	DefaultAttributes []interface{}            `json:"default_attributes,omitempty"`
	Variations        []interface{}            `json:"variations,omitempty"`
	GroupedProducts   []interface{}            `json:"grouped_products,omitempty"`
	MenuOrder         int                      `json:"menu_order,omitempty"`
	RelatedIds        []interface{}            `json:"related_ids,omitempty"`
	MetaData          []interface{}            `json:"meta_data,omitempty"`
	StockStatus       string                   `json:"stock_status,omitempty"`
}

type ProductsBodyDimensions struct {
	Length string `json:"length,omitempty"`
	Width  string `json:"width,omitempty"`
	Height string `json:"height,omitempty"`
}

type ProductsBodyCategories struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

type ProductsBodyImages struct {
	Src  string `json:"src,omitempty"`
	Name string `json:"name,omitempty"`
	Alt  string `json:"alt,omitempty"`
}

type ProductsBodyAttributes struct {
	Name string `json:"name"`
	Position int `json:"position"`
	Visible bool `json:"visible"`
	Variation bool `json:"variation"`
	Options []string `json:"options"`
}

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
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Status int `json:"status"`
		Params struct {
			Backorders string `json:"backorders"`
		} `json:"params"`
		Details struct {
			Backorders struct {
				Code    string      `json:"code"`
				Message string      `json:"message"`
				Data    interface{} `json:"data"`
			} `json:"backorders"`
		} `json:"details"`
	} `json:"data"`
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

// CreateProduct is to create a new product
func CreateProduct(body ProductsBody, r *Request) (ProductsReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return ProductsReturn{}, err
	}

	// Set config for new request
	c := Config{"/wp-json/wc/v3/products", "POST", convert}

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

// UpdateProduct is to update a product
func UpdateProduct(id int, body ProductsBody, r *Request) (ProductsReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return ProductsReturn{}, err
	}

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/products/%d", id), "PUT", convert}

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
