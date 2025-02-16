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

// Status2 struct for Status2
type Status2 struct {
	Status *Status2Status `json:"status,omitempty"`
}

// NewStatus2 instantiates a new Status2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatus2() *Status2 {
	this := Status2{}
	return &this
}

// NewStatus2WithDefaults instantiates a new Status2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatus2WithDefaults() *Status2 {
	this := Status2{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Status2) GetStatus() Status2Status {
	if o == nil || o.Status == nil {
		var ret Status2Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Status2) GetStatusOk() (*Status2Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Status2) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status2Status and assigns it to the Status field.
func (o *Status2) SetStatus(v Status2Status) {
	o.Status = &v
}

func (o Status2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableStatus2 struct {
	value *Status2
	isSet bool
}

func (v NullableStatus2) Get() *Status2 {
	return v.value
}

func (v *NullableStatus2) Set(val *Status2) {
	v.value = val
	v.isSet = true
}

func (v NullableStatus2) IsSet() bool {
	return v.isSet
}

func (v *NullableStatus2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatus2(val *Status2) *NullableStatus2 {
	return &NullableStatus2{value: val, isSet: true}
}

func (v NullableStatus2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatus2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


