/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// UserInvitationPayload Object to invite users to join the organization.
type UserInvitationPayload struct {
	// List of user invitations.
	Data *[]UserInvitationData `json:"data,omitempty"`
}

// NewUserInvitationPayload instantiates a new UserInvitationPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInvitationPayload() *UserInvitationPayload {
	this := UserInvitationPayload{}
	return &this
}

// NewUserInvitationPayloadWithDefaults instantiates a new UserInvitationPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInvitationPayloadWithDefaults() *UserInvitationPayload {
	this := UserInvitationPayload{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UserInvitationPayload) GetData() []UserInvitationData {
	if o == nil || o.Data == nil {
		var ret []UserInvitationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvitationPayload) GetDataOk() (*[]UserInvitationData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UserInvitationPayload) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []UserInvitationData and assigns it to the Data field.
func (o *UserInvitationPayload) SetData(v []UserInvitationData) {
	o.Data = &v
}

func (o UserInvitationPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableUserInvitationPayload struct {
	value *UserInvitationPayload
	isSet bool
}

func (v NullableUserInvitationPayload) Get() *UserInvitationPayload {
	return v.value
}

func (v *NullableUserInvitationPayload) Set(val *UserInvitationPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInvitationPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInvitationPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInvitationPayload(val *UserInvitationPayload) *NullableUserInvitationPayload {
	return &NullableUserInvitationPayload{value: val, isSet: true}
}

func (v NullableUserInvitationPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInvitationPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}