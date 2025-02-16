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

// GetClientApps200ResponseInnerScopesInner struct for GetClientApps200ResponseInnerScopesInner
type GetClientApps200ResponseInnerScopesInner struct {
	Description *string `json:"description,omitempty"`
	Value *string `json:"value,omitempty"`
	Id *int32 `json:"id,omitempty"`
}

// NewGetClientApps200ResponseInnerScopesInner instantiates a new GetClientApps200ResponseInnerScopesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetClientApps200ResponseInnerScopesInner() *GetClientApps200ResponseInnerScopesInner {
	this := GetClientApps200ResponseInnerScopesInner{}
	return &this
}

// NewGetClientApps200ResponseInnerScopesInnerWithDefaults instantiates a new GetClientApps200ResponseInnerScopesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetClientApps200ResponseInnerScopesInnerWithDefaults() *GetClientApps200ResponseInnerScopesInner {
	this := GetClientApps200ResponseInnerScopesInner{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetClientApps200ResponseInnerScopesInner) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClientApps200ResponseInnerScopesInner) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetClientApps200ResponseInnerScopesInner) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetClientApps200ResponseInnerScopesInner) SetDescription(v string) {
	o.Description = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetClientApps200ResponseInnerScopesInner) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClientApps200ResponseInnerScopesInner) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetClientApps200ResponseInnerScopesInner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *GetClientApps200ResponseInnerScopesInner) SetValue(v string) {
	o.Value = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetClientApps200ResponseInnerScopesInner) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClientApps200ResponseInnerScopesInner) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetClientApps200ResponseInnerScopesInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetClientApps200ResponseInnerScopesInner) SetId(v int32) {
	o.Id = &v
}

func (o GetClientApps200ResponseInnerScopesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGetClientApps200ResponseInnerScopesInner struct {
	value *GetClientApps200ResponseInnerScopesInner
	isSet bool
}

func (v NullableGetClientApps200ResponseInnerScopesInner) Get() *GetClientApps200ResponseInnerScopesInner {
	return v.value
}

func (v *NullableGetClientApps200ResponseInnerScopesInner) Set(val *GetClientApps200ResponseInnerScopesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetClientApps200ResponseInnerScopesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetClientApps200ResponseInnerScopesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetClientApps200ResponseInnerScopesInner(val *GetClientApps200ResponseInnerScopesInner) *NullableGetClientApps200ResponseInnerScopesInner {
	return &NullableGetClientApps200ResponseInnerScopesInner{value: val, isSet: true}
}

func (v NullableGetClientApps200ResponseInnerScopesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetClientApps200ResponseInnerScopesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


