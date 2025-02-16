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

// Registration struct for Registration
type Registration struct {
	// Registration identifier.
	Id *string `json:"id,omitempty"`
	// pending registration has not been completed successfully. accepted registration has successfully completed.
	Status *string `json:"status,omitempty"`
	// Device id to be used with Verify the Factor.
	DeviceId *string `json:"device_id,omitempty"`
}

// NewRegistration instantiates a new Registration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistration() *Registration {
	this := Registration{}
	return &this
}

// NewRegistrationWithDefaults instantiates a new Registration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationWithDefaults() *Registration {
	this := Registration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Registration) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Registration) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Registration) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Registration) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Registration) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Registration) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Registration) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Registration) SetStatus(v string) {
	o.Status = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *Registration) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Registration) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *Registration) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *Registration) SetDeviceId(v string) {
	o.DeviceId = &v
}

func (o Registration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.DeviceId != nil {
		toSerialize["device_id"] = o.DeviceId
	}
	return json.Marshal(toSerialize)
}

type NullableRegistration struct {
	value *Registration
	isSet bool
}

func (v NullableRegistration) Get() *Registration {
	return v.value
}

func (v *NullableRegistration) Set(val *Registration) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistration(val *Registration) *NullableRegistration {
	return &NullableRegistration{value: val, isSet: true}
}

func (v NullableRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


