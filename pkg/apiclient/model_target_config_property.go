/*
Daytona Server API

Daytona Server API

API version: v0.0.0-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// checks if the TargetConfigProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetConfigProperty{}

// TargetConfigProperty struct for TargetConfigProperty
type TargetConfigProperty struct {
	// DefaultValue is converted into the appropriate type based on the Type If the property is a FilePath, the DefaultValue is a path to a directory
	DefaultValue *string `json:"defaultValue,omitempty"`
	// Brief description of the property
	Description *string `json:"description,omitempty"`
	// A regex string matched with the name of the target config to determine if the property should be disabled If the regex matches the target config name, the property will be disabled E.g. \"^local$\" will disable the property for the local target
	DisabledPredicate *string `json:"disabledPredicate,omitempty"`
	InputMasked       *bool   `json:"inputMasked,omitempty"`
	// Options is only used if the Type is TargetConfigPropertyTypeOption
	Options []string `json:"options,omitempty"`
	// Suggestions is an optional list of auto-complete values to assist the user while filling the field
	Suggestions []string                          `json:"suggestions,omitempty"`
	Type        *ProviderTargetConfigPropertyType `json:"type,omitempty"`
}

// NewTargetConfigProperty instantiates a new TargetConfigProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetConfigProperty() *TargetConfigProperty {
	this := TargetConfigProperty{}
	return &this
}

// NewTargetConfigPropertyWithDefaults instantiates a new TargetConfigProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetConfigPropertyWithDefaults() *TargetConfigProperty {
	this := TargetConfigProperty{}
	return &this
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *TargetConfigProperty) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetConfigProperty) GetDefaultValueOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *TargetConfigProperty) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *TargetConfigProperty) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TargetConfigProperty) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetConfigProperty) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TargetConfigProperty) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TargetConfigProperty) SetDescription(v string) {
	o.Description = &v
}

// GetDisabledPredicate returns the DisabledPredicate field value if set, zero value otherwise.
func (o *TargetConfigProperty) GetDisabledPredicate() string {
	if o == nil || IsNil(o.DisabledPredicate) {
		var ret string
		return ret
	}
	return *o.DisabledPredicate
}

// GetDisabledPredicateOk returns a tuple with the DisabledPredicate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetConfigProperty) GetDisabledPredicateOk() (*string, bool) {
	if o == nil || IsNil(o.DisabledPredicate) {
		return nil, false
	}
	return o.DisabledPredicate, true
}

// HasDisabledPredicate returns a boolean if a field has been set.
func (o *TargetConfigProperty) HasDisabledPredicate() bool {
	if o != nil && !IsNil(o.DisabledPredicate) {
		return true
	}

	return false
}

// SetDisabledPredicate gets a reference to the given string and assigns it to the DisabledPredicate field.
func (o *TargetConfigProperty) SetDisabledPredicate(v string) {
	o.DisabledPredicate = &v
}

// GetInputMasked returns the InputMasked field value if set, zero value otherwise.
func (o *TargetConfigProperty) GetInputMasked() bool {
	if o == nil || IsNil(o.InputMasked) {
		var ret bool
		return ret
	}
	return *o.InputMasked
}

// GetInputMaskedOk returns a tuple with the InputMasked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetConfigProperty) GetInputMaskedOk() (*bool, bool) {
	if o == nil || IsNil(o.InputMasked) {
		return nil, false
	}
	return o.InputMasked, true
}

// HasInputMasked returns a boolean if a field has been set.
func (o *TargetConfigProperty) HasInputMasked() bool {
	if o != nil && !IsNil(o.InputMasked) {
		return true
	}

	return false
}

// SetInputMasked gets a reference to the given bool and assigns it to the InputMasked field.
func (o *TargetConfigProperty) SetInputMasked(v bool) {
	o.InputMasked = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *TargetConfigProperty) GetOptions() []string {
	if o == nil || IsNil(o.Options) {
		var ret []string
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetConfigProperty) GetOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *TargetConfigProperty) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []string and assigns it to the Options field.
func (o *TargetConfigProperty) SetOptions(v []string) {
	o.Options = v
}

// GetSuggestions returns the Suggestions field value if set, zero value otherwise.
func (o *TargetConfigProperty) GetSuggestions() []string {
	if o == nil || IsNil(o.Suggestions) {
		var ret []string
		return ret
	}
	return o.Suggestions
}

// GetSuggestionsOk returns a tuple with the Suggestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetConfigProperty) GetSuggestionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Suggestions) {
		return nil, false
	}
	return o.Suggestions, true
}

// HasSuggestions returns a boolean if a field has been set.
func (o *TargetConfigProperty) HasSuggestions() bool {
	if o != nil && !IsNil(o.Suggestions) {
		return true
	}

	return false
}

// SetSuggestions gets a reference to the given []string and assigns it to the Suggestions field.
func (o *TargetConfigProperty) SetSuggestions(v []string) {
	o.Suggestions = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TargetConfigProperty) GetType() ProviderTargetConfigPropertyType {
	if o == nil || IsNil(o.Type) {
		var ret ProviderTargetConfigPropertyType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetConfigProperty) GetTypeOk() (*ProviderTargetConfigPropertyType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TargetConfigProperty) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ProviderTargetConfigPropertyType and assigns it to the Type field.
func (o *TargetConfigProperty) SetType(v ProviderTargetConfigPropertyType) {
	o.Type = &v
}

func (o TargetConfigProperty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetConfigProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultValue) {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisabledPredicate) {
		toSerialize["disabledPredicate"] = o.DisabledPredicate
	}
	if !IsNil(o.InputMasked) {
		toSerialize["inputMasked"] = o.InputMasked
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Suggestions) {
		toSerialize["suggestions"] = o.Suggestions
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTargetConfigProperty struct {
	value *TargetConfigProperty
	isSet bool
}

func (v NullableTargetConfigProperty) Get() *TargetConfigProperty {
	return v.value
}

func (v *NullableTargetConfigProperty) Set(val *TargetConfigProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetConfigProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetConfigProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetConfigProperty(val *TargetConfigProperty) *NullableTargetConfigProperty {
	return &NullableTargetConfigProperty{value: val, isSet: true}
}

func (v NullableTargetConfigProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetConfigProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}