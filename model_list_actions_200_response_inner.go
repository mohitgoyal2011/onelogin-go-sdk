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

// ListActions200ResponseInner struct for ListActions200ResponseInner
type ListActions200ResponseInner struct {
	// The name of the Action.
	Name *string `json:"name,omitempty"`
	// The unique identifier of the action. This should be used when defining actions for a User Mapping.
	Value *string `json:"value,omitempty"`
}

// NewListActions200ResponseInner instantiates a new ListActions200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListActions200ResponseInner() *ListActions200ResponseInner {
	this := ListActions200ResponseInner{}
	return &this
}

// NewListActions200ResponseInnerWithDefaults instantiates a new ListActions200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListActions200ResponseInnerWithDefaults() *ListActions200ResponseInner {
	this := ListActions200ResponseInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListActions200ResponseInner) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListActions200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListActions200ResponseInner) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListActions200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ListActions200ResponseInner) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListActions200ResponseInner) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ListActions200ResponseInner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ListActions200ResponseInner) SetValue(v string) {
	o.Value = &v
}

func (o ListActions200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableListActions200ResponseInner struct {
	value *ListActions200ResponseInner
	isSet bool
}

func (v NullableListActions200ResponseInner) Get() *ListActions200ResponseInner {
	return v.value
}

func (v *NullableListActions200ResponseInner) Set(val *ListActions200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListActions200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListActions200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListActions200ResponseInner(val *ListActions200ResponseInner) *NullableListActions200ResponseInner {
	return &NullableListActions200ResponseInner{value: val, isSet: true}
}

func (v NullableListActions200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListActions200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


