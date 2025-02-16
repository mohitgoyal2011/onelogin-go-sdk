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

// UpdateEnvironmentVariableRequest struct for UpdateEnvironmentVariableRequest
type UpdateEnvironmentVariableRequest struct {
	// The secret value that will be encrypted at rest and injected in applicable hook functions at run time.
	Value string `json:"value"`
}

// NewUpdateEnvironmentVariableRequest instantiates a new UpdateEnvironmentVariableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEnvironmentVariableRequest(value string) *UpdateEnvironmentVariableRequest {
	this := UpdateEnvironmentVariableRequest{}
	this.Value = value
	return &this
}

// NewUpdateEnvironmentVariableRequestWithDefaults instantiates a new UpdateEnvironmentVariableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEnvironmentVariableRequestWithDefaults() *UpdateEnvironmentVariableRequest {
	this := UpdateEnvironmentVariableRequest{}
	return &this
}

// GetValue returns the Value field value
func (o *UpdateEnvironmentVariableRequest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UpdateEnvironmentVariableRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *UpdateEnvironmentVariableRequest) SetValue(v string) {
	o.Value = v
}

func (o UpdateEnvironmentVariableRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateEnvironmentVariableRequest struct {
	value *UpdateEnvironmentVariableRequest
	isSet bool
}

func (v NullableUpdateEnvironmentVariableRequest) Get() *UpdateEnvironmentVariableRequest {
	return v.value
}

func (v *NullableUpdateEnvironmentVariableRequest) Set(val *UpdateEnvironmentVariableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEnvironmentVariableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEnvironmentVariableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEnvironmentVariableRequest(val *UpdateEnvironmentVariableRequest) *NullableUpdateEnvironmentVariableRequest {
	return &NullableUpdateEnvironmentVariableRequest{value: val, isSet: true}
}

func (v NullableUpdateEnvironmentVariableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEnvironmentVariableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


