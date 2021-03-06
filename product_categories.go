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

// ProductCategory is to structure the category data
type ProductCategory struct {
	Name        string                `json:"name,omitempty"`
	Parent      int                   `json:"parent,omitempty"`
	Description string                `json:"description,omitempty"`
	Display     string                `json:"display,omitempty"`
	MenuOrder   int                   `json:"menu_order,omitempty"`
	Image       *ProductCategoryImage `json:"image,omitempty"`
}

type ProductCategoryImage struct {
	Src string `json:"src,omitempty"`
}

// ProductCategoriesReturn is to decode the json data
type ProductCategoriesReturn struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Slug        string      `json:"slug"`
	Parent      int         `json:"parent"`
	Description string      `json:"description"`
	Display     string      `json:"display"`
	Image       interface{} `json:"image"`
	MenuOrder   int         `json:"menu_order"`
	Count       int         `json:"count"`
	Links       struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
	} `json:"_links"`
}

// ProductCategories is to get a list of all product categories
func ProductCategories(r Request) ([]ProductCategoriesReturn, error) {

	// Set config for new request
	c := Config{"/wp-json/wc/v3/products/categories", "GET", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return nil, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode []ProductCategoriesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return nil, err
	}

	// Return data
	return decode, err

}

// AddProductCategory is to create a new category
func AddProductCategory(body ProductCategory, r Request) (ProductCategoriesReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return ProductCategoriesReturn{}, err
	}

	// Set config for new request
	c := Config{"/wp-json/wc/v3/products/categories", "POST", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductCategoriesReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductCategoriesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductCategoriesReturn{}, err
	}

	// Return data
	return decode, err

}

// UpdateProductCategory is to update an existing category
func UpdateProductCategory(id int, body ProductCategory, r Request) (ProductCategoriesReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return ProductCategoriesReturn{}, err
	}

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/products/categories/%d", id), "POST", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductCategoriesReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductCategoriesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductCategoriesReturn{}, err
	}

	// Return data
	return decode, err

}

// DeleteProductCategory is to delete a product category
func DeleteProductCategory(id int, force bool, r Request) (ProductCategoriesReturn, error) {

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/products/categories/%d?force=%t", id, force), "DELETE", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductCategoriesReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductCategoriesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductCategoriesReturn{}, err
	}

	// Return data
	return decode, err

}
