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
	Attributes        []interface{}            `json:"attributes,omitempty"`
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

// ProductsReturn is to decode the product return
type ProductsReturn struct {
	Id                int           `json:"id,omitempty"`
	Name              string        `json:"name,omitempty"`
	Slug              string        `json:"slug,omitempty"`
	Permalink         string        `json:"permalink,omitempty"`
	DateCreated       string        `json:"date_created,omitempty"`
	DateCreatedGmt    string        `json:"date_created_gmt,omitempty"`
	DateModified      string        `json:"date_modified,omitempty"`
	DateModifiedGmt   string        `json:"date_modified_gmt,omitempty"`
	Type              string        `json:"type,omitempty"`
	Status            string        `json:"status,omitempty"`
	Featured          bool          `json:"featured,omitempty"`
	CatalogVisibility string        `json:"catalog_visibility,omitempty"`
	Description       string        `json:"description,omitempty"`
	ShortDescription  string        `json:"short_description,omitempty"`
	Sku               string        `json:"sku,omitempty"`
	Price             string        `json:"price,omitempty"`
	RegularPrice      string        `json:"regular_price,omitempty"`
	SalePrice         string        `json:"sale_price,omitempty"`
	DateOnSaleFrom    interface{}   `json:"date_on_sale_from,omitempty"`
	DateOnSaleFromGmt interface{}   `json:"date_on_sale_from_gmt,omitempty"`
	DateOnSaleTo      interface{}   `json:"date_on_sale_to,omitempty"`
	DateOnSaleToGmt   interface{}   `json:"date_on_sale_to_gmt,omitempty"`
	OnSale            bool          `json:"on_sale,omitempty"`
	Purchasable       bool          `json:"purchasable,omitempty"`
	TotalSales        int           `json:"total_sales,omitempty"`
	Virtual           bool          `json:"virtual,omitempty"`
	Downloadable      bool          `json:"downloadable,omitempty"`
	Downloads         []interface{} `json:"downloads,omitempty"`
	DownloadLimit     int           `json:"download_limit,omitempty"`
	DownloadExpiry    int           `json:"download_expiry,omitempty"`
	ExternalUrl       string        `json:"external_url,omitempty"`
	ButtonText        string        `json:"button_text,omitempty"`
	TaxStatus         string        `json:"tax_status,omitempty"`
	TaxClass          string        `json:"tax_class,omitempty"`
	ManageStock       bool          `json:"manage_stock,omitempty"`
	StockQuantity     interface{}   `json:"stock_quantity,omitempty"`
	Backorders        string        `json:"backorders,omitempty"`
	BackordersAllowed bool          `json:"backorders_allowed,omitempty"`
	Backordered       bool          `json:"backordered,omitempty"`
	LowStockAmount    interface{}   `json:"low_stock_amount,omitempty"`
	SoldIndividually  bool          `json:"sold_individually,omitempty"`
	Weight            string        `json:"weight,omitempty"`
	Dimensions        struct {
		Length string `json:"length,omitempty"`
		Width  string `json:"width,omitempty"`
		Height string `json:"height,omitempty"`
	} `json:"dimensions,omitempty"`
	ShippingRequired bool          `json:"shipping_required,omitempty"`
	ShippingTaxable  bool          `json:"shipping_taxable,omitempty"`
	ShippingClass    string        `json:"shipping_class,omitempty"`
	ShippingClassId  int           `json:"shipping_class_id,omitempty"`
	ReviewsAllowed   bool          `json:"reviews_allowed,omitempty"`
	AverageRating    string        `json:"average_rating,omitempty"`
	RatingCount      int           `json:"rating_count,omitempty"`
	UpsellIds        []interface{} `json:"upsell_ids,omitempty"`
	CrossSellIds     []interface{} `json:"cross_sell_ids,omitempty"`
	ParentId         int           `json:"parent_id,omitempty"`
	PurchaseNote     string        `json:"purchase_note,omitempty"`
	Categories       []struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Slug string `json:"slug,omitempty"`
	} `json:"categories,omitempty"`
	Tags              []interface{} `json:"tags,omitempty"`
	Images            []interface{} `json:"images,omitempty"`
	Attributes        []interface{} `json:"attributes,omitempty"`
	DefaultAttributes []interface{} `json:"default_attributes,omitempty"`
	Variations        []interface{} `json:"variations,omitempty"`
	GroupedProducts   []interface{} `json:"grouped_products,omitempty"`
	MenuOrder         int           `json:"menu_order,omitempty"`
	PriceHtml         string        `json:"price_html,omitempty"`
	RelatedIds        []interface{} `json:"related_ids,omitempty"`
	MetaData          []interface{} `json:"meta_data,omitempty"`
	StockStatus       string        `json:"stock_status,omitempty"`
	Links             struct {
		Self []struct {
			Href string `json:"href,omitempty"`
		} `json:"self,omitempty"`
		Collection []struct {
			Href string `json:"href,omitempty"`
		} `json:"collection,omitempty"`
	} `json:"_links,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		Status int `json:"status,omitempty"`
		Params struct {
			Backorders string `json:"backorders,omitempty"`
		} `json:"params,omitempty"`
		Details struct {
			Backorders struct {
				Code    string      `json:"code,omitempty"`
				Message string      `json:"message,omitempty"`
				Data    interface{} `json:"data,omitempty"`
			} `json:"backorders,omitempty"`
		} `json:"details,omitempty"`
	} `json:"data,omitempty"`
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
