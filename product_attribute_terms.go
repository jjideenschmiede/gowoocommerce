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

// ProductAttributeTermBody is to structure the request data
type ProductAttributeTermBody struct {
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"`
	MenuOrder   int    `json:"menu_order,omitempty"`
}

// ProductAttributeTermsReturn is to decode the data
type ProductAttributeTermsReturn struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	MenuOrder   int    `json:"menu_order"`
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

// ProductAttributeTerms is to get all terms
func ProductAttributeTerms(attributeId int, r *Request) ([]ProductAttributeTermsReturn, error) {

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/products/attributes/%d/terms", attributeId), "GET", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return nil, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode []ProductAttributeTermsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return nil, err
	}

	// Return data
	return decode, err

}

// CreateProductAttributeTerms is to create a new term
func CreateProductAttributeTerms(attributeId int, body ProductAttributeTermBody, r *Request) (ProductAttributeTermsReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return ProductAttributeTermsReturn{}, err
	}

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/products/attributes/%d/terms", attributeId), "POST", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductAttributeTermsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductAttributeTermsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductAttributeTermsReturn{}, err
	}

	// Return data
	return decode, err

}

// UpdateProductAttributeTerms is to update a term
func UpdateProductAttributeTerms(attributeId, termId int, body ProductAttributeTermBody, r *Request) (ProductAttributeTermsReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return ProductAttributeTermsReturn{}, err
	}

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/products/attributes/%d/terms/%d", attributeId, termId), "PUT", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductAttributeTermsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductAttributeTermsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductAttributeTermsReturn{}, err
	}

	// Return data
	return decode, err

}

// DeleteProductAttributeTerms is to delete a term
func DeleteProductAttributeTerms(attributeId, termId int, force bool, r *Request) (ProductAttributeTermsReturn, error) {

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/products/attributes/%d/terms/%d?force=%t", attributeId, termId, force), "DELETE", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductAttributeTermsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductAttributeTermsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductAttributeTermsReturn{}, err
	}

	// Return data
	return decode, err

}
