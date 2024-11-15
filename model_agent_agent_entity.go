/*
agent/service.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package agent

import (
	"encoding/json"
	"time"
)

// checks if the AgentAgentEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentAgentEntity{}

// AgentAgentEntity struct for AgentAgentEntity
type AgentAgentEntity struct {
	Id *string `json:"id,omitempty"`
	Grn *string `json:"grn,omitempty"`
	Code *string `json:"code,omitempty"`
	Firstname *string `json:"firstname,omitempty"`
	Lastname *string `json:"lastname,omitempty"`
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AgentAgentEntity AgentAgentEntity

// NewAgentAgentEntity instantiates a new AgentAgentEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentAgentEntity() *AgentAgentEntity {
	this := AgentAgentEntity{}
	return &this
}

// NewAgentAgentEntityWithDefaults instantiates a new AgentAgentEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentAgentEntityWithDefaults() *AgentAgentEntity {
	this := AgentAgentEntity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AgentAgentEntity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentAgentEntity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AgentAgentEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AgentAgentEntity) SetId(v string) {
	o.Id = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *AgentAgentEntity) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentAgentEntity) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *AgentAgentEntity) HasGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *AgentAgentEntity) SetGrn(v string) {
	o.Grn = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AgentAgentEntity) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentAgentEntity) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AgentAgentEntity) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AgentAgentEntity) SetCode(v string) {
	o.Code = &v
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *AgentAgentEntity) GetFirstname() string {
	if o == nil || IsNil(o.Firstname) {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentAgentEntity) GetFirstnameOk() (*string, bool) {
	if o == nil || IsNil(o.Firstname) {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *AgentAgentEntity) HasFirstname() bool {
	if o != nil && !IsNil(o.Firstname) {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *AgentAgentEntity) SetFirstname(v string) {
	o.Firstname = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *AgentAgentEntity) GetLastname() string {
	if o == nil || IsNil(o.Lastname) {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentAgentEntity) GetLastnameOk() (*string, bool) {
	if o == nil || IsNil(o.Lastname) {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *AgentAgentEntity) HasLastname() bool {
	if o != nil && !IsNil(o.Lastname) {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *AgentAgentEntity) SetLastname(v string) {
	o.Lastname = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AgentAgentEntity) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentAgentEntity) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AgentAgentEntity) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AgentAgentEntity) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *AgentAgentEntity) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentAgentEntity) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *AgentAgentEntity) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *AgentAgentEntity) SetPhone(v string) {
	o.Phone = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AgentAgentEntity) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentAgentEntity) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AgentAgentEntity) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AgentAgentEntity) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AgentAgentEntity) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentAgentEntity) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AgentAgentEntity) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AgentAgentEntity) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o AgentAgentEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentAgentEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
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
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentAgentEntity) UnmarshalJSON(data []byte) (err error) {
	varAgentAgentEntity := _AgentAgentEntity{}

	err = json.Unmarshal(data, &varAgentAgentEntity)

	if err != nil {
		return err
	}

	*o = AgentAgentEntity(varAgentAgentEntity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "grn")
		delete(additionalProperties, "code")
		delete(additionalProperties, "firstname")
		delete(additionalProperties, "lastname")
		delete(additionalProperties, "email")
		delete(additionalProperties, "phone")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *AgentAgentEntity) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *AgentAgentEntity) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableAgentAgentEntity struct {
	value *AgentAgentEntity
	isSet bool
}

func (v NullableAgentAgentEntity) Get() *AgentAgentEntity {
	return v.value
}

func (v *NullableAgentAgentEntity) Set(val *AgentAgentEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentAgentEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentAgentEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentAgentEntity(val *AgentAgentEntity) *NullableAgentAgentEntity {
	return &NullableAgentAgentEntity{value: val, isSet: true}
}

func (v NullableAgentAgentEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentAgentEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


