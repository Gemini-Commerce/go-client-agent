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

// checks if the AgentUpdateAgentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentUpdateAgentRequest{}

// AgentUpdateAgentRequest struct for AgentUpdateAgentRequest
type AgentUpdateAgentRequest struct {
	Payload *AgentUpdatePayload `json:"payload,omitempty"`
	PayloadMask *string `json:"payloadMask,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AgentUpdateAgentRequest AgentUpdateAgentRequest

// NewAgentUpdateAgentRequest instantiates a new AgentUpdateAgentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentUpdateAgentRequest() *AgentUpdateAgentRequest {
	this := AgentUpdateAgentRequest{}
	return &this
}

// NewAgentUpdateAgentRequestWithDefaults instantiates a new AgentUpdateAgentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentUpdateAgentRequestWithDefaults() *AgentUpdateAgentRequest {
	this := AgentUpdateAgentRequest{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *AgentUpdateAgentRequest) GetPayload() AgentUpdatePayload {
	if o == nil || IsNil(o.Payload) {
		var ret AgentUpdatePayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentUpdateAgentRequest) GetPayloadOk() (*AgentUpdatePayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *AgentUpdateAgentRequest) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given AgentUpdatePayload and assigns it to the Payload field.
func (o *AgentUpdateAgentRequest) SetPayload(v AgentUpdatePayload) {
	o.Payload = &v
}

// GetPayloadMask returns the PayloadMask field value if set, zero value otherwise.
func (o *AgentUpdateAgentRequest) GetPayloadMask() string {
	if o == nil || IsNil(o.PayloadMask) {
		var ret string
		return ret
	}
	return *o.PayloadMask
}

// GetPayloadMaskOk returns a tuple with the PayloadMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentUpdateAgentRequest) GetPayloadMaskOk() (*string, bool) {
	if o == nil || IsNil(o.PayloadMask) {
		return nil, false
	}
	return o.PayloadMask, true
}

// HasPayloadMask returns a boolean if a field has been set.
func (o *AgentUpdateAgentRequest) HasPayloadMask() bool {
	if o != nil && !IsNil(o.PayloadMask) {
		return true
	}

	return false
}

// SetPayloadMask gets a reference to the given string and assigns it to the PayloadMask field.
func (o *AgentUpdateAgentRequest) SetPayloadMask(v string) {
	o.PayloadMask = &v
}

func (o AgentUpdateAgentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentUpdateAgentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.PayloadMask) {
		toSerialize["payloadMask"] = o.PayloadMask
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentUpdateAgentRequest) UnmarshalJSON(data []byte) (err error) {
	varAgentUpdateAgentRequest := _AgentUpdateAgentRequest{}

	err = json.Unmarshal(data, &varAgentUpdateAgentRequest)

	if err != nil {
		return err
	}

	*o = AgentUpdateAgentRequest(varAgentUpdateAgentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "payload")
		delete(additionalProperties, "payloadMask")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *AgentUpdateAgentRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *AgentUpdateAgentRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableAgentUpdateAgentRequest struct {
	value *AgentUpdateAgentRequest
	isSet bool
}

func (v NullableAgentUpdateAgentRequest) Get() *AgentUpdateAgentRequest {
	return v.value
}

func (v *NullableAgentUpdateAgentRequest) Set(val *AgentUpdateAgentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentUpdateAgentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentUpdateAgentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentUpdateAgentRequest(val *AgentUpdateAgentRequest) *NullableAgentUpdateAgentRequest {
	return &NullableAgentUpdateAgentRequest{value: val, isSet: true}
}

func (v NullableAgentUpdateAgentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentUpdateAgentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


