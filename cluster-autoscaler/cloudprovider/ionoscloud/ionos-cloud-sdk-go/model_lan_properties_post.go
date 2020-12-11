/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionossdk

import (
	"encoding/json"
)

// LanPropertiesPost struct for LanPropertiesPost
type LanPropertiesPost struct {
	// A name of that resource
	Name *string `json:"name,omitempty"`
	// Does this LAN faces the public Internet or not
	Public *bool `json:"public,omitempty"`
}



// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LanPropertiesPost) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanPropertiesPost) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *LanPropertiesPost) SetName(v string) {
	o.Name = &v
}

// HasName returns a boolean if a field has been set.
func (o *LanPropertiesPost) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}



// GetPublic returns the Public field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *LanPropertiesPost) GetPublic() *bool {
	if o == nil {
		return nil
	}

	return o.Public
}

// GetPublicOk returns a tuple with the Public field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanPropertiesPost) GetPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Public, true
}

// SetPublic sets field value
func (o *LanPropertiesPost) SetPublic(v bool) {
	o.Public = &v
}

// HasPublic returns a boolean if a field has been set.
func (o *LanPropertiesPost) HasPublic() bool {
	if o != nil && o.Public != nil {
		return true
	}

	return false
}


func (o LanPropertiesPost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	

	if o.Public != nil {
		toSerialize["public"] = o.Public
	}
	
	return json.Marshal(toSerialize)
}

type NullableLanPropertiesPost struct {
	value *LanPropertiesPost
	isSet bool
}

func (v NullableLanPropertiesPost) Get() *LanPropertiesPost {
	return v.value
}

func (v *NullableLanPropertiesPost) Set(val *LanPropertiesPost) {
	v.value = val
	v.isSet = true
}

func (v NullableLanPropertiesPost) IsSet() bool {
	return v.isSet
}

func (v *NullableLanPropertiesPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanPropertiesPost(val *LanPropertiesPost) *NullableLanPropertiesPost {
	return &NullableLanPropertiesPost{value: val, isSet: true}
}

func (v NullableLanPropertiesPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanPropertiesPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

