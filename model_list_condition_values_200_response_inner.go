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

// ListConditionValues200ResponseInner struct for ListConditionValues200ResponseInner
type ListConditionValues200ResponseInner struct {
	Name *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewListConditionValues200ResponseInner instantiates a new ListConditionValues200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConditionValues200ResponseInner() *ListConditionValues200ResponseInner {
	this := ListConditionValues200ResponseInner{}
	return &this
}

// NewListConditionValues200ResponseInnerWithDefaults instantiates a new ListConditionValues200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConditionValues200ResponseInnerWithDefaults() *ListConditionValues200ResponseInner {
	this := ListConditionValues200ResponseInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListConditionValues200ResponseInner) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConditionValues200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListConditionValues200ResponseInner) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListConditionValues200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ListConditionValues200ResponseInner) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConditionValues200ResponseInner) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ListConditionValues200ResponseInner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ListConditionValues200ResponseInner) SetValue(v string) {
	o.Value = &v
}

func (o ListConditionValues200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableListConditionValues200ResponseInner struct {
	value *ListConditionValues200ResponseInner
	isSet bool
}

func (v NullableListConditionValues200ResponseInner) Get() *ListConditionValues200ResponseInner {
	return v.value
}

func (v *NullableListConditionValues200ResponseInner) Set(val *ListConditionValues200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListConditionValues200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListConditionValues200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConditionValues200ResponseInner(val *ListConditionValues200ResponseInner) *NullableListConditionValues200ResponseInner {
	return &NullableListConditionValues200ResponseInner{value: val, isSet: true}
}

func (v NullableListConditionValues200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConditionValues200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


