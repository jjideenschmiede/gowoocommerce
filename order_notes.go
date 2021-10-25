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

// OrderNotesBody is to structure the data
type OrderNotesBody struct {
	Note         string `json:"note"`
	CustomerNote bool   `json:"customer_note"`
	AddedByUser  bool   `json:"added_by_user"`
}

// OrderNotesReturn is to decode the json return
type OrderNotesReturn struct {
	Id             int    `json:"id"`
	Author         string `json:"author"`
	DateCreated    string `json:"date_created"`
	DateCreatedGmt string `json:"date_created_gmt"`
	Note           string `json:"note"`
	CustomerNote   bool   `json:"customer_note"`
	Links          struct {
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
		Status int `json:"status,omitempty"`
	} `json:"data,omitempty"`
}

// CreateOrderNote is to create an order note
func CreateOrderNote(id int, body OrderNotesBody, r Request) (OrderNotesReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return OrderNotesReturn{}, err
	}

	// Set config for new request
	c := Config{fmt.Sprintf("/wp-json/wc/v3/orders/%d/notes", id), "POST", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return OrderNotesReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode OrderNotesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return OrderNotesReturn{}, err
	}

	// Return data
	return decode, err

}
