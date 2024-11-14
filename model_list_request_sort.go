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

// checks if the ListRequestSort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRequestSort{}

// ListRequestSort struct for ListRequestSort
type ListRequestSort struct {
	EvaluationOrder *int64 `json:"evaluationOrder,omitempty"`
	Field *SortSortField `json:"field,omitempty"`
	Order *AgentSortOrder `json:"order,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListRequestSort ListRequestSort

// NewListRequestSort instantiates a new ListRequestSort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRequestSort() *ListRequestSort {
	this := ListRequestSort{}
	var field SortSortField = SORTSORTFIELD_UNKNOWN
	this.Field = &field
	var order AgentSortOrder = AGENTSORTORDER_DESC
	this.Order = &order
	return &this
}

// NewListRequestSortWithDefaults instantiates a new ListRequestSort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRequestSortWithDefaults() *ListRequestSort {
	this := ListRequestSort{}
	var field SortSortField = SORTSORTFIELD_UNKNOWN
	this.Field = &field
	var order AgentSortOrder = AGENTSORTORDER_DESC
	this.Order = &order
	return &this
}

// GetEvaluationOrder returns the EvaluationOrder field value if set, zero value otherwise.
func (o *ListRequestSort) GetEvaluationOrder() int64 {
	if o == nil || IsNil(o.EvaluationOrder) {
		var ret int64
		return ret
	}
	return *o.EvaluationOrder
}

// GetEvaluationOrderOk returns a tuple with the EvaluationOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRequestSort) GetEvaluationOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.EvaluationOrder) {
		return nil, false
	}
	return o.EvaluationOrder, true
}

// &#39;Has&#39;EvaluationOrder returns a boolean if a field has been set.
func (o *ListRequestSort) &#39;Has&#39;EvaluationOrder() bool {
	if o != nil && !IsNil(o.EvaluationOrder) {
		return true
	}

	return false
}

// SetEvaluationOrder gets a reference to the given int64 and assigns it to the EvaluationOrder field.
func (o *ListRequestSort) SetEvaluationOrder(v int64) {
	o.EvaluationOrder = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *ListRequestSort) GetField() SortSortField {
	if o == nil || IsNil(o.Field) {
		var ret SortSortField
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRequestSort) GetFieldOk() (*SortSortField, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// &#39;Has&#39;Field returns a boolean if a field has been set.
func (o *ListRequestSort) &#39;Has&#39;Field() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given SortSortField and assigns it to the Field field.
func (o *ListRequestSort) SetField(v SortSortField) {
	o.Field = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ListRequestSort) GetOrder() AgentSortOrder {
	if o == nil || IsNil(o.Order) {
		var ret AgentSortOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRequestSort) GetOrderOk() (*AgentSortOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// &#39;Has&#39;Order returns a boolean if a field has been set.
func (o *ListRequestSort) &#39;Has&#39;Order() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given AgentSortOrder and assigns it to the Order field.
func (o *ListRequestSort) SetOrder(v AgentSortOrder) {
	o.Order = &v
}

func (o ListRequestSort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListRequestSort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EvaluationOrder) {
		toSerialize["evaluationOrder"] = o.EvaluationOrder
	}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListRequestSort) UnmarshalJSON(data []byte) (err error) {
	varListRequestSort := _ListRequestSort{}

	err = json.Unmarshal(data, &varListRequestSort)

	if err != nil {
		return err
	}

	*o = ListRequestSort(varListRequestSort)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "evaluationOrder")
		delete(additionalProperties, "field")
		delete(additionalProperties, "order")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ListRequestSort) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ListRequestSort) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableListRequestSort struct {
	value *ListRequestSort
	isSet bool
}

func (v NullableListRequestSort) Get() *ListRequestSort {
	return v.value
}

func (v *NullableListRequestSort) Set(val *ListRequestSort) {
	v.value = val
	v.isSet = true
}

func (v NullableListRequestSort) IsSet() bool {
	return v.isSet
}

func (v *NullableListRequestSort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRequestSort(val *ListRequestSort) *NullableListRequestSort {
	return &NullableListRequestSort{value: val, isSet: true}
}

func (v NullableListRequestSort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRequestSort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


