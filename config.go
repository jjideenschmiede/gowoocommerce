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
	"bytes"
	"net/http"
)

// Config is to define config data
type Config struct {
	Path, Method string
	Body         []byte
}

// Request is to define the request data
type Request struct {
	BaseUrl, ConsumerKey, ConsumerSecret string
}

// Send is to send a new request
func (c *Config) Send(r *Request) (*http.Response, error) {

	// Set url
	url := r.BaseUrl + c.Path

	// Define client
	client := &http.Client{}

	// Request
	request, err := http.NewRequest(c.Method, url, bytes.NewBuffer(c.Body))
	if err != nil {
		return nil, err
	}

	// Define header
	request.Header.Set("Content-Type", "application/json")

	// Define url parameter
	parameter := request.URL.Query()

	parameter.Add("consumer_key", r.ConsumerKey)
	parameter.Add("consumer_secret", r.ConsumerSecret)

	// Set new url
	request.URL.RawQuery = parameter.Encode()

	// Send request & get response
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Return data
	return response, nil

}
