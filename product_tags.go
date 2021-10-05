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

// ProductTagsBody is to structure the body data
type ProductTagsBody struct {
	Name        string `json:"name"`
	Slug        string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"`
}

// ProductTagsReturn is to decode the json data
type ProductTagsReturn struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Count       int    `json:"count"`
	Links       struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
	} `json:"_links"`
}

// ProductTags is get a list of all product tags
func ProductTags(r *Request) ([]ProductTagsReturn, error) {

	// Set config for new request
	c := Config{"/wp-json/wc/v3/products/tags", "GET", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return nil, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode []ProductTagsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return nil, err
	}

	// Return data
	return decode, err

}

// ProductTag is get a product tag
func ProductTag(id int, r *Request) (ProductTagsReturn, error) {

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/products/tags/%d", id), "GET", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductTagsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductTagsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductTagsReturn{}, err
	}

	// Return data
	return decode, err

}

// CreateProductTag is to create a new product tag
func CreateProductTag(body ProductTagsBody, r *Request) (ProductTagsReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return ProductTagsReturn{}, err
	}

	// Set config for new request
	c := Config{"/wp-json/wc/v3/products/tags", "POST", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductTagsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductTagsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductTagsReturn{}, err
	}

	// Return data
	return decode, err

}

// UpdateProductTag is to update an existing product tag
func UpdateProductTag(id int, body ProductTagsBody, r *Request) (ProductTagsReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return ProductTagsReturn{}, err
	}

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/products/tags/%d", id), "PUT", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductTagsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductTagsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductTagsReturn{}, err
	}

	// Return data
	return decode, err

}

// DeleteProductTag is to delete an existing product tag
func DeleteProductTag(id int, force bool, r *Request) (ProductTagsReturn, error) {

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/products/tags/%d?force=%t", id, force), "DELETE", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductTagsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductTagsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductTagsReturn{}, err
	}

	// Return data
	return decode, err

}
