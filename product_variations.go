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

// ProductVariationsBody are to structure the data
type ProductVariationsBody struct {
	Description       string                            `json:"description,omitempty"`
	Permalink         string                            `json:"permalink,omitempty"`
	Sku               string                            `json:"sku,omitempty"`
	Price             string                            `json:"price,omitempty"`
	RegularPrice      string                            `json:"regular_price,omitempty"`
	SalePrice         string                            `json:"sale_price,omitempty"`
	DateOnSaleFrom    interface{}                       `json:"date_on_sale_from,omitempty"`
	DateOnSaleFromGmt interface{}                       `json:"date_on_sale_from_gmt,omitempty"`
	DateOnSaleTo      interface{}                       `json:"date_on_sale_to,omitempty"`
	DateOnSaleToGmt   interface{}                       `json:"date_on_sale_to_gmt,omitempty"`
	OnSale            bool                              `json:"on_sale,omitempty"`
	Status            string                            `json:"status,omitempty"`
	Purchasable       bool                              `json:"purchasable,omitempty"`
	Virtual           bool                              `json:"virtual,omitempty"`
	Downloadable      bool                              `json:"downloadable,omitempty"`
	Downloads         []interface{}                     `json:"downloads,omitempty"`
	DownloadLimit     int                               `json:"download_limit,omitempty"`
	DownloadExpiry    int                               `json:"download_expiry,omitempty"`
	TaxStatus         string                            `json:"tax_status,omitempty"`
	TaxClass          string                            `json:"tax_class,omitempty"`
	ManageStock       bool                              `json:"manage_stock,omitempty"`
	StockQuantity     interface{}                       `json:"stock_quantity,omitempty"`
	StockStatus       string                            `json:"stock_status,omitempty"`
	Backorders        string                            `json:"backorders,omitempty"`
	BackordersAllowed bool                              `json:"backorders_allowed,omitempty"`
	Backordered       bool                              `json:"backordered,omitempty"`
	Weight            string                            `json:"weight,omitempty"`
	Dimensions        ProductVariationsBodyDimensions   `json:"dimensions,omitempty"`
	ShippingClass     string                            `json:"shipping_class,omitempty"`
	ShippingClassId   int                               `json:"shipping_class_id,omitempty"`
	Image             ProductVariationsBodyImages       `json:"image,omitempty"`
	Attributes        []ProductVariationsBodyAttributes `json:"attributes,omitempty"`
	MenuOrder         int                               `json:"menu_order,omitempty"`
	MetaData          []interface{}                     `json:"meta_data,omitempty"`
}

type ProductVariationsBodyDimensions struct {
	Length string `json:"length,omitempty"`
	Width  string `json:"width,omitempty"`
	Height string `json:"height,omitempty"`
}

type ProductVariationsBodyImages struct {
	Src  string `json:"src,omitempty"`
	Name string `json:"name,omitempty"`
	Alt  string `json:"alt,omitempty"`
}

type ProductVariationsBodyAttributes struct {
	Name   string `json:"name,omitempty"`
	Option string `json:"option,omitempty"`
}

