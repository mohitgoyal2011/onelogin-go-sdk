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

// RuleId struct for RuleId
type RuleId struct {
	Id *int32 `json:"id,omitempty"`
}

// NewRuleId instantiates a new RuleId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleId() *RuleId {
	this := RuleId{}
	return &this
}

// NewRuleIdWithDefaults instantiates a new RuleId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleIdWithDefaults() *RuleId {
	this := RuleId{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RuleId) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleId) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RuleId) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RuleId) SetId(v int32) {
	o.Id = &v
}

func (o RuleId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableRuleId struct {
	value *RuleId
	isSet bool
}

func (v NullableRuleId) Get() *RuleId {
	return v.value
}

func (v *NullableRuleId) Set(val *RuleId) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleId) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleId(val *RuleId) *NullableRuleId {
	return &NullableRuleId{value: val, isSet: true}
}

func (v NullableRuleId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


