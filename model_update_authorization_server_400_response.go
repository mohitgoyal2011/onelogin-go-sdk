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

// UpdateAuthorizationServer400Response struct for UpdateAuthorizationServer400Response
type UpdateAuthorizationServer400Response struct {
	Field *string `json:"field,omitempty"`
	Message *string `json:"message,omitempty"`
	StatusCode *int32 `json:"statusCode,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewUpdateAuthorizationServer400Response instantiates a new UpdateAuthorizationServer400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAuthorizationServer400Response() *UpdateAuthorizationServer400Response {
	this := UpdateAuthorizationServer400Response{}
	return &this
}

// NewUpdateAuthorizationServer400ResponseWithDefaults instantiates a new UpdateAuthorizationServer400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAuthorizationServer400ResponseWithDefaults() *UpdateAuthorizationServer400Response {
	this := UpdateAuthorizationServer400Response{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *UpdateAuthorizationServer400Response) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthorizationServer400Response) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *UpdateAuthorizationServer400Response) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *UpdateAuthorizationServer400Response) SetField(v string) {
	o.Field = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UpdateAuthorizationServer400Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthorizationServer400Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UpdateAuthorizationServer400Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UpdateAuthorizationServer400Response) SetMessage(v string) {
	o.Message = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *UpdateAuthorizationServer400Response) GetStatusCode() int32 {
	if o == nil || o.StatusCode == nil {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthorizationServer400Response) GetStatusCodeOk() (*int32, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *UpdateAuthorizationServer400Response) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *UpdateAuthorizationServer400Response) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateAuthorizationServer400Response) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthorizationServer400Response) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateAuthorizationServer400Response) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateAuthorizationServer400Response) SetName(v string) {
	o.Name = &v
}

func (o UpdateAuthorizationServer400Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.StatusCode != nil {
		toSerialize["statusCode"] = o.StatusCode
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateAuthorizationServer400Response struct {
	value *UpdateAuthorizationServer400Response
	isSet bool
}

func (v NullableUpdateAuthorizationServer400Response) Get() *UpdateAuthorizationServer400Response {
	return v.value
}

func (v *NullableUpdateAuthorizationServer400Response) Set(val *UpdateAuthorizationServer400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAuthorizationServer400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAuthorizationServer400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAuthorizationServer400Response(val *UpdateAuthorizationServer400Response) *NullableUpdateAuthorizationServer400Response {
	return &NullableUpdateAuthorizationServer400Response{value: val, isSet: true}
}

func (v NullableUpdateAuthorizationServer400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAuthorizationServer400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


