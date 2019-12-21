/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"bytes"
	"encoding/json"
)

// Client struct for Client
type Client struct {
	Client *string `json:"client,omitempty"`
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *Client) GetClient() string {
	if o == nil || o.Client == nil {
		var ret string
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetClientOk() (string, bool) {
	if o == nil || o.Client == nil {
		var ret string
		return ret, false
	}
	return *o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *Client) HasClient() bool {
	if o != nil && o.Client != nil {
		return true
	}

	return false
}

// SetClient gets a reference to the given string and assigns it to the Client field.
func (o *Client) SetClient(v string) {
	o.Client = &v
}

type NullableClient struct {
	Value Client
	ExplicitNull bool
}

func (v NullableClient) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableClient) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}