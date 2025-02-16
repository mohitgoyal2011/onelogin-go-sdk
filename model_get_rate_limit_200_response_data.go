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

// GetRateLimit200ResponseData struct for GetRateLimit200ResponseData
type GetRateLimit200ResponseData struct {
	XRateLimitLimit *int32 `json:"X-RateLimit-Limit,omitempty"`
	XRateLimitRemaining *int32 `json:"X-RateLimit-Remaining,omitempty"`
	XRateLimitReset *int32 `json:"X-RateLimit-Reset,omitempty"`
}

// NewGetRateLimit200ResponseData instantiates a new GetRateLimit200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRateLimit200ResponseData() *GetRateLimit200ResponseData {
	this := GetRateLimit200ResponseData{}
	return &this
}

// NewGetRateLimit200ResponseDataWithDefaults instantiates a new GetRateLimit200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRateLimit200ResponseDataWithDefaults() *GetRateLimit200ResponseData {
	this := GetRateLimit200ResponseData{}
	return &this
}

// GetXRateLimitLimit returns the XRateLimitLimit field value if set, zero value otherwise.
func (o *GetRateLimit200ResponseData) GetXRateLimitLimit() int32 {
	if o == nil || o.XRateLimitLimit == nil {
		var ret int32
		return ret
	}
	return *o.XRateLimitLimit
}

// GetXRateLimitLimitOk returns a tuple with the XRateLimitLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRateLimit200ResponseData) GetXRateLimitLimitOk() (*int32, bool) {
	if o == nil || o.XRateLimitLimit == nil {
		return nil, false
	}
	return o.XRateLimitLimit, true
}

// HasXRateLimitLimit returns a boolean if a field has been set.
func (o *GetRateLimit200ResponseData) HasXRateLimitLimit() bool {
	if o != nil && o.XRateLimitLimit != nil {
		return true
	}

	return false
}

// SetXRateLimitLimit gets a reference to the given int32 and assigns it to the XRateLimitLimit field.
func (o *GetRateLimit200ResponseData) SetXRateLimitLimit(v int32) {
	o.XRateLimitLimit = &v
}

// GetXRateLimitRemaining returns the XRateLimitRemaining field value if set, zero value otherwise.
func (o *GetRateLimit200ResponseData) GetXRateLimitRemaining() int32 {
	if o == nil || o.XRateLimitRemaining == nil {
		var ret int32
		return ret
	}
	return *o.XRateLimitRemaining
}

// GetXRateLimitRemainingOk returns a tuple with the XRateLimitRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRateLimit200ResponseData) GetXRateLimitRemainingOk() (*int32, bool) {
	if o == nil || o.XRateLimitRemaining == nil {
		return nil, false
	}
	return o.XRateLimitRemaining, true
}

// HasXRateLimitRemaining returns a boolean if a field has been set.
func (o *GetRateLimit200ResponseData) HasXRateLimitRemaining() bool {
	if o != nil && o.XRateLimitRemaining != nil {
		return true
	}

	return false
}

// SetXRateLimitRemaining gets a reference to the given int32 and assigns it to the XRateLimitRemaining field.
func (o *GetRateLimit200ResponseData) SetXRateLimitRemaining(v int32) {
	o.XRateLimitRemaining = &v
}

// GetXRateLimitReset returns the XRateLimitReset field value if set, zero value otherwise.
func (o *GetRateLimit200ResponseData) GetXRateLimitReset() int32 {
	if o == nil || o.XRateLimitReset == nil {
		var ret int32
		return ret
	}
	return *o.XRateLimitReset
}

// GetXRateLimitResetOk returns a tuple with the XRateLimitReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRateLimit200ResponseData) GetXRateLimitResetOk() (*int32, bool) {
	if o == nil || o.XRateLimitReset == nil {
		return nil, false
	}
	return o.XRateLimitReset, true
}

// HasXRateLimitReset returns a boolean if a field has been set.
func (o *GetRateLimit200ResponseData) HasXRateLimitReset() bool {
	if o != nil && o.XRateLimitReset != nil {
		return true
	}

	return false
}

// SetXRateLimitReset gets a reference to the given int32 and assigns it to the XRateLimitReset field.
func (o *GetRateLimit200ResponseData) SetXRateLimitReset(v int32) {
	o.XRateLimitReset = &v
}

func (o GetRateLimit200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.XRateLimitLimit != nil {
		toSerialize["X-RateLimit-Limit"] = o.XRateLimitLimit
	}
	if o.XRateLimitRemaining != nil {
		toSerialize["X-RateLimit-Remaining"] = o.XRateLimitRemaining
	}
	if o.XRateLimitReset != nil {
		toSerialize["X-RateLimit-Reset"] = o.XRateLimitReset
	}
	return json.Marshal(toSerialize)
}

type NullableGetRateLimit200ResponseData struct {
	value *GetRateLimit200ResponseData
	isSet bool
}

func (v NullableGetRateLimit200ResponseData) Get() *GetRateLimit200ResponseData {
	return v.value
}

func (v *NullableGetRateLimit200ResponseData) Set(val *GetRateLimit200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRateLimit200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRateLimit200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRateLimit200ResponseData(val *GetRateLimit200ResponseData) *NullableGetRateLimit200ResponseData {
	return &NullableGetRateLimit200ResponseData{value: val, isSet: true}
}

func (v NullableGetRateLimit200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRateLimit200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


