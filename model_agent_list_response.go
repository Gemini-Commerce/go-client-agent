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

// checks if the AgentListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentListResponse{}

// AgentListResponse struct for AgentListResponse
type AgentListResponse struct {
	Agents []AgentAgentEntity `json:"agents,omitempty"`
	PageToken *string `json:"pageToken,omitempty"`
}

// NewAgentListResponse instantiates a new AgentListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentListResponse() *AgentListResponse {
	this := AgentListResponse{}
	return &this
}

// NewAgentListResponseWithDefaults instantiates a new AgentListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentListResponseWithDefaults() *AgentListResponse {
	this := AgentListResponse{}
	return &this
}

// GetAgents returns the Agents field value if set, zero value otherwise.
func (o *AgentListResponse) GetAgents() []AgentAgentEntity {
	if o == nil || IsNil(o.Agents) {
		var ret []AgentAgentEntity
		return ret
	}
	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentListResponse) GetAgentsOk() ([]AgentAgentEntity, bool) {
	if o == nil || IsNil(o.Agents) {
		return nil, false
	}
	return o.Agents, true
}

// HasAgents returns a boolean if a field has been set.
func (o *AgentListResponse) HasAgents() bool {
	if o != nil && !IsNil(o.Agents) {
		return true
	}

	return false
}

// SetAgents gets a reference to the given []AgentAgentEntity and assigns it to the Agents field.
func (o *AgentListResponse) SetAgents(v []AgentAgentEntity) {
	o.Agents = v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *AgentListResponse) GetPageToken() string {
	if o == nil || IsNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentListResponse) GetPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PageToken) {
		return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *AgentListResponse) HasPageToken() bool {
	if o != nil && !IsNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *AgentListResponse) SetPageToken(v string) {
	o.PageToken = &v
}

func (o AgentListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Agents) {
		toSerialize["agents"] = o.Agents
	}
	if !IsNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}
	return toSerialize, nil
}

type NullableAgentListResponse struct {
	value *AgentListResponse
	isSet bool
}

func (v NullableAgentListResponse) Get() *AgentListResponse {
	return v.value
}

func (v *NullableAgentListResponse) Set(val *AgentListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentListResponse(val *AgentListResponse) *NullableAgentListResponse {
	return &NullableAgentListResponse{value: val, isSet: true}
}

func (v NullableAgentListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

