/*
OneLogin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GenerateTokenRequest struct for GenerateTokenRequest
type GenerateTokenRequest struct {
	GrantType *string `json:"grant_type,omitempty"`
}

// NewGenerateTokenRequest instantiates a new GenerateTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateTokenRequest() *GenerateTokenRequest {
	this := GenerateTokenRequest{}
	return &this
}

// NewGenerateTokenRequestWithDefaults instantiates a new GenerateTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateTokenRequestWithDefaults() *GenerateTokenRequest {
	this := GenerateTokenRequest{}
	return &this
}

// GetGrantType returns the GrantType field value if set, zero value otherwise.
func (o *GenerateTokenRequest) GetGrantType() string {
	if o == nil || o.GrantType == nil {
		var ret string
		return ret
	}
	return *o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateTokenRequest) GetGrantTypeOk() (*string, bool) {
	if o == nil || o.GrantType == nil {
		return nil, false
	}
	return o.GrantType, true
}

// HasGrantType returns a boolean if a field has been set.
func (o *GenerateTokenRequest) HasGrantType() bool {
	if o != nil && o.GrantType != nil {
		return true
	}

	return false
}

// SetGrantType gets a reference to the given string and assigns it to the GrantType field.
func (o *GenerateTokenRequest) SetGrantType(v string) {
	o.GrantType = &v
}

func (o GenerateTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GrantType != nil {
		toSerialize["grant_type"] = o.GrantType
	}
	return json.Marshal(toSerialize)
}

type NullableGenerateTokenRequest struct {
	value *GenerateTokenRequest
	isSet bool
}

func (v NullableGenerateTokenRequest) Get() *GenerateTokenRequest {
	return v.value
}

func (v *NullableGenerateTokenRequest) Set(val *GenerateTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateTokenRequest(val *GenerateTokenRequest) *NullableGenerateTokenRequest {
	return &NullableGenerateTokenRequest{value: val, isSet: true}
}

func (v NullableGenerateTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


