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

// ActivateFactorRequest struct for ActivateFactorRequest
type ActivateFactorRequest struct {
	// Required. Specifies the factor to be verified.
	DeviceId *int32 `json:"device_id,omitempty"`
	// Optional. Sets the window of time in seconds that the factor must be verified within. 
	ExpiresIn *int32 `json:"expires_in,omitempty"`
	// Optional. Only applies to Email MagicLink factor.
	RedirectTo *string `json:"redirect_to,omitempty"`
	// Optional. Only applies to SMS factor. A message template that will be sent via SMS.
	CustomMessage *string `json:"custom_message,omitempty"`
}

// NewActivateFactorRequest instantiates a new ActivateFactorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivateFactorRequest() *ActivateFactorRequest {
	this := ActivateFactorRequest{}
	return &this
}

// NewActivateFactorRequestWithDefaults instantiates a new ActivateFactorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivateFactorRequestWithDefaults() *ActivateFactorRequest {
	this := ActivateFactorRequest{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *ActivateFactorRequest) GetDeviceId() int32 {
	if o == nil || o.DeviceId == nil {
		var ret int32
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateFactorRequest) GetDeviceIdOk() (*int32, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *ActivateFactorRequest) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given int32 and assigns it to the DeviceId field.
func (o *ActivateFactorRequest) SetDeviceId(v int32) {
	o.DeviceId = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *ActivateFactorRequest) GetExpiresIn() int32 {
	if o == nil || o.ExpiresIn == nil {
		var ret int32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateFactorRequest) GetExpiresInOk() (*int32, bool) {
	if o == nil || o.ExpiresIn == nil {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *ActivateFactorRequest) HasExpiresIn() bool {
	if o != nil && o.ExpiresIn != nil {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int32 and assigns it to the ExpiresIn field.
func (o *ActivateFactorRequest) SetExpiresIn(v int32) {
	o.ExpiresIn = &v
}

// GetRedirectTo returns the RedirectTo field value if set, zero value otherwise.
func (o *ActivateFactorRequest) GetRedirectTo() string {
	if o == nil || o.RedirectTo == nil {
		var ret string
		return ret
	}
	return *o.RedirectTo
}

// GetRedirectToOk returns a tuple with the RedirectTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateFactorRequest) GetRedirectToOk() (*string, bool) {
	if o == nil || o.RedirectTo == nil {
		return nil, false
	}
	return o.RedirectTo, true
}

// HasRedirectTo returns a boolean if a field has been set.
func (o *ActivateFactorRequest) HasRedirectTo() bool {
	if o != nil && o.RedirectTo != nil {
		return true
	}

	return false
}

// SetRedirectTo gets a reference to the given string and assigns it to the RedirectTo field.
func (o *ActivateFactorRequest) SetRedirectTo(v string) {
	o.RedirectTo = &v
}

// GetCustomMessage returns the CustomMessage field value if set, zero value otherwise.
func (o *ActivateFactorRequest) GetCustomMessage() string {
	if o == nil || o.CustomMessage == nil {
		var ret string
		return ret
	}
	return *o.CustomMessage
}

// GetCustomMessageOk returns a tuple with the CustomMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateFactorRequest) GetCustomMessageOk() (*string, bool) {
	if o == nil || o.CustomMessage == nil {
		return nil, false
	}
	return o.CustomMessage, true
}

// HasCustomMessage returns a boolean if a field has been set.
func (o *ActivateFactorRequest) HasCustomMessage() bool {
	if o != nil && o.CustomMessage != nil {
		return true
	}

	return false
}

// SetCustomMessage gets a reference to the given string and assigns it to the CustomMessage field.
func (o *ActivateFactorRequest) SetCustomMessage(v string) {
	o.CustomMessage = &v
}

func (o ActivateFactorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceId != nil {
		toSerialize["device_id"] = o.DeviceId
	}
	if o.ExpiresIn != nil {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if o.RedirectTo != nil {
		toSerialize["redirect_to"] = o.RedirectTo
	}
	if o.CustomMessage != nil {
		toSerialize["custom_message"] = o.CustomMessage
	}
	return json.Marshal(toSerialize)
}

type NullableActivateFactorRequest struct {
	value *ActivateFactorRequest
	isSet bool
}

func (v NullableActivateFactorRequest) Get() *ActivateFactorRequest {
	return v.value
}

func (v *NullableActivateFactorRequest) Set(val *ActivateFactorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableActivateFactorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableActivateFactorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivateFactorRequest(val *ActivateFactorRequest) *NullableActivateFactorRequest {
	return &NullableActivateFactorRequest{value: val, isSet: true}
}

func (v NullableActivateFactorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivateFactorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