// ProductVariationsReturn is to decode the json data
type ProductVariationsReturn struct {
	Id                int           `json:"id"`
	DateCreated       string        `json:"date_created"`
	DateCreatedGmt    string        `json:"date_created_gmt"`
	DateModified      string        `json:"date_modified"`
	DateModifiedGmt   string        `json:"date_modified_gmt"`
	Description       string        `json:"description"`
	Permalink         string        `json:"permalink"`
	Sku               string        `json:"sku"`
	Price             string        `json:"price"`
	RegularPrice      string        `json:"regular_price"`
	SalePrice         string        `json:"sale_price"`
	DateOnSaleFrom    interface{}   `json:"date_on_sale_from"`
	DateOnSaleFromGmt interface{}   `json:"date_on_sale_from_gmt"`
	DateOnSaleTo      interface{}   `json:"date_on_sale_to"`
	DateOnSaleToGmt   interface{}   `json:"date_on_sale_to_gmt"`
	OnSale            bool          `json:"on_sale"`
	Status            string        `json:"status"`
	Purchasable       bool          `json:"purchasable"`
	Virtual           bool          `json:"virtual"`
	Downloadable      bool          `json:"downloadable"`
	Downloads         []interface{} `json:"downloads"`
	DownloadLimit     int           `json:"download_limit"`
	DownloadExpiry    int           `json:"download_expiry"`
	TaxStatus         string        `json:"tax_status"`
	TaxClass          string        `json:"tax_class"`
	ManageStock       bool          `json:"manage_stock"`
	StockQuantity     int           `json:"stock_quantity"`
	StockStatus       string        `json:"stock_status"`
	Backorders        string        `json:"backorders"`
	BackordersAllowed bool          `json:"backorders_allowed"`
	Backordered       bool          `json:"backordered"`
	LowStockAmount    interface{}   `json:"low_stock_amount"`
	Weight            string        `json:"weight"`
	Dimensions        struct {
		Length string `json:"length"`
		Width  string `json:"width"`
		Height string `json:"height"`
	} `json:"dimensions"`
	ShippingClass   string `json:"shipping_class"`
	ShippingClassId int    `json:"shipping_class_id"`
	Image           struct {
		Id              int    `json:"id"`
		DateCreated     string `json:"date_created"`
		DateCreatedGmt  string `json:"date_created_gmt"`
		DateModified    string `json:"date_modified"`
		DateModifiedGmt string `json:"date_modified_gmt"`
		Src             string `json:"src"`
		Name            string `json:"name"`
		Alt             string `json:"alt"`
	} `json:"image"`
	Attributes []struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Option string `json:"option"`
	} `json:"attributes"`
	MenuOrder int           `json:"menu_order"`
	MetaData  []interface{} `json:"meta_data"`
	Links     struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
		Up []struct {
			Href string `json:"href"`
		} `json:"up"`
	} `json:"_links"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		Status int `json:"status"`
		Params struct {
			StockStatus string `json:"stock_status"`
		} `json:"params"`
		Details struct {
			StockStatus struct {
				Code    string      `json:"code"`
				Message string      `json:"message"`
				Data    interface{} `json:"data"`
			} `json:"stock_status"`
		} `json:"details"`
	} `json:"data,omitempty"`
}

// ProductVariations is to get a list of all product variations
func ProductVariations(productId, page int, r Request) ([]ProductVariationsReturn, error) {

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/products/%d/variations?page=%d&per_page=100", productId, page), "GET", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return nil, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode []ProductVariationsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return nil, err
	}

	// Return data
	return decode, err

}

// CreateProductVariations is to create a new variant to a product
func CreateProductVariations(productId int, body ProductVariationsBody, r Request) (ProductVariationsReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return ProductVariationsReturn{}, err
	}

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/products/%d/variations", productId), "POST", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductVariationsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductVariationsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductVariationsReturn{}, err
	}

	// Return data
	return decode, err

}

// UpdateProductVariations is to create a new variant to a product
func UpdateProductVariations(productId, variantId int, body ProductVariationsBody, r Request) (ProductVariationsReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return ProductVariationsReturn{}, err
	}

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/products/%d/variations/%d", productId, variantId), "PUT", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductVariationsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductVariationsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductVariationsReturn{}, err
	}

	// Return data
	return decode, err

}

// UpdateProductVariationStock is to update the stock
func UpdateProductVariationStock(productId, variantId, stock int, r Request) (ProductVariationsReturn, error) {

	// Define request body & set stock to body struct
	type RequestBody struct {
		StockQuantity int `json:"stock_quantity"`
	}

	body := RequestBody{
		StockQuantity: stock,
	}

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return ProductVariationsReturn{}, err
	}

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/products/%d/variations/%d", productId, variantId), "PUT", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductVariationsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductVariationsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductVariationsReturn{}, err
	}

	// Return data
	return decode, err

}

// DeleteProductVariations is to remove a variant from a product
func DeleteProductVariations(productId, variantId int, r Request) (ProductVariationsReturn, error) {

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/products/%d/variations/%d", productId, variantId), "DELETE", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductVariationsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductVariationsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductVariationsReturn{}, err
	}

	// Return data
	return decode, err

}
