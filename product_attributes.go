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

// ProductAttributesBody is to structure the data
type ProductAttributesBody struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Type        string `json:"type"`
	OrderBy     string `json:"order_by"`
	HasArchives bool   `json:"has_archives"`
}

// ProductAttributesReturn is to decode the data
type ProductAttributesReturn struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Type        string `json:"type"`
	OrderBy     string `json:"order_by"`
	HasArchives bool   `json:"has_archives"`
	Links       struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
	} `json:"_links"`
}

// ProductAttributes is to get a list of all product attributes
func ProductAttributes(r *Request) ([]ProductAttributesReturn, error) {

	// Set config for new request
	c := Config{"/wp-json/wc/v3/products/attributes", "GET", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return nil, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode []ProductAttributesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return nil, err
	}

	// Return data
	return decode, err

}

// CreateProductAttributes is to create a new attribute
func CreateProductAttributes(body ProductAttributesBody, r *Request) (ProductAttributesReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return ProductAttributesReturn{}, err
	}

	// Set config for new request
	c := Config{"/wp-json/wc/v3/products/attributes/", "POST", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductAttributesReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductAttributesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductAttributesReturn{}, err
	}

	// Return data
	return decode, err

}

// UpdateProductAttributes is to update a attribute
func UpdateProductAttributes(id int, body ProductAttributesBody, r *Request) (ProductAttributesReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return ProductAttributesReturn{}, err
	}

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/products/attributes/%d", id), "PUT", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductAttributesReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductAttributesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductAttributesReturn{}, err
	}

	// Return data
	return decode, err

}

// DeleteProductAttributes is to delete a product attribute
func DeleteProductAttributes(id int, r *Request) (ProductAttributesReturn, error) {

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/products/attributes/%d", id), "DELETE", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductAttributesReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductAttributesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductAttributesReturn{}, err
	}

	// Return data
	return decode, err

}
