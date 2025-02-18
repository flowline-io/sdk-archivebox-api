/*
ArchiveBox API

 <h3>Welcome to your ArchiveBox server's REST API <code>[v1 ALPHA]</code> homepage!</h3> <br/> <i><b>WARNING: This API is still in an early development stage and may change!</b></i> <br/> <ul> <li>⬅️ Manage your server: <a href=\"/admin/api/\"><b>Setup API Keys</b></a>, <a href=\"/admin/\">Go to your Server Admin UI</a>, <a href=\"/\">Go to your Snapshots list</a>  <li>💬 Ask questions and get help here: <a href=\"https://zulip.archivebox.io\">ArchiveBox Chat Forum</a></li> <li>🐞 Report API bugs here: <a href=\"https://github.com/ArchiveBox/ArchiveBox/issues\">Github Issues</a></li> <li>📚 ArchiveBox Documentation: <a href=\"https://github.com/ArchiveBox/ArchiveBox/wiki\">Github Wiki</a></li> <li>📜 See the API source code: <a href=\"https://github.com/ArchiveBox/ArchiveBox/blob/dev/archivebox/api\"><code>archivebox/api/</code></a></li> </ul> <small>Served by ArchiveBox v0.8.5rc50 (<a href=\"https://github.com/ArchiveBox/ArchiveBox/commit/1dff8bae4043efb678c1a6fb517dfa7e7a211150\"><code>1dff8bae</code></a>), API powered by <a href=\"https://django-ninja.dev/\"><code>django-ninja</code></a>.</small> 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TokenAuthSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenAuthSchema{}

// TokenAuthSchema Schema for a /check_api_token request
type TokenAuthSchema struct {
	Token string `json:"token"`
}

type _TokenAuthSchema TokenAuthSchema

// NewTokenAuthSchema instantiates a new TokenAuthSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenAuthSchema(token string) *TokenAuthSchema {
	this := TokenAuthSchema{}
	this.Token = token
	return &this
}

// NewTokenAuthSchemaWithDefaults instantiates a new TokenAuthSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenAuthSchemaWithDefaults() *TokenAuthSchema {
	this := TokenAuthSchema{}
	return &this
}

// GetToken returns the Token field value
func (o *TokenAuthSchema) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *TokenAuthSchema) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *TokenAuthSchema) SetToken(v string) {
	o.Token = v
}

func (o TokenAuthSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenAuthSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

func (o *TokenAuthSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTokenAuthSchema := _TokenAuthSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenAuthSchema)

	if err != nil {
		return err
	}

	*o = TokenAuthSchema(varTokenAuthSchema)

	return err
}

type NullableTokenAuthSchema struct {
	value *TokenAuthSchema
	isSet bool
}

func (v NullableTokenAuthSchema) Get() *TokenAuthSchema {
	return v.value
}

func (v *NullableTokenAuthSchema) Set(val *TokenAuthSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenAuthSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenAuthSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenAuthSchema(val *TokenAuthSchema) *NullableTokenAuthSchema {
	return &NullableTokenAuthSchema{value: val, isSet: true}
}

func (v NullableTokenAuthSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenAuthSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


