/*
agent/service.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package agent

import (
	"encoding/json"
)

// checks if the AgentUpdatePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentUpdatePayload{}

// AgentUpdatePayload struct for AgentUpdatePayload
type AgentUpdatePayload struct {
	Firstname *string `json:"firstname,omitempty"`
	Lastname *string `json:"lastname,omitempty"`
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

// NewAgentUpdatePayload instantiates a new AgentUpdatePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentUpdatePayload() *AgentUpdatePayload {
	this := AgentUpdatePayload{}
	return &this
}

// NewAgentUpdatePayloadWithDefaults instantiates a new AgentUpdatePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentUpdatePayloadWithDefaults() *AgentUpdatePayload {
	this := AgentUpdatePayload{}
	return &this
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *AgentUpdatePayload) GetFirstname() string {
	if o == nil || IsNil(o.Firstname) {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentUpdatePayload) GetFirstnameOk() (*string, bool) {
	if o == nil || IsNil(o.Firstname) {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *AgentUpdatePayload) HasFirstname() bool {
	if o != nil && !IsNil(o.Firstname) {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *AgentUpdatePayload) SetFirstname(v string) {
	o.Firstname = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *AgentUpdatePayload) GetLastname() string {
	if o == nil || IsNil(o.Lastname) {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentUpdatePayload) GetLastnameOk() (*string, bool) {
	if o == nil || IsNil(o.Lastname) {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *AgentUpdatePayload) HasLastname() bool {
	if o != nil && !IsNil(o.Lastname) {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *AgentUpdatePayload) SetLastname(v string) {
	o.Lastname = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AgentUpdatePayload) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentUpdatePayload) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AgentUpdatePayload) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AgentUpdatePayload) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *AgentUpdatePayload) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentUpdatePayload) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *AgentUpdatePayload) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *AgentUpdatePayload) SetPhone(v string) {
	o.Phone = &v
}

func (o AgentUpdatePayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentUpdatePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Firstname) {
		toSerialize["firstname"] = o.Firstname
	}
	if !IsNil(o.Lastname) {
		toSerialize["lastname"] = o.Lastname
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	return toSerialize, nil
}

type NullableAgentUpdatePayload struct {
	value *AgentUpdatePayload
	isSet bool
}

func (v NullableAgentUpdatePayload) Get() *AgentUpdatePayload {
	return v.value
}

func (v *NullableAgentUpdatePayload) Set(val *AgentUpdatePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentUpdatePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentUpdatePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentUpdatePayload(val *AgentUpdatePayload) *NullableAgentUpdatePayload {
	return &NullableAgentUpdatePayload{value: val, isSet: true}
}

func (v NullableAgentUpdatePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentUpdatePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


