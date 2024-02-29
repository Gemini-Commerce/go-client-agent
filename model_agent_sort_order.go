/*
agent/service.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package agent

import (
	"encoding/json"
	"fmt"
)

// AgentSortOrder the model 'AgentSortOrder'
type AgentSortOrder string

// List of agentSortOrder
const (
	DESC AgentSortOrder = "DESC"
	ASC AgentSortOrder = "ASC"
)

// All allowed values of AgentSortOrder enum
var AllowedAgentSortOrderEnumValues = []AgentSortOrder{
	"DESC",
	"ASC",
}

func (v *AgentSortOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AgentSortOrder(value)
	for _, existing := range AllowedAgentSortOrderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AgentSortOrder", value)
}

// NewAgentSortOrderFromValue returns a pointer to a valid AgentSortOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentSortOrderFromValue(v string) (*AgentSortOrder, error) {
	ev := AgentSortOrder(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentSortOrder: valid values are %v", v, AllowedAgentSortOrderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentSortOrder) IsValid() bool {
	for _, existing := range AllowedAgentSortOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agentSortOrder value
func (v AgentSortOrder) Ptr() *AgentSortOrder {
	return &v
}

type NullableAgentSortOrder struct {
	value *AgentSortOrder
	isSet bool
}

func (v NullableAgentSortOrder) Get() *AgentSortOrder {
	return v.value
}

func (v *NullableAgentSortOrder) Set(val *AgentSortOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentSortOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentSortOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentSortOrder(val *AgentSortOrder) *NullableAgentSortOrder {
	return &NullableAgentSortOrder{value: val, isSet: true}
}

func (v NullableAgentSortOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentSortOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
